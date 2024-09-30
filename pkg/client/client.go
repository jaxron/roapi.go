// package client provides HTTP request handling functionality with various middleware options.
package client

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	apierrors "github.com/jaxron/roapi.go/pkg/client/errors"
	"github.com/jaxron/roapi.go/pkg/client/logger"
	"go.uber.org/zap"
)

// Default values for various Client settings.
const (
	LogMaxBodyLength = 1024
)

// RequestOptions represents the options for a single HTTP request.
type RequestOptions struct {
	Method    string
	URL       string
	Query     Query
	Headers   map[string]string
	Body      []byte
	Result    interface{}
	UseCookie bool
	UseToken  bool
}

// Middleware interface for all HTTP middleware components.
type Middleware interface {
	Process(ctx context.Context, opts *RequestOptions, next func(context.Context, *RequestOptions) (*http.Response, error)) (*http.Response, error)
	SetLogger(l logger.Logger)
}

// Client manages HTTP requests with various middleware options.
type Client struct {
	middlewares    []Middleware
	httpClient     *http.Client
	ProxyManager   *ProxyManager
	Auth           *Auth
	Logger         logger.Logger
	DefaultHeaders map[string]string
	MaxTimeout     time.Duration
}

// NewClient creates a new Client instance with default settings.
func NewClient(opts ...Option) *Client {
	// Create a new client with default settings
	logger := &logger.NoOpLogger{}
	client := &Client{
		middlewares: []Middleware{},
		httpClient: &http.Client{
			Transport:     http.DefaultTransport,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0, // No client timeout as context timeout is used
		},
		ProxyManager:   NewProxyManager(logger),
		Auth:           nil,
		Logger:         logger,
		DefaultHeaders: make(map[string]string),
		MaxTimeout:     10 * time.Second,
	}
	client.Auth = NewAuth(logger, client.Do)

	// Apply all provided options to customize the client
	for _, opt := range opts {
		opt(client)
	}

	// Set up proxy connection logging
	transport := client.httpClient.Transport.(*http.Transport)
	transport.OnProxyConnectResponse = func(ctx context.Context, proxyURL *url.URL, connectReq *http.Request, connectRes *http.Response) error {
		client.Logger.Debug("Proxy connection established", zap.String("proxy", proxyURL.Host))
		return nil
	}

	return client
}

// Do performs an HTTP request with the specified options.
func (h *Client) Do(ctx context.Context, opts *RequestOptions) (*http.Response, error) {
	// Create a new context with the maximum timeout if the original context has no deadline
	ctx, cancel := context.WithTimeout(ctx, h.MaxTimeout)
	defer cancel()

	return h.executeMiddlewareChain(ctx, opts, 0)
}

// executeMiddlewareChain recursively executes the middleware chain.
func (h *Client) executeMiddlewareChain(ctx context.Context, opts *RequestOptions, index int) (*http.Response, error) {
	if index == len(h.middlewares) {
		// Base case: all middleware processed, execute the actual request
		return h.performRequest(ctx, opts)
	}

	// Execute the current middleware
	return h.middlewares[index].Process(ctx, opts, func(ctx context.Context, opts *RequestOptions) (*http.Response, error) {
		// Move to the next middleware in the chain
		return h.executeMiddlewareChain(ctx, opts, index+1)
	})
}

// performRequest executes the actual HTTP request.
func (h *Client) performRequest(ctx context.Context, opts *RequestOptions) (*http.Response, error) {
	// Check for context cancellation
	if err := ctx.Err(); err != nil {
		return nil, apierrors.NewError(apierrors.ErrorTypeTimeout, "Context error before request", err, nil)
	}

	// Parse and update URL with query parameters
	url, err := url.Parse(opts.URL)
	if err != nil {
		return nil, apierrors.NewError(apierrors.ErrorTypeInternal, "Failed to parse URL", err, nil)
	}
	url.RawQuery = opts.Query.Encode()

	// Prepare headers
	reqHeaders, err := h.prepareHeaders(ctx, opts)
	if err != nil {
		return nil, err
	}

	// Make the HTTP request
	resp, err := h.makeRequest(ctx, opts, url.String(), reqHeaders)
	if err != nil {
		return nil, err
	}

	// Process the response
	return h.processResponse(resp, opts.Result)
}

// makeRequest creates and sends an HTTP request.
func (h *Client) makeRequest(ctx context.Context, opts *RequestOptions, url string, headers map[string]string) (*http.Response, error) {
	// Log the request details
	h.logRequest(opts.Method, url, headers, string(opts.Body))

	// Set up proxy if available
	if h.ProxyManager.GetProxyCount() > 0 {
		transport := h.httpClient.Transport.(*http.Transport).Clone()
		transport.Proxy = h.ProxyManager.NextProxy

		h.httpClient.Transport = transport
	}

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, opts.Method, url, strings.NewReader(string(opts.Body)))
	if err != nil {
		return nil, apierrors.NewError(apierrors.ErrorTypeInternal, "Failed to create request", err, nil)
	}

	// Set request headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Send the request
	resp, err := h.httpClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, apierrors.NewError(apierrors.ErrorTypeTimeout, "Request timed out", err, nil)
		}
		return nil, apierrors.NewError(apierrors.ErrorTypeNetwork, "Network error occurred", err, nil)
	}

	return resp, nil
}

// prepareHeaders combines default headers, authentication headers, and request-specific headers.
func (h *Client) prepareHeaders(ctx context.Context, opts *RequestOptions) (map[string]string, error) {
	result := make(map[string]string)

	// Set default Content-Type and Accept headers
	for k, v := range map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	} {
		result[k] = v
	}

	// Add client's default headers
	for k, v := range h.DefaultHeaders {
		result[k] = v
	}

	// Add authentication headers if cookies are available
	if h.Auth.GetCookieCount() > 0 {
		authHeaders, err := h.Auth.GetAuthHeaders(ctx, opts.UseCookie, opts.UseToken)
		if err != nil {
			return nil, apierrors.NewError(apierrors.ErrorTypeAuth, "Failed to get auth headers", err, nil)
		}

		for k, v := range authHeaders {
			result[k] = v
		}
	}

	// Add request-specific headers, overriding any existing headers
	for k, v := range opts.Headers {
		result[k] = v
	}

	return result, nil
}

// processResponse handles the HTTP response, including error checking and JSON unmarshaling.
func (h *Client) processResponse(resp *http.Response, result interface{}) (*http.Response, error) {
	defer resp.Body.Close()

	// Read the entire body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, apierrors.NewError(apierrors.ErrorTypeHTTP, "Failed to read response body", err, nil)
	}

	// Log the response
	h.logResponse(resp.StatusCode, resp.Header, string(body))

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		// Handle rate limiting
		if resp.StatusCode == http.StatusTooManyRequests {
			return nil, apierrors.NewError(apierrors.ErrorTypeTooManyRequests, "Rate limited by Roblox API", nil, body)
		}
		// Try to parse API errors
		apiErrors, parseErr := apierrors.ParseAPIErrors(body)
		if parseErr == nil && len(apiErrors.Errors) > 0 {
			return nil, apierrors.NewError(apierrors.ErrorTypeAPI, apiErrors.Errors[0].Message, nil, body)
		}
		// Generic HTTP error
		return nil, apierrors.NewError(apierrors.ErrorTypeHTTP, "Unexpected status code", nil, body)
	}

	// Unmarshal JSON response if result interface is provided
	if result != nil {
		if err := sonic.ConfigFastest.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
			return nil, apierrors.NewError(apierrors.ErrorTypeUnmarshal, "Failed to unmarshal response", err, body)
		}
	}

	// Create a new response with the same data but with a new body reader
	newResp := *resp
	newResp.Body = io.NopCloser(bytes.NewReader(body))

	return &newResp, nil
}

// SetLogger updates the client's logger and propagates it to all middleware.
func (h *Client) SetLogger(l logger.Logger) {
	// Update all middleware loggers
	for _, m := range h.middlewares {
		m.SetLogger(l)
	}
	h.Logger = l
	h.ProxyManager.logger = l
}

// UpdateMiddleware adds or replaces a middleware in the client's middleware chain.
func (h *Client) UpdateMiddleware(newMiddleware Middleware) {
	for i, m := range h.middlewares {
		if reflect.TypeOf(m) == reflect.TypeOf(newMiddleware) {
			h.middlewares[i] = newMiddleware
			m.SetLogger(h.Logger)
			return
		}
	}
	h.middlewares = append(h.middlewares, newMiddleware)
	newMiddleware.SetLogger(h.Logger)
}

// logRequest logs the details of an outgoing HTTP request.
func (h *Client) logRequest(method, url string, headers map[string]string, body string) {
	// Truncate body if it's too long
	logBody := body
	if len(logBody) > LogMaxBodyLength {
		logBody = logBody[:LogMaxBodyLength] + "...TRUNCATED"
	}

	h.Logger.Debug("Request",
		zap.String("method", method),
		zap.String("url", url),
		zap.Any("len_headers", len(headers)),
		zap.String("body", logBody),
	)
}

// logResponse logs the details of an incoming HTTP response.
func (h *Client) logResponse(statusCode int, headers http.Header, body string) {
	// Truncate body if it's too long
	logBody := body
	if len(logBody) > LogMaxBodyLength {
		logBody = logBody[:LogMaxBodyLength] + "...TRUNCATED"
	}

	h.Logger.Debug("Response",
		zap.Int("status_code", statusCode),
		zap.Any("headers", headers),
		zap.String("body", logBody),
	)
}
