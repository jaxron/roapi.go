package errors

import (
	"errors"
	"fmt"

	"github.com/bytedance/sonic"
)

var (
	ErrCSRFTokenNotFound = errors.New("CSRF token not found in response headers")
	ErrNoCookie          = errors.New("no .ROBLOSECURITY cookie found")
)

// ErrorType represents the type of error encountered.
type ErrorType int

const (
	ErrorTypeNetwork ErrorType = iota
	ErrorTypeTimeout
	ErrorTypeAPI
	ErrorTypeHTTP
	ErrorTypeAuth
	ErrorTypeRateLimit
	ErrorTypeTooManyRequests
	ErrorTypeUnmarshal
	ErrorTypeInternal
	ErrorTypeCircuitOpen
	ErrorTypeCircuitExhausted
)

// Error represents a custom error type for the client package.
type Error struct {
	Type    ErrorType // The type of error
	Message string    // A human-readable error message
	Cause   error     // The underlying cause of the error
	Body    []byte    // The raw response body, if applicable
}

// NewError creates a new Error with the given parameters.
func NewError(errType ErrorType, message string, cause error, body []byte) *Error {
	return &Error{
		Type:    errType,
		Message: message,
		Cause:   cause,
		Body:    body,
	}
}

// Error returns the error message, including the cause if present.
func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

// Unwrap returns the underlying cause of the error.
func (e *Error) Unwrap() error {
	return e.Cause
}

// APIErrors represents a list of errors returned by the Roblox API.
type APIErrors struct {
	Errors []APIError `json:"errors"`
}

// APIError represents a single error returned by the Roblox API.
type APIError struct {
	Code              int    `json:"code"`              // The error code
	Message           string `json:"message"`           // The error message
	UserFacingMessage string `json:"userFacingMessage"` // A user-friendly error message
}

// ParseAPIErrors parses the response body into APIErrors.
func ParseAPIErrors(body []byte) (*APIErrors, error) {
	var apiErrors APIErrors
	err := sonic.Unmarshal(body, &apiErrors)
	if err != nil {
		return nil, err
	}
	return &apiErrors, nil
}

// IsTemporary returns true if the error is temporary and the request can be retried.
func IsTemporary(err error) bool {
	var clientErr *Error
	if errors.As(err, &clientErr) {
		// Check if the error type is one that can be retried
		return clientErr.Type == ErrorTypeNetwork ||
			clientErr.Type == ErrorTypeTimeout ||
			clientErr.Type == ErrorTypeRateLimit ||
			clientErr.Type == ErrorTypeTooManyRequests
	}
	return false
}
