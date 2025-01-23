package jsonheader

import (
	"context"
	"net/http"

	"github.com/jaxron/axonet/pkg/client/logger"
	"github.com/jaxron/axonet/pkg/client/middleware"
)

// JSONHeaderMiddleware adds application/json headers to HTTP requests.
type JSONHeaderMiddleware struct {
	logger logger.Logger
}

// New creates a new JSONHeaderMiddleware instance.
func New() *JSONHeaderMiddleware {
	return &JSONHeaderMiddleware{
		logger: &logger.NoOpLogger{},
	}
}

// Process applies headers to the request before passing it to the next middleware.
func (m *JSONHeaderMiddleware) Process(ctx context.Context, httpClient *http.Client, req *http.Request, next middleware.NextFunc) (*http.Response, error) {
	m.logger.Debug("Adding JSON headers to request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return next(ctx, httpClient, req)
}

// SetLogger sets the logger for the middleware.
func (m *JSONHeaderMiddleware) SetLogger(l logger.Logger) {
	m.logger = l
}
