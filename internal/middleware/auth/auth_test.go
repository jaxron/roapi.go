package auth_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jaxron/axonet/pkg/client/logger"
	"github.com/jaxron/roapi.go/internal/middleware/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockTransport is a custom http.RoundTripper for mocking HTTP responses
type mockTransport struct {
	roundTripFunc func(*http.Request) (*http.Response, error)
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.roundTripFunc(req)
}

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
				roundTripFunc: func(*http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"X-Csrf-Token": []string{"mocked-csrf-token"}},
					}, nil
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

func TestCSRFTokenCaching(t *testing.T) {
	t.Run("CSRF token caching and refreshing", func(t *testing.T) {
		t.Parallel()

		cookies := []string{"cookie1", "cookie2"}
		middleware := auth.New(cookies)
		middleware.SetLogger(logger.NewBasicLogger())

		mockTime := time.Now()
		middleware.SetNowFunc(func() time.Time {
			return mockTime
		})

		tokenCount := 0
		mockTransport := &mockTransport{
			roundTripFunc: func(*http.Request) (*http.Response, error) {
				tokenCount++
				return &http.Response{
					StatusCode: http.StatusOK,
					Header:     http.Header{"X-Csrf-Token": []string{fmt.Sprintf("mocked-csrf-token-%d", tokenCount)}},
				}, nil
			},
		}

		mockClient := &http.Client{Transport: mockTransport}

		ctx := context.WithValue(context.Background(), auth.KeyAddCookie, true)
		ctx = context.WithValue(ctx, auth.KeyAddToken, true)

		// First request should fetch a new token
		req1 := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
		_, err := middleware.Process(ctx, mockClient, req1, func(ctx context.Context, httpClient *http.Client, req *http.Request) (*http.Response, error) {
			assert.Equal(t, "mocked-csrf-token-1", req.Header.Get("X-Csrf-Token"))
			return &http.Response{StatusCode: http.StatusOK}, nil
		})
		require.NoError(t, err)

		// Second request within 5 minutes should use the cached token
		req2 := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
		_, err = middleware.Process(ctx, mockClient, req2, func(ctx context.Context, httpClient *http.Client, req *http.Request) (*http.Response, error) {
			assert.Equal(t, "mocked-csrf-token-1", req.Header.Get("X-Csrf-Token"))
			return &http.Response{StatusCode: http.StatusOK}, nil
		})
		require.NoError(t, err)

		// Simulate passage of 5 minutes
		mockTime = mockTime.Add(5 * time.Minute)

		// Third request after 5 minutes should fetch a new token
		req3 := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
		_, err = middleware.Process(ctx, mockClient, req3, func(ctx context.Context, httpClient *http.Client, req *http.Request) (*http.Response, error) {
			assert.Equal(t, "mocked-csrf-token-2", req.Header.Get("X-Csrf-Token"))
			return &http.Response{StatusCode: http.StatusOK}, nil
		})
		require.NoError(t, err)

		// Verify that only two tokens were requested
		assert.Equal(t, 2, tokenCount)
	})
}
