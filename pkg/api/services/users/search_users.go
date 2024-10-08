package users

import (
	"context"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/models"
)

// SearchUsers searches for a user with the given username.
// GET https://users.roblox.com/v1/users/search
func (s *Service) SearchUsers(ctx context.Context, params SearchUsersParams) (*models.UserSearchPageResponse, error) {
	if err := s.validate.Struct(params); err != nil {
		return nil, err
	}

	var result models.UserSearchPageResponse
	resp, err := s.client.NewRequest().
		Method(http.MethodGet).
		URL(UsersEndpoint+"/v1/users/search").
		Query("keyword", params.Username).
		Query("limit", strconv.FormatUint(params.Limit, 10)).
		Query("cursor", params.Cursor).
		Result(&result).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return &result, nil
}

// SearchUsersParams holds the parameters for searching users.
type SearchUsersParams struct {
	Username string `json:"username" validate:"required"`           // Required: Username to search for
	Limit    uint64 `json:"limit"    validate:"oneof=10 25 50 100"` // Optional: Maximum number of results to return (default: 10)
	Cursor   string `json:"cursor"   validate:"omitempty,base64"`   // Optional: Cursor for pagination
}

// SearchUsersBuilder is a builder for SearchUsersParams.
type SearchUsersBuilder struct {
	params SearchUsersParams
}

// NewSearchUsersBuilder creates a new SearchUsersBuilder with default values.
func NewSearchUsersBuilder(username string) *SearchUsersBuilder {
	return &SearchUsersBuilder{
		params: SearchUsersParams{
			Username: username,
			Limit:    10,
			Cursor:   "",
		},
	}
}

// WithLimit sets the maximum number of results to return.
func (b *SearchUsersBuilder) WithLimit(limit uint64) *SearchUsersBuilder {
	b.params.Limit = limit
	return b
}

// WithCursor sets the cursor for pagination.
func (b *SearchUsersBuilder) WithCursor(cursor string) *SearchUsersBuilder {
	b.params.Cursor = cursor
	return b
}

// Build returns the SearchUsersParams.
func (b *SearchUsersBuilder) Build() SearchUsersParams {
	return b.params
}
