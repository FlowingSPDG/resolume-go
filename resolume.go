package resolume

import (
	"net"
	"net/url"

	"github.com/FlowingSPDG/resolume-go/internal"
)

// Client API Client
type Client struct {
	c *internal.ClientWithResponses
}

// NewClient Get API client
func NewClient(host, port string) (*Client, error) {
	u := &url.URL{
		Host:   net.JoinHostPort(host, port),
		Scheme: "http",
		Path:   "/api/v1/", // Fixed for v1 only
	}
	c, err := internal.NewClientWithResponses(u.String())
	if err != nil {
		return nil, err
	}
	return &Client{
		c: c,
	}, nil
}
