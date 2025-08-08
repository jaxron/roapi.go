package friends

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// FindFriends fetches the paginated list of friends for a user.
// GET https://friends.roblox.com/v1/users/{userID}/friends/find
func (r *Resource) FindFriends(ctx context.Context, p FindFriendsParams) (*types.FriendPageResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var friends types.FriendPageResponse

	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/friends/find", types.FriendsEndpoint, p.UserID)).
		Query("userSort", strconv.FormatInt(p.UserSort, 10)).
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

// FindFriendsParams holds the parameters for finding friends.
type FindFriendsParams struct {
	UserID   int64  `json:"userId"   validate:"required,gt=0"` // Required: ID of the user to fetch friends for
	UserSort int64  `json:"userSort" validate:"oneof=0 1 2"`   // Optional: Sort order for results
	Limit    int64  `json:"limit"    validate:"min=1,max=50"`  // Optional: Maximum number of results to return (default: 50)
	Cursor   string `json:"cursor"`                            // Optional: Cursor for pagination
}

// FindFriendsBuilder is a builder for FindFriendsParams.
type FindFriendsBuilder struct {
	params FindFriendsParams
}

// NewFindFriendsBuilder creates a new FindFriendsBuilder with default values.
func NewFindFriendsBuilder(userID int64) *FindFriendsBuilder {
	return &FindFriendsBuilder{
		params: FindFriendsParams{
			UserID:   userID,
			UserSort: 2,
			Limit:    10,
			Cursor:   "",
		},
	}
}

// WithUserSort sets the user sort.
func (b *FindFriendsBuilder) WithUserSort(userSort int64) *FindFriendsBuilder {
	b.params.UserSort = userSort
	return b
}

// WithLimit sets the limit.
func (b *FindFriendsBuilder) WithLimit(limit int64) *FindFriendsBuilder {
	b.params.Limit = limit
	return b
}

// WithCursor sets the cursor.
func (b *FindFriendsBuilder) WithCursor(cursor string) *FindFriendsBuilder {
	b.params.Cursor = cursor
	return b
}

// Build returns the FindFriendsParams.
func (b *FindFriendsBuilder) Build() FindFriendsParams {
	return b.params
}
