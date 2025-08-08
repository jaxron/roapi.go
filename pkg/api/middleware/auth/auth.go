package auth

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

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

// Middleware manages cookie rotation and CSRF token caching for HTTP requests.
type Middleware struct {
	cookies      []string
	cookieCount  int
	cookiesMux   sync.RWMutex
	current      atomic.Uint64
	csrfToken    string
	csrfTokenExp time.Time
	csrfTokenMux sync.RWMutex
	logger       logger.Logger
	now          func() time.Time
}

// New creates a new AuthMiddleware instance.
func New(cookies []string) *Middleware {
	m := &Middleware{
		cookies:      cookies,
		current:      atomic.Uint64{},
		cookieCount:  len(cookies),
		cookiesMux:   sync.RWMutex{},
		csrfToken:    "",
		csrfTokenExp: time.Time{},
		csrfTokenMux: sync.RWMutex{},
		logger:       &logger.NoOpLogger{},
		now:          time.Now,
	}
	m.current.Store(0)
	return m
}

// Process applies cookie logic before passing the request to the next middleware.
func (m *Middleware) Process(ctx context.Context, httpClient *http.Client, req *http.Request, next middleware.NextFunc) (*http.Response, error) {
	isCookieEnabled, cookieOk := ctx.Value(KeyAddCookie).(bool)
	isTokenEnabled, tokenOk := ctx.Value(KeyAddToken).(bool)

	// Skip middleware if cookies and tokens are disabled
	if (!cookieOk || !isCookieEnabled) && (!tokenOk || !isTokenEnabled) {
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

// UpdateCookies updates the list of cookies at runtime.
func (m *Middleware) UpdateCookies(cookies []string) {
	m.cookiesMux.Lock()
	defer m.cookiesMux.Unlock()

	m.cookies = cookies
	m.current.Store(0)
	m.cookieCount = len(cookies)
	m.logger.WithFields(logger.Int("cookies", len(cookies))).Debug("Cookies updated")
}

// Shuffle randomizes the order of the cookies.
func (m *Middleware) Shuffle() {
	m.cookiesMux.Lock()
	defer m.cookiesMux.Unlock()

	rand.New(rand.NewSource(time.Now().UnixNano())).Shuffle(len(m.cookies), func(i, j int) {
		m.cookies[i], m.cookies[j] = m.cookies[j], m.cookies[i]
	})

	m.logger.Debug("Cookies shuffled")
}

// GetCookieCount returns the current number of cookies in the list.
func (m *Middleware) GetCookieCount() int {
	m.cookiesMux.RLock()
	defer m.cookiesMux.RUnlock()
	return m.cookieCount
}

// SetLogger sets the logger for the middleware.
func (m *Middleware) SetLogger(l logger.Logger) {
	m.logger = l
}

// SetNowFunc sets a custom function for getting the current time (useful for testing).
func (m *Middleware) SetNowFunc(f func() time.Time) {
	m.now = f
}

func (m *Middleware) getAndValidateCookie() (string, error) {
	m.cookiesMux.RLock()
	defer m.cookiesMux.RUnlock()

	if m.cookieCount == 0 {
		return "", ErrNoCookie
	}

	m.logger.Debug("Processing request with cookie middleware")

	current := m.current.Add(1) - 1
	index := current % uint64(m.cookieCount) // #nosec G115
	return m.cookies[index], nil
}

func (m *Middleware) applyCookieAndToken(ctx context.Context, httpClient *http.Client, req *http.Request, cookie string, isCookieEnabled, isTokenEnabled bool) error {
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

// getCSRFToken retrieves a valid CSRF token, either from cache or by making a new request.
func (m *Middleware) getCSRFToken(ctx context.Context, httpClient *http.Client, cookie string) (string, error) {
	m.csrfTokenMux.RLock()
	csrfToken := m.csrfToken
	csrfTokenExp := m.csrfTokenExp
	m.csrfTokenMux.RUnlock()

	// Return cached token if it's still valid
	if csrfToken != "" && m.now().Before(csrfTokenExp) {
		return csrfToken, nil
	}

	// Otherwise, refresh the token
	token, err := m.refreshCSRFToken(ctx, httpClient, cookie)
	if err != nil {
		return "", err
	}

	return token, nil
}

// refreshCSRFToken sends a POST request to generate a new CSRF token and caches it.
func (m *Middleware) refreshCSRFToken(ctx context.Context, httpClient *http.Client, cookie string) (string, error) {
	m.csrfTokenMux.Lock()
	defer m.csrfTokenMux.Unlock()

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
	defer func() { _ = resp.Body.Close() }()

	// Get the CSRF token from the response
	csrfToken := resp.Header.Get("X-Csrf-Token")
	if csrfToken == "" {
		return "", ErrTokenNotFound
	}

	// Cache the new token
	m.csrfToken = csrfToken
	m.csrfTokenExp = m.now().Add(5 * time.Minute)

	return csrfToken, nil
}
