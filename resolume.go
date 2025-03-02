package resolume

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"path"
)

// Client represents a Resolume API client
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
}

// NewClient creates a new Resolume API client
func NewClient(host, port string) (*Client, error) {
	baseURL := &url.URL{
		Host:   net.JoinHostPort(host, port),
		Scheme: "http",
		Path:   "/api/v1",
	}

	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}, nil
}

// GetProduct retrieves product information
func (c *Client) GetProduct() (*ProductInfo, error) {
	endpoint := "/product"
	var product ProductInfo
	if err := c.get(endpoint, &product); err != nil {
		return nil, err
	}
	return &product, nil
}

// GetEffects retrieves available effects
func (c *Client) GetEffects() (*Effects, error) {
	endpoint := "/effects"
	var effects Effects
	if err := c.get(endpoint, &effects); err != nil {
		return nil, err
	}
	return &effects, nil
}

// GetSources retrieves available sources
func (c *Client) GetSources() (*Sources, error) {
	endpoint := "/sources"
	var sources Sources
	if err := c.get(endpoint, &sources); err != nil {
		return nil, err
	}
	return &sources, nil
}

// GetDummyThumbnail retrieves the dummy thumbnail used for clips without a thumbnail
func (c *Client) GetDummyThumbnail() (io.ReadCloser, error) {
	endpoint := "/composition/thumbnail/dummy"
	resp, err := c.getRaw(endpoint)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

// GetClipThumbnail retrieves the thumbnail for a specific clip
func (c *Client) GetClipThumbnail(layerIndex, clipIndex int) (io.ReadCloser, error) {
	endpoint := fmt.Sprintf("/composition/layers/%d/clips/%d/thumbnail", layerIndex, clipIndex)
	resp, err := c.getRaw(endpoint)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

// SetClipThumbnail sets a custom thumbnail for a specific clip
func (c *Client) SetClipThumbnail(layerIndex, clipIndex int, thumbnail io.Reader) error {
	endpoint := fmt.Sprintf("/composition/layers/%d/clips/%d/thumbnail", layerIndex, clipIndex)

	// Create multipart form
	body := &bytes.Buffer{}
	writer := c.createMultipartForm(body, thumbnail)

	// Set content type
	contentType := fmt.Sprintf("multipart/form-data; boundary=%s", writer.Boundary())

	// Send request
	req, err := http.NewRequest(http.MethodPost, c.url(endpoint), body)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", contentType)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

// ResetClipThumbnail resets a clip's thumbnail to the default
func (c *Client) ResetClipThumbnail(layerIndex, clipIndex int) error {
	endpoint := fmt.Sprintf("/composition/layers/%d/clips/%d/thumbnail", layerIndex, clipIndex)
	return c.delete(endpoint)
}

// Error represents an API error response
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("API error: %d - %s", e.Code, e.Message)
}

// get performs a GET request to the specified endpoint and decodes the JSON response
func (c *Client) get(endpoint string, v interface{}) error {
	resp, err := c.getRaw(endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	return nil
}

// getRaw performs a GET request and returns the raw response
func (c *Client) getRaw(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.url(endpoint), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		var apiErr Error
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
		}
		return nil, &apiErr
	}

	return resp, nil
}

// post performs a POST request to the specified endpoint
func (c *Client) post(endpoint string, body interface{}, v interface{}) error {
	return c.doRequest(http.MethodPost, endpoint, body, v)
}

// put performs a PUT request to the specified endpoint
func (c *Client) put(endpoint string, body interface{}, v interface{}) error {
	return c.doRequest(http.MethodPut, endpoint, body, v)
}

// delete performs a DELETE request to the specified endpoint
func (c *Client) delete(endpoint string) error {
	return c.doRequest(http.MethodDelete, endpoint, nil, nil)
}

// doRequest performs an HTTP request
func (c *Client) doRequest(method, endpoint string, body, v interface{}) error {
	url := c.url(endpoint)

	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %v", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var apiErr Error
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return fmt.Errorf("HTTP error: %d", resp.StatusCode)
		}
		return &apiErr
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return fmt.Errorf("failed to decode response: %v", err)
		}
	}

	return nil
}

// url builds the full URL for an endpoint
func (c *Client) url(endpoint string) string {
	u := *c.baseURL
	u.Path = path.Join(u.Path, endpoint)
	return u.String()
}

// createMultipartForm creates a multipart form writer
func (c *Client) createMultipartForm(body *bytes.Buffer, thumbnail io.Reader) *multipart.Writer {
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "thumbnail")
	if err != nil {
		return nil
	}

	io.Copy(part, thumbnail)
	writer.Close()
	return writer
}
