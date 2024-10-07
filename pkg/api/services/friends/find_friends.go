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

// FindFriends fetches the paginated list of friends for a user.
// GET https://friends.roblox.com/v1/users/{userID}/friends/find
func (s *Service) FindFriends(ctx context.Context, params FindFriendsParams) (*models.FriendPageResponse, error) {
	if err := s.validate.Struct(params); err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, auth.KeyAddCookie, true)

	var friends models.FriendPageResponse
	resp, err := s.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/friends/find", FriendsEndpoint, params.UserID)).
		Query("userSort", strconv.FormatUint(params.UserSort, 10)).
		Query("cursor", params.Cursor).
		Query("limit", strconv.FormatUint(params.Limit, 10)).
		Result(&friends).
		JSONHeaders().
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return &friends, nil
}

// FindFriendsParams holds the parameters for finding friends.
type FindFriendsParams struct {
	UserID   uint64 `json:"userId"   validate:"required"`     // Required: ID of the user to fetch friends for
	UserSort uint64 `json:"userSort" validate:"oneof=0 1 2"`  // Optional: Sort order for results
	Limit    uint64 `json:"limit"    validate:"min=1,max=50"` // Optional: Maximum number of results to return (default: 50)
	Cursor   string `json:"cursor"`                           // Optional: Cursor for pagination
}

// FindFriendsBuilder is a builder for FindFriendsParams.
type FindFriendsBuilder struct {
	params FindFriendsParams
}

// NewFindFriendsBuilder creates a new FindFriendsBuilder with default values.
func NewFindFriendsBuilder(userID uint64) *FindFriendsBuilder {
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
func (b *FindFriendsBuilder) WithUserSort(userSort uint64) *FindFriendsBuilder {
	b.params.UserSort = userSort
	return b
}

// WithLimit sets the limit.
func (b *FindFriendsBuilder) WithLimit(limit uint64) *FindFriendsBuilder {
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
