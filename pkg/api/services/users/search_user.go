package users

import (
	"context"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/models"
	"github.com/jaxron/roapi.go/pkg/client"
)

// SearchUser searches for a user with the given username.
// GET https://users.roblox.com/v1/users/search
func (s *Service) SearchUser(ctx context.Context, b *SearchUserBuilder) (*models.SearchResult, error) {
	var result models.SearchResult
	req := client.NewRequest().
		Method(http.MethodGet).
		URL(UsersEndpoint+"/v1/users/search").
		Query("keyword", b.username).
		Query("limit", strconv.FormatUint(b.limit, 10)).
		Query("cursor", b.cursor).
		Result(&result)

	resp, err := s.Client.Do(ctx, req.Build())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &result, nil
}

// SearchUserBuilder builds parameters for SearchUser API call.
type SearchUserBuilder struct {
	username string // Required: Username to search for
	limit    uint64 // Optional: Maximum number of results to return (default: 10)
	cursor   string // Optional: Cursor for pagination
}

// NewSearchUserBuilder creates a new SearchUserBuilder with the given username.
func NewSearchUserBuilder(username string) *SearchUserBuilder {
	return &SearchUserBuilder{
		username: username,
		limit:    DefaultLimit,
		cursor:   "",
	}
}

// Limit sets the maximum number of results to return.
func (b *SearchUserBuilder) Limit(limit uint64) *SearchUserBuilder {
	b.limit = limit
	return b
}

// Cursor sets the cursor for pagination.
func (b *SearchUserBuilder) Cursor(cursor string) *SearchUserBuilder {
	b.cursor = cursor
	return b
}
