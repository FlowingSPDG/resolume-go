package resolume

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("localhost", "8080")
	if err != nil {
		t.Errorf("NewClient() error = %v", err)
		return
	}
	if client == nil {
		t.Error("NewClient() returned nil client")
	}
}

func TestGetProduct(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/product" {
			t.Errorf("Expected path /api/v1/product, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{
			"name": "Arena",
			"major": 7,
			"minor": 8,
			"micro": 0,
			"revision": 1234
		}`))
	}))
	defer server.Close()

	// Create client using test server URL
	serverURL, _ := url.Parse(server.URL + "/api/v1")
	client := &Client{
		baseURL:    serverURL,
		httpClient: server.Client(),
	}

	// Test GetProduct
	product, err := client.GetProduct(context.Background())
	if err != nil {
		t.Errorf("GetProduct() error = %v", err)
		return
	}

	// Verify response
	if product.Name != "Arena" {
		t.Errorf("Expected product name Arena, got %s", product.Name)
	}
	if product.Major != 7 {
		t.Errorf("Expected major version 7, got %d", product.Major)
	}
}

func TestGetEffects(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/effects" {
			t.Errorf("Expected path /api/v1/effects, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{
			"video": [
				{
					"idstring": "A139",
					"name": "Blow",
					"presets": [
						{
							"id": 2026888427,
							"name": "Bright Lines"
						}
					]
				}
			]
		}`))
	}))
	defer server.Close()

	// Create client using test server URL
	serverURL, _ := url.Parse(server.URL + "/api/v1")
	client := &Client{
		baseURL:    serverURL,
		httpClient: server.Client(),
	}

	// Test GetEffects
	effects, err := client.GetEffects(context.Background())
	if err != nil {
		t.Errorf("GetEffects() error = %v", err)
		return
	}

	// Verify response
	if len(effects.Video) != 1 {
		t.Errorf("Expected 1 video effect, got %d", len(effects.Video))
		return
	}
	if effects.Video[0].Name != "Blow" {
		t.Errorf("Expected effect name Blow, got %s", effects.Video[0].Name)
	}
}

func TestGetSources(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/sources" {
			t.Errorf("Expected path /api/v1/sources, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{
			"video": [
				{
					"idstring": "A401",
					"name": "Gradient",
					"presets": [
						{
							"id": 2026883901,
							"name": "Dutch Sky"
						}
					]
				}
			]
		}`))
	}))
	defer server.Close()

	// Create client using test server URL
	serverURL, _ := url.Parse(server.URL + "/api/v1")
	client := &Client{
		baseURL:    serverURL,
		httpClient: server.Client(),
	}

	// Test GetSources
	sources, err := client.GetSources(context.Background())
	if err != nil {
		t.Errorf("GetSources() error = %v", err)
		return
	}

	// Verify response
	if len(sources.Video) != 1 {
		t.Errorf("Expected 1 video source, got %d", len(sources.Video))
		return
	}
	if sources.Video[0].Name != "Gradient" {
		t.Errorf("Expected source name Gradient, got %s", sources.Video[0].Name)
	}
}
