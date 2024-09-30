package client

import (
	"net/url"
	"time"

	"github.com/jaxron/roapi.go/internal/handler"
	"github.com/jaxron/roapi.go/pkg/logger"
)

// Option is a function type that modifies the Handler configuration.
type Option func(*handler.Handler)

// WithCookies sets the .ROBLOSECURITY cookie values for authentication.
func WithCookies(cookies []string) Option {
	return func(h *handler.Handler) {
		h.Auth.UpdateCookies(cookies)
	}
}

// WithProxies sets the list of proxy URLs to use for requests.
func WithProxies(proxies []*url.URL) Option {
	return func(h *handler.Handler) {
		h.ProxyManager.UpdateProxies(proxies)
	}
}

// WithDefaultHeader adds a default header to be sent with all requests.
func WithDefaultHeader(key, value string) Option {
	return func(h *handler.Handler) {
		h.DefaultHeaders[key] = value
	}
}

// WithTimeout sets the maximum timeout for all requests.
func WithTimeout(timeout time.Duration) Option {
	return func(h *handler.Handler) {
		h.MaxTimeout = timeout
	}
}

// WithMiddleware adds or updates the middleware for the Handler.
func WithMiddleware(middleware handler.Middleware) Option {
	return func(h *handler.Handler) {
		h.UpdateMiddleware(middleware)
	}
}

// WithRateLimit enables the rate limiter middleware with the specified options.
func WithRateLimit(requestsPerSecond float64, burst int) Option {
	return WithMiddleware(handler.NewRateLimiterMiddleware(requestsPerSecond, burst))
}

// WithRetry enables the retry middleware with the specified options.
func WithRetry(maxAttempts uint64, initialInterval, maxInterval time.Duration) Option {
	return WithMiddleware(handler.NewRetryMiddleware(maxAttempts, initialInterval, maxInterval))
}

// WithCircuitBreaker enables the circuit breaker middleware with the specified options.
func WithCircuitBreaker(maxRequests uint32, interval, timeout time.Duration) Option {
	return WithMiddleware(handler.NewCircuitBreakerMiddleware(maxRequests, interval, timeout))
}

// WithSingleFlight enables the single flight middleware.
func WithSingleFlight() Option {
	return WithMiddleware(handler.NewSingleFlightMiddleware())
}

// WithLogger sets the logger for the Handler and its middleware.
func WithLogger(logger logger.Logger) Option {
	return func(h *handler.Handler) {
		h.SetLogger(logger)
	}
}

// Request helps build requests using method chaining.
type Request struct {
	opts *handler.RequestOptions
}

// NewRequest creates a new Request with default options.
func NewRequest() *Request {
	return &Request{
		opts: &handler.RequestOptions{
			Method:    "",
			URL:       "",
			Query:     make(handler.Query),
			Headers:   make(map[string]string),
			Body:      nil,
			Result:    nil,
			UseCookie: false,
			UseToken:  false,
		},
	}
}

// Method sets the HTTP method for the request.
func (rb *Request) Method(method string) *Request {
	rb.opts.Method = method
	return rb
}

// URL sets the URL for the request.
func (rb *Request) URL(url string) *Request {
	rb.opts.URL = url
	return rb
}

// Query adds a query parameter to the request.
func (rb *Request) Query(key, value string) *Request {
	rb.opts.Query.Add(key, value)
	return rb
}

// Header adds a header to the request.
func (rb *Request) Header(key, value string) *Request {
	rb.opts.Headers[key] = value
	return rb
}

// Body sets the body of the request.
func (rb *Request) Body(body []byte) *Request {
	rb.opts.Body = body
	return rb
}

// JSONBody sets the body of the request as JSON.
func (rb *Request) JSONBody(fn func() ([]byte, error)) (*Request, error) {
	// Generate and set JSON body using the provided function
	body, err := fn()
	if err != nil {
		return nil, err
	}

	rb.opts.Body = body
	return rb, nil
}

// Result sets the result pointer for JSON unmarshaling of the response.
func (rb *Request) Result(v interface{}) *Request {
	rb.opts.Result = v
	return rb
}

// UseCookie sets whether to use the .ROBLOSECURITY cookie for authentication.
func (rb *Request) UseCookie(use bool) *Request {
	rb.opts.UseCookie = use
	return rb
}

// UseToken sets whether to use the X-CSRF-TOKEN for CSRF protection.
func (rb *Request) UseToken(use bool) *Request {
	rb.opts.UseToken = use
	return rb
}

// Build returns the final RequestOptions for execution.
func (rb *Request) Build() *handler.RequestOptions {
	return rb.opts
}
