package jsonheader_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jaxron/axonet/pkg/client/logger"
	"github.com/jaxron/roapi.go/internal/middleware/jsonheader"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSONHeaderMiddleware(t *testing.T) {
	t.Run("Apply JSON headers to request", func(t *testing.T) {
		t.Parallel()

		middleware := jsonheader.New()
		middleware.SetLogger(logger.NewBasicLogger())

		req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)

		resp, err := middleware.Process(context.Background(), &http.Client{}, req, func(ctx context.Context, httpClient *http.Client, req *http.Request) (*http.Response, error) {
			assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
			assert.Equal(t, "application/json", req.Header.Get("Accept"))
			return &http.Response{StatusCode: http.StatusOK}, nil
		})

		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
