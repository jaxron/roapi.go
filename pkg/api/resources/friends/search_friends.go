package friends

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/middleware/auth"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// SearchFriends fetches the paginated list of friends for a user.
// GET https://friends.roblox.com/v1/users/{userID}/friends/search
func (r *Resource) SearchFriends(ctx context.Context, p SearchFriendsParams) (*types.FriendPageResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	ctx = context.WithValue(ctx, auth.KeyAddCookie, true)

	var friends types.FriendPageResponse

	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/friends/search", types.FriendsEndpoint, p.UserID)).
		Query("query", p.Query).
		Query("cursor", p.Cursor).
		Query("limit", strconv.FormatInt(p.Limit, 10)).
		Result(&friends).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}

	defer func() { _ = resp.Body.Close() }()

	if err := r.validate.Struct(&friends); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &friends, nil
}

// SearchFriendsParams holds the parameters for searching friends.
type SearchFriendsParams struct {
	UserID int64  `json:"userId" validate:"required,gt=0"`    // Required: ID of the user to fetch friends for
	Query  string `json:"query"`                              // Optional: Search keyword
	Limit  int64  `json:"limit"  validate:"min=1,max=50"`     // Optional: Maximum number of results to return (default: 20)
	Cursor string `json:"cursor" validate:"omitempty,base64"` // Optional: Cursor for pagination
}

// SearchFriendsBuilder is a builder for SearchFriendsParams.
type SearchFriendsBuilder struct {
	params SearchFriendsParams
}

// NewSearchFriendsBuilder creates a new SearchFriendsBuilder with default values.
func NewSearchFriendsBuilder(userID int64) *SearchFriendsBuilder {
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
func (b *SearchFriendsBuilder) WithLimit(limit int64) *SearchFriendsBuilder {
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
