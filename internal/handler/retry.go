package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/jaxron/roapi.go/pkg/errors"
	"go.uber.org/zap"
)

// Retry implements retry logic for HTTP requests with exponential backoff.
type Retry struct {
	handler *Handler
	limiter *RateLimiter
}

// NewRetry creates a new Retry instance with the specified Handler.
func NewRetry(handler *Handler) *Retry {
	return &Retry{
		handler: handler,
		limiter: NewRateLimiter(handler),
	}
}

// do performs an HTTP request, potentially using retry logic with exponential backoff.
func (r *Retry) do(ctx context.Context, options *RequestOptions) (*http.Response, error) {
	if r.handler.UseRetry {
		return r.executeWithRetry(ctx, options)
	}
	return r.executeWithoutRetry(ctx, options)
}

// executeWithRetry performs an HTTP request with retry logic using exponential backoff.
func (r *Retry) executeWithRetry(ctx context.Context, options *RequestOptions) (*http.Response, error) {
	// Create an exponential backoff strategy with a maximum number of retries
	expBackoff := backoff.WithMaxRetries(backoff.NewExponentialBackOff(
		backoff.WithInitialInterval(r.handler.RetryInitialInterval),
		backoff.WithMaxInterval(r.handler.RetryMaxInterval),
	), r.handler.RetryMaxAttempts)
	backoffStrategy := backoff.WithContext(expBackoff, ctx)

	var resp *http.Response
	var err error

	// Retry the request using the backoff strategy
	retryErr := backoff.RetryNotify(
		func() error {
			// Execute the request through the rate limiter
			resp, err = r.limiter.do(ctx, options)
			// Handle the error and determine if a retry is needed
			return r.handleRetryError(err)
		},
		backoffStrategy,
		func(err error, duration time.Duration) {
			// Log retry attempts
			r.handler.Logger.Warn("Retrying request", zap.Error(err), zap.Duration("retry_in", duration))
		},
	)

	// If all retries have been exhausted and there's still an error, return it
	if retryErr != nil {
		return nil, retryErr
	}

	// Return the successful response
	return resp, nil
}

// executeWithoutRetry performs an HTTP request without using retry logic.
func (r *Retry) executeWithoutRetry(ctx context.Context, options *RequestOptions) (*http.Response, error) {
	// Execute the request through the rate limiter
	return r.limiter.do(ctx, options)
}

// handleRetryError determines whether to retry the request based on the error type.
func (r *Retry) handleRetryError(err error) error {
	if err != nil {
		if errors.IsTemporary(err) {
			return err // This will trigger a retry for temporary errors
		}
		return backoff.Permanent(err) // This will stop retries for permanent errors
	}
	return nil // Success, stop retrying
}
