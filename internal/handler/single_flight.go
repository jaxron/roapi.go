package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/logger"
	"golang.org/x/sync/singleflight"
)

// SingleFlightMiddleware implements the singleflight pattern to deduplicate concurrent identical requests.
type SingleFlightMiddleware struct {
	sfGroup *singleflight.Group
	logger  logger.Logger
}

// NewSingleFlightMiddleware creates a new SingleFlightMiddleware instance.
func NewSingleFlightMiddleware() *SingleFlightMiddleware {
	return &SingleFlightMiddleware{
		sfGroup: &singleflight.Group{},
		logger:  &logger.NoOpLogger{},
	}
}

// Process applies the singleflight pattern before passing the request to the next middleware.
func (m *SingleFlightMiddleware) Process(ctx context.Context, opts *RequestOptions, next func(context.Context, *RequestOptions) (*http.Response, error)) (*http.Response, error) {
	m.logger.Debug("Processing request with singleflight middleware")

	// Generate a unique key for the request
	key := fmt.Sprintf("%s:%s:%x", opts.Method, opts.URL, opts.Body)

	// Use singleflight to execute the request
	result, err, _ := m.sfGroup.Do(key, func() (interface{}, error) {
		return next(ctx, opts)
	})
	if err != nil {
		return nil, err
	}

	return result.(*http.Response), nil
}

// SetLogger sets the logger for the middleware.
func (m *SingleFlightMiddleware) SetLogger(l logger.Logger) {
	m.logger = l
}
