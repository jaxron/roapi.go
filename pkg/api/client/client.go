package client

import (
	"context"
	"net/http"
	"time"

	"github.com/jaxron/roapi.go/internal/handler"
)

// Client represents the main client for interacting with the Roblox API.
type Client struct {
	Handler *handler.Handler
}

// NewClient creates a new Client with the specified options.
func NewClient(opts ...Option) *Client {
	// Create a new handler with default middleware
	h := handler.NewHandler(10 * time.Second)

	// Apply all provided options to customize the handler
	for _, opt := range opts {
		opt(h)
	}

	return &Client{
		Handler: h,
	}
}

// Do executes an HTTP request with the given context and options.
// It returns the HTTP response and any error encountered.
func (c *Client) Do(ctx context.Context, opts *handler.RequestOptions) (*http.Response, error) {
	return c.Handler.Do(ctx, opts)
}
