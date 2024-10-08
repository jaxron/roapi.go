package friends

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/internal/middleware/auth"
	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/models"
)

// SearchFriends fetches the paginated list of friends for a user.
// GET https://friends.roblox.com/v1/users/{userID}/friends/search
func (s *Service) SearchFriends(ctx context.Context, params SearchFriendsParams) (*models.FriendPageResponse, error) {
	if err := s.validate.Struct(params); err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, auth.KeyAddCookie, true)

	var friends models.FriendPageResponse
	resp, err := s.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/friends/search", FriendsEndpoint, params.UserID)).
		Query("query", params.Query).
		Query("cursor", params.Cursor).
		Query("limit", strconv.FormatUint(params.Limit, 10)).
		Result(&friends).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return &friends, nil
}

// SearchFriendsParams holds the parameters for searching friends.
type SearchFriendsParams struct {
	UserID uint64 `json:"userId" validate:"required"`         // Required: ID of the user to fetch friends for
	Query  string `json:"query"`                              // Optional: Search keyword
	Limit  uint64 `json:"limit"  validate:"min=1,max=50"`     // Optional: Maximum number of results to return (default: 20)
	Cursor string `json:"cursor" validate:"omitempty,base64"` // Optional: Cursor for pagination
}

// SearchFriendsBuilder is a builder for SearchFriendsParams.
type SearchFriendsBuilder struct {
	params SearchFriendsParams
}

// NewSearchFriendsBuilder creates a new SearchFriendsBuilder with default values.
func NewSearchFriendsBuilder(userID uint64) *SearchFriendsBuilder {
	return &SearchFriendsBuilder{
		params: SearchFriendsParams{
			UserID: userID,
			Query:  "",
			Cursor: "",
			Limit:  20,
		},
	}
}

// WithQuery sets the search keyword.
func (b *SearchFriendsBuilder) WithQuery(query string) *SearchFriendsBuilder {
	b.params.Query = query
	return b
}

// WithLimit sets the maximum number of results to return.
func (b *SearchFriendsBuilder) WithLimit(limit uint64) *SearchFriendsBuilder {
	b.params.Limit = limit
	return b
}

// WithCursor sets the cursor for pagination.
func (b *SearchFriendsBuilder) WithCursor(cursor string) *SearchFriendsBuilder {
	b.params.Cursor = cursor
	return b
}

// Build returns the SearchFriendsParams.
func (b *SearchFriendsBuilder) Build() SearchFriendsParams {
	return b.params
}
