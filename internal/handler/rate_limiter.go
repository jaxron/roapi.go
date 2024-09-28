package handler

import (
	"context"
	"net/http"

	"golang.org/x/time/rate"
)

// RateLimiter implements a rate limiting middleware for HTTP requests.
type RateLimiter struct {
	handler *Handler
	limiter *rate.Limiter
}

// NewRateLimiter creates a new RateLimiter instance with the specified Handler.
// It initializes the rate limiter based on the Handler's configuration.
func NewRateLimiter(handler *Handler) *RateLimiter {
	return &RateLimiter{
		handler: handler,
		limiter: rate.NewLimiter(rate.Limit(handler.RateLimitRequestsPerSecond), handler.RateLimitBurst),
	}
}

// do performs a rate-limited HTTP request.
// It waits for the rate limiter before executing the request.
func (r *RateLimiter) do(ctx context.Context, options *RequestOptions) (*http.Response, error) {
	if r.handler.UseRateLimiter {
		// Wait for rate limiter permission
		if err := r.limiter.Wait(ctx); err != nil {
			return nil, err
		}
	}

	// Execute the HTTP request
	return r.handler.performRequest(ctx, options)
}
