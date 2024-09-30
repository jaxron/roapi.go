package client_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDo tests the basic functionality of the Client's Do method.
func TestDo(t *testing.T) {
	t.Parallel()

	// Set up a mock server to respond to requests
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"message": "success"}`))
		assert.NoError(t, err)
	}))
	defer mockServer.Close()

	// Create a new test client
	c := utils.NewTestClient(false, false)

	// Prepare the request
	ctx := context.Background()
	reqOptions := client.NewRequest().
		Method(http.MethodGet).
		URL(mockServer.URL).
		Build()

	// Execute the request
	resp, err := c.Do(ctx, reqOptions)

	// Assert the results
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// TestDoWithRetry tests the retry functionality of the Client's Do method.
func TestDoWithRetry(t *testing.T) {
	t.Parallel()

	attempts := 0
	// Set up a mock server that fails twice before succeeding
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		attempts++
		if attempts < 3 {
			w.WriteHeader(http.StatusTooManyRequests)
		} else {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{"message": "success after retry"}`))
			assert.NoError(t, err)
		}
	}))
	defer mockServer.Close()

	// Create a new test client with retry configuration
	c := utils.NewTestClient(false, false, client.WithRetry(3, 10*time.Millisecond, 100*time.Millisecond))

	// Prepare the request
	ctx := context.Background()
	reqOptions := client.NewRequest().
		Method(http.MethodGet).
		URL(mockServer.URL).
		Build()

	// Execute the request
	resp, err := c.Do(ctx, reqOptions)

	// Assert the results
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, 3, attempts) // Ensure it took 3 attempts to succeed
}
