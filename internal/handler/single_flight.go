package handler

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/sync/singleflight"
)

// SingleFlight implements the singleflight pattern to deduplicate concurrent identical requests.
type SingleFlight struct {
	handler        *Handler
	sfGroup        *singleflight.Group
	circuitBreaker *CircuitBreaker
}

// NewSingleFlight creates a new SingleFlight instance with the specified Handler.
func NewSingleFlight(handler *Handler) *SingleFlight {
	return &SingleFlight{
		handler:        handler,
		sfGroup:        &singleflight.Group{},
		circuitBreaker: NewCircuitBreaker(handler),
	}
}

// do performs an HTTP request, potentially using singleflight to deduplicate concurrent identical requests.
func (s *SingleFlight) do(ctx context.Context, options *RequestOptions) (*http.Response, error) {
	if s.handler.UseSingleFlight {
		return s.doWithSingleFlight(ctx, options)
	}
	return s.doWithoutSingleFlight(ctx, options)
}

// doWithSingleFlight performs an HTTP request using singleflight to deduplicate concurrent identical requests.
func (s *SingleFlight) doWithSingleFlight(ctx context.Context, options *RequestOptions) (*http.Response, error) {
	// Generate a unique key for the request
	key := fmt.Sprintf("%s:%s:%x", options.Method, options.URL, options.Body)

	// Use singleflight to execute the request
	result, err, _ := s.sfGroup.Do(key, func() (interface{}, error) {
		return s.circuitBreaker.do(ctx, options)
	})
	if err != nil {
		return nil, err
	}

	return result.(*http.Response), nil
}

// doWithoutSingleFlight performs an HTTP request without using singleflight.
func (s *SingleFlight) doWithoutSingleFlight(ctx context.Context, options *RequestOptions) (*http.Response, error) {
	// Execute the request through the circuit breaker
	return s.circuitBreaker.do(ctx, options)
}
