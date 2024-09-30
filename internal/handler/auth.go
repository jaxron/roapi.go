package handler

import (
	"context"
	"fmt"
	"math/rand/v2"
	"net/http"
	"sync"

	apierrors "github.com/jaxron/roapi.go/pkg/errors"
	"github.com/jaxron/roapi.go/pkg/logger"
	"go.uber.org/zap"
)

// HTTPRequester is a function type that performs HTTP requests.
type HTTPRequester func(ctx context.Context, opts *RequestOptions) (*http.Response, error)

// Auth represents the authentication information for Roblox users.
type Auth struct {
	cookies   []string
	current   int
	doRequest HTTPRequester
	logger    logger.Logger
	mu        sync.RWMutex
}

// NewAuth creates a new Auth instance with the specified Handler.
func NewAuth(logger logger.Logger, doRequest HTTPRequester) *Auth {
	return &Auth{
		cookies:   []string{},
		current:   0,
		doRequest: doRequest,
		logger:    logger,
		mu:        sync.RWMutex{},
	}
}

// GetAuthHeaders returns a map of authentication headers based on the specified options.
func (a *Auth) GetAuthHeaders(ctx context.Context, useCookie bool, useToken bool) (map[string]string, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	headers := make(map[string]string)

	// Include the .ROBLOSECURITY cookie if requested
	if useCookie {
		headers["Cookie"] = ".ROBLOSECURITY=" + a.nextCookie()
		a.logger.Debug("Using .ROBLOSECURITY cookie")
	}

	// Include the CSRF token if requested
	if useToken {
		// Get a new CSRF token
		token, err := a.getCSRFToken(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get CSRF token: %w", err)
		}

		headers["X-CSRF-TOKEN"] = token
		a.logger.Debug("Using CSRF token")
	}

	return headers, nil
}

// nextCookie returns the next .ROBLOSECURITY cookie for each request using round-robin selection.
func (a *Auth) nextCookie() string {
	// Select the current cookie
	cookie := a.cookies[a.current]

	// Move to the next cookie for the next request (round-robin)
	a.current = (a.current + 1) % len(a.cookies)

	// Log the first 50 characters of the cookie for debugging
	a.logger.Debug("Next Cookie", zap.String("cookie", cookie[:50]))

	return cookie
}

// getCSRFToken sends a POST request to generate a new CSRF token.
func (a *Auth) getCSRFToken(ctx context.Context) (string, error) {
	// Prepare the request options
	options := &RequestOptions{
		Method:    http.MethodPost,
		URL:       "https://auth.roblox.com/v2/logout",
		Query:     make(Query),
		Headers:   make(map[string]string),
		Body:      nil,
		Result:    nil,
		UseCookie: true,
		UseToken:  false,
	}

	// Send the request to get a new CSRF token
	resp, err := a.doRequest(ctx, options)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Extract the CSRF token from the response headers
	csrfToken := resp.Header.Get("X-Csrf-Token")
	if csrfToken == "" {
		return "", apierrors.ErrCSRFTokenNotFound
	}

	return csrfToken, nil
}

// UpdateCookies updates the list of cookies at runtime.
func (a *Auth) UpdateCookies(cookies []string) {
	a.mu.Lock()
	defer a.mu.Unlock()

	// Replace the existing cookie list with the new one
	a.cookies = cookies
	if len(cookies) > 0 {
		// Randomize the starting cookie to distribute load
		a.current = rand.IntN(len(cookies))
	}

	a.logger.Debug("Cookie list updated", zap.Int("cookie_count", len(cookies)))
}

// GetCookieCount returns the current number of cookies in the list.
func (a *Auth) GetCookieCount() int {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return len(a.cookies)
}
