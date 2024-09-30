package client

import (
	"context"
	"net/http"
	"time"

	"github.com/jaxron/roapi.go/pkg/client/errors"
	"github.com/jaxron/roapi.go/pkg/client/logger"
	"github.com/sony/gobreaker"
	"go.uber.org/zap"
)

// CircuitBreakerMiddleware implements the circuit breaker pattern to prevent cascading failures.
type CircuitBreakerMiddleware struct {
	breaker *gobreaker.CircuitBreaker
	logger  logger.Logger
}

// NewCircuitBreakerMiddleware creates a new CircuitBreakerMiddleware instance.
func NewCircuitBreakerMiddleware(maxRequests uint32, interval, timeout time.Duration) *CircuitBreakerMiddleware {
	middleware := &CircuitBreakerMiddleware{
		breaker: nil,
		logger:  &logger.NoOpLogger{},
	}

	breaker := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "HTTPCircuitBreaker",
		MaxRequests: maxRequests,
		Interval:    interval,
		Timeout:     timeout,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.6
		},
		OnStateChange: func(name string, from, to gobreaker.State) {
			middleware.logger.Warn("Circuit breaker state changed",
				zap.String("name", name),
				zap.String("from", from.String()),
				zap.String("to", to.String()))
		},
		IsSuccessful: nil,
	})
	middleware.breaker = breaker

	return middleware
}

// Process applies the circuit breaker before passing the request to the next middleware.
func (m *CircuitBreakerMiddleware) Process(ctx context.Context, opts *RequestOptions, next func(context.Context, *RequestOptions) (*http.Response, error)) (*http.Response, error) {
	m.logger.Debug("Processing request with circuit breaker middleware")

	result, err := m.breaker.Execute(func() (interface{}, error) {
		return next(ctx, opts)
	})
	if err != nil {
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

// SetLogger sets the logger for the middleware.
func (m *CircuitBreakerMiddleware) SetLogger(l logger.Logger) {
	m.logger = l
}
