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
func (s *Service) SearchUsers(ctx context.Context, b *SearchUsersBuilder) (*models.UserSearchPageResponse, error) {
	var result models.UserSearchPageResponse
	resp, err := s.client.NewRequest().
		Method(http.MethodGet).
		URL(UsersEndpoint+"/v1/users/search").
		Query("keyword", b.username).
		Query("limit", strconv.FormatUint(b.limit, 10)).
		Query("cursor", b.cursor).
		Result(&result).
		JSONHeaders().
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return &result, nil
}

// SearchUsersBuilder builds parameters for SearchUsers API call.
type SearchUsersBuilder struct {
	username string // Required: Username to search for
	limit    uint64 // Optional: Maximum number of results to return (default: 10)
	cursor   string // Optional: Cursor for pagination
}

// NewSearchUsersBuilder creates a new SearchUsersBuilder with the given username.
func NewSearchUsersBuilder(username string) *SearchUsersBuilder {
	return &SearchUsersBuilder{
		username: username,
		limit:    DefaultLimit,
		cursor:   "",
	}
}

// Limit sets the maximum number of results to return.
func (b *SearchUsersBuilder) Limit(limit uint64) *SearchUsersBuilder {
	b.limit = limit
	return b
}

// Cursor sets the cursor for pagination.
func (b *SearchUsersBuilder) Cursor(cursor string) *SearchUsersBuilder {
	b.cursor = cursor
	return b
}
