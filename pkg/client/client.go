package client

import (
	"context"
	"net/http"

	"github.com/jaxron/roapi.go/internal/handler"
)

// Client represents the main client for interacting with the Roblox API.
type Client struct {
	Handler *handler.Handler
}

// NewClient creates a new Client with the specified options.
func NewClient(opts ...Option) *Client {
	// Create a new handler with default settings
	handler := handler.NewHandler()

	// Apply all provided options to customize the handler
	for _, opt := range opts {
		opt(handler)
	}

	return &Client{
		Handler: handler,
	}
}

// Do executes an HTTP request with the given context and options.
// It returns the HTTP response and any error encountered.
func (c *Client) Do(ctx context.Context, options *handler.RequestOptions) (*http.Response, error) {
	return c.Handler.Do(ctx, options)
}
