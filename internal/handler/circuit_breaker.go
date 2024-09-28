package handler

import (
	"context"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/errors"
	"github.com/sony/gobreaker"
	"go.uber.org/zap"
)

// CircuitBreaker implements the circuit breaker pattern to prevent cascading failures.
type CircuitBreaker struct {
	handler *Handler
	breaker *gobreaker.CircuitBreaker
	retry   *Retry
}

// NewCircuitBreaker creates a new CircuitBreaker instance with the specified Handler.
func NewCircuitBreaker(handler *Handler) *CircuitBreaker {
	return &CircuitBreaker{
		handler: handler,
		breaker: gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "HTTPCircuitBreaker",
			MaxRequests: handler.CircuitBreakerMaxRequests,
			Interval:    handler.CircuitBreakerInterval,
			Timeout:     handler.CircuitBreakerTimeout,
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
				return counts.Requests >= 3 && failureRatio >= 0.6
			},
			OnStateChange: func(name string, from, to gobreaker.State) {
				handler.Logger.Warn("Circuit breaker state changed",
					zap.String("name", name),
					zap.String("from", from.String()),
					zap.String("to", to.String()))
			},
			IsSuccessful: nil,
		}),
		retry: NewRetry(handler),
	}
}

// do performs an HTTP request, potentially using the circuit breaker to prevent cascading failures.
func (c *CircuitBreaker) do(ctx context.Context, options *RequestOptions) (*http.Response, error) {
	if c.handler.UseCircuitBreaker {
		return c.executeWithCircuitBreaker(ctx, options)
	}
	return c.executeWithoutCircuitBreaker(ctx, options)
}

// executeWithCircuitBreaker performs an HTTP request using the circuit breaker.
func (c *CircuitBreaker) executeWithCircuitBreaker(ctx context.Context, options *RequestOptions) (*http.Response, error) {
	// Execute the request through the circuit breaker
	result, err := c.breaker.Execute(func() (interface{}, error) {
		return c.retry.do(ctx, options)
	})
	if err != nil {
		// Handle different circuit breaker errors
		switch err {
		case gobreaker.ErrOpenState:
			return nil, errors.NewError(errors.ErrorTypeCircuitOpen, "Circuit breaker is open", err, nil)
		case gobreaker.ErrTooManyRequests:
			return nil, errors.NewError(errors.ErrorTypeCircuitExhausted, "Circuit breaker request limit reached", err, nil)
		default:
			return nil, err
		}
	}

	return result.(*http.Response), nil
}

// executeWithoutCircuitBreaker performs an HTTP request without using the circuit breaker.
func (c *CircuitBreaker) executeWithoutCircuitBreaker(ctx context.Context, options *RequestOptions) (*http.Response, error) {
	// Execute the request with retry logic
	return c.retry.do(ctx, options)
}
