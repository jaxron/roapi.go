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

// WithLogger sets the logger for the Handler.
func WithLogger(logger logger.Logger) Option {
	return func(h *handler.Handler) {
		h.Logger = logger
	}
}

// WithRateLimit sets the rate limit options for the Handler.
func WithRateLimit(requestsPerSecond float64, burst int) Option {
	return func(h *handler.Handler) {
		h.RateLimitRequestsPerSecond = requestsPerSecond
		h.RateLimitBurst = burst
	}
}

// WithDefaultHeader adds a default header to be sent with all requests.
func WithDefaultHeader(key, value string) Option {
	return func(h *handler.Handler) {
		h.DefaultHeaders[key] = value
	}
}

// WithRetry sets the retry options for failed requests.
func WithRetry(maxAttempts uint64, initialInterval, maxInterval time.Duration) Option {
	return func(h *handler.Handler) {
		h.RetryMaxAttempts = maxAttempts
		h.RetryInitialInterval = initialInterval
		h.RetryMaxInterval = maxInterval
	}
}

// WithRequestTimeout sets the timeout for individual requests.
func WithRequestTimeout(timeout time.Duration) Option {
	return func(h *handler.Handler) {
		h.RequestTimeout = timeout
	}
}

// WithCircuitBreaker sets the circuit breaker options to prevent cascading failures.
func WithCircuitBreaker(maxRequests uint32, interval, timeout time.Duration) Option {
	return func(h *handler.Handler) {
		h.CircuitBreakerMaxRequests = maxRequests
		h.CircuitBreakerInterval = interval
		h.CircuitBreakerTimeout = timeout
	}
}

// WithCircuitBreakerEnabled enables or disables the circuit breaker.
func WithCircuitBreakerEnabled(use bool) Option {
	return func(h *handler.Handler) {
		h.UseCircuitBreaker = use
	}
}

// WithRateLimiterEnabled enables or disables the rate limiter.
func WithRateLimiterEnabled(use bool) Option {
	return func(h *handler.Handler) {
		h.UseRateLimiter = use
	}
}

// WithSingleFlightEnabled enables or disables the single flight feature to deduplicate in-flight requests.
func WithSingleFlightEnabled(use bool) Option {
	return func(h *handler.Handler) {
		h.UseSingleFlight = use
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
