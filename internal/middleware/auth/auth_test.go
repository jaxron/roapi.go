package auth_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jaxron/axonet/pkg/client/logger"
	"github.com/jaxron/roapi.go/internal/middleware/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthMiddleware(t *testing.T) {
	t.Run("Apply cookie and CSRF token to request", func(t *testing.T) {
		t.Parallel()

		cookies := []string{"cookie1", "cookie2"}
		middleware := auth.New(cookies)
		middleware.SetLogger(logger.NewBasicLogger())

		req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
		ctx := context.WithValue(context.Background(), auth.KeyAddCookie, true)
		ctx = context.WithValue(ctx, auth.KeyAddToken, true)

		// Mock HTTP client to return a CSRF token
		mockClient := &http.Client{
			Transport: &mockTransport{
				response: &http.Response{
					StatusCode: http.StatusOK,
					Header:     http.Header{"X-Csrf-Token": []string{"mocked-csrf-token"}},
				},
			},
		}

		resp, err := middleware.Process(ctx, mockClient, req, func(ctx context.Context, httpClient *http.Client, req *http.Request) (*http.Response, error) {
			assert.Contains(t, req.Header.Get("Cookie"), ".ROBLOSECURITY=cookie")
			assert.Equal(t, "mocked-csrf-token", req.Header.Get("X-Csrf-Token"))
			return &http.Response{StatusCode: http.StatusOK}, nil
		})

		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Skip middleware when cookies and tokens are disabled", func(t *testing.T) {
		t.Parallel()

		cookies := []string{"cookie1", "cookie2"}
		middleware := auth.New(cookies)
		middleware.SetLogger(logger.NewBasicLogger())

		req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
		ctx := context.Background()

		resp, err := middleware.Process(ctx, &http.Client{}, req, func(ctx context.Context, httpClient *http.Client, req *http.Request) (*http.Response, error) {
			assert.Empty(t, req.Header.Get("Cookie"))
			assert.Empty(t, req.Header.Get("X-Csrf-Token"))
			return &http.Response{StatusCode: http.StatusOK}, nil
		})

		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Update cookies at runtime", func(t *testing.T) {
		t.Parallel()

		cookies := []string{"cookie1", "cookie2"}
		middleware := auth.New(cookies)
		middleware.SetLogger(logger.NewBasicLogger())

		assert.Equal(t, 2, middleware.GetCookieCount())

		newCookies := []string{"cookie3", "cookie4", "cookie5"}
		middleware.UpdateCookies(newCookies)

		assert.Equal(t, 3, middleware.GetCookieCount())
	})
}

// mockTransport is a custom http.RoundTripper for mocking HTTP responses
type mockTransport struct {
	response *http.Response
}

func (m *mockTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return m.response, nil
}
