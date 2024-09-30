package handler

import (
	"context"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/logger"
	"golang.org/x/time/rate"
)

// RateLimiterMiddleware implements a rate limiting middleware for HTTP requests.
type RateLimiterMiddleware struct {
	limiter *rate.Limiter
	logger  logger.Logger
}

// NewRateLimiterMiddleware creates a new RateLimiterMiddleware instance.
func NewRateLimiterMiddleware(requestsPerSecond float64, burst int) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		limiter: rate.NewLimiter(rate.Limit(requestsPerSecond), burst),
		logger:  &logger.NoOpLogger{},
	}
}

// Process applies rate limiting before passing the request to the next middleware.
func (m *RateLimiterMiddleware) Process(ctx context.Context, opts *RequestOptions, next func(context.Context, *RequestOptions) (*http.Response, error)) (*http.Response, error) {
	m.logger.Debug("Processing request with rate limiter middleware")

	// Wait for rate limiter permission
	if err := m.limiter.Wait(ctx); err != nil {
		return nil, err
	}

	// Execute the next middleware in the chain
	return next(ctx, opts)
}

// SetLogger sets the logger for the middleware.
func (m *RateLimiterMiddleware) SetLogger(l logger.Logger) {
	m.logger = l
}
