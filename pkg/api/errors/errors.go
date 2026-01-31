package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	clientErrors "github.com/jaxron/axonet/pkg/client/errs"
)

var (
	ErrNoMessage = errors.New("no error message available")
	ErrReadBody  = errors.New("failed to read response body")
	ErrParseJSON = errors.New("invalid JSON response")

	ErrInvalidRequest  = errors.New("invalid request")
	ErrInvalidResponse = errors.New("invalid response")
)

// APIError represents a list of errors returned by the Roblox API.
type APIError struct {
	Errors []APIErrorData `json:"errors"`
}

// APIErrorData represents a single error returned by the Roblox API.
type APIErrorData struct {
	Code              int    `json:"code"`
	Message           string `json:"message"`
	UserFacingMessage string `json:"userFacingMessage"`
}

// Error implements the error interface for APIErrors.
func (ae *APIError) Error() string {
	if len(ae.Errors) == 0 {
		return "roblox API error: " + ErrNoMessage.Error()
	}

	err := ae.Errors[0]

	return fmt.Sprintf("roblox API error (%d): %s", err.Code, err.Message)
}

// HandleAPIError checks if the error is a bad status error and parses the API error if so.
// It returns the original error if it's not a bad status error.
func HandleAPIError(resp *http.Response, err error) error {
	if errors.Is(err, clientErrors.ErrBadStatus) {
		return New(resp)
	}

	return err
}

// New parses the response body into an error type.
func New(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%w: code %d", ErrReadBody, resp.StatusCode)
	}

	var apiErrors APIError
	if err := json.Unmarshal(body, &apiErrors); err != nil {
		return fmt.Errorf("%w: code %d", ErrParseJSON, resp.StatusCode)
	}

	if len(apiErrors.Errors) == 0 {
		return fmt.Errorf("roblox API error (%d): %w", resp.StatusCode, ErrNoMessage)
	}

	return &apiErrors
}
