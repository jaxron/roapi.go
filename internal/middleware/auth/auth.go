package auth

import (
	"context"
	"errors"
	"net/http"
	"sync/atomic"

	"github.com/jaxron/axonet/pkg/client/logger"
	"github.com/jaxron/axonet/pkg/client/middleware"
)

type contextKey int

const (
	KeyAddCookie contextKey = iota
	KeyAddToken
)

var (
	ErrNoCookie      = errors.New("no cookie available")
	ErrTokenNotFound = errors.New("CSRF token not found")
)

// AuthMiddleware manages cookie rotation for HTTP requests.
type AuthMiddleware struct {
	cookies atomic.Value
	current atomic.Uint64
	logger  logger.Logger
}

type cookieState struct {
	cookies []string
}

// New creates a new AuthMiddleware instance.
func New(cookies []string) *AuthMiddleware {
	m := &AuthMiddleware{
		cookies: atomic.Value{},
		current: atomic.Uint64{},
		logger:  &logger.NoOpLogger{},
	}
	m.cookies.Store(&cookieState{cookies: cookies})
	return m
}

// Process applies cookie logic before passing the request to the next middleware.
func (m *AuthMiddleware) Process(ctx context.Context, httpClient *http.Client, req *http.Request, next middleware.NextFunc) (*http.Response, error) {
	isCookieEnabled, cookieOk := ctx.Value(KeyAddCookie).(bool)
	isTokenEnabled, tokenOk := ctx.Value(KeyAddToken).(bool)

	// Skip middleware if cookies and tokens are disabled
	if !((cookieOk && isCookieEnabled) || (tokenOk && isTokenEnabled)) {
		return next(ctx, httpClient, req)
	}

	// Apply cookie and token to the request if required
	cookie, err := m.getAndValidateCookie()
	if err != nil {
		return nil, err
	}

	if err := m.applyCookieAndToken(ctx, httpClient, req, cookie, isCookieEnabled, isTokenEnabled); err != nil {
		return nil, err
	}

	return next(ctx, httpClient, req)
}

func (m *AuthMiddleware) getAndValidateCookie() (string, error) {
	state := m.cookies.Load().(*cookieState)
	if len(state.cookies) == 0 {
		return "", ErrNoCookie
	}

	m.logger.Debug("Processing request with cookie middleware")

	current := m.current.Add(1) - 1
	index := int(current % uint64(len(state.cookies))) // #nosec G115
	return state.cookies[index], nil
}

func (m *AuthMiddleware) applyCookieAndToken(ctx context.Context, httpClient *http.Client, req *http.Request, cookie string, isCookieEnabled, isTokenEnabled bool) error {
	if isCookieEnabled {
		req.Header.Add("Cookie", ".ROBLOSECURITY="+cookie)
		m.logger.Debug("Applied cookie to request")
	}

	if isTokenEnabled {
		token, err := m.getCSRFToken(ctx, httpClient, cookie)
		if err != nil {
			return err
		}
		req.Header.Set("X-Csrf-Token", token)
		m.logger.Debug("Applied CSRF token to request")
	}

	return nil
}

// getCSRFToken sends a POST request to generate a new CSRF token.
func (m *AuthMiddleware) getCSRFToken(ctx context.Context, httpClient *http.Client, cookie string) (string, error) {
	// Create a new request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://auth.roblox.com/v2/logout", nil)
	if err != nil {
		return "", err
	}

	// Add the cookie to the request
	req.Header.Add("Cookie", ".ROBLOSECURITY="+cookie)

	// Send the request
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Get the CSRF token from the response
	csrfToken := resp.Header.Get("X-Csrf-Token")
	if csrfToken == "" {
		return "", ErrTokenNotFound
	}

	return csrfToken, nil
}

// UpdateCookies updates the list of cookies at runtime.
func (m *AuthMiddleware) UpdateCookies(cookies []string) {
	newState := &cookieState{cookies: cookies}
	m.cookies.Store(newState)
	m.current.Store(0)

	m.logger.WithFields(logger.Int("cookies", len(cookies))).Debug("Cookies updated")
}

// GetCookieCount returns the current number of cookies in the list.
func (m *AuthMiddleware) GetCookieCount() int {
	state := m.cookies.Load().(*cookieState)
	return len(state.cookies)
}

// SetLogger sets the logger for the middleware.
func (m *AuthMiddleware) SetLogger(l logger.Logger) {
	m.logger = l
}
