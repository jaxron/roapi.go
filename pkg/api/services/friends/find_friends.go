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
func (s *Service) FindFriends(ctx context.Context, b *FindFriendsBuilder) (*models.FriendPageResponse, error) {
	ctx = context.WithValue(ctx, auth.KeyAddCookie, true)

	var friends models.FriendPageResponse
	resp, err := s.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/friends/find", FriendsEndpoint, b.userID)).
		Query("userSort", strconv.FormatUint(b.userSort, 10)).
		Query("cursor", b.cursor).
		Query("limit", strconv.FormatUint(b.limit, 10)).
		Result(&friends).
		JSONHeaders().
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return &friends, nil
}

// FindFriendsBuilder builds parameters for FindFriends API call.
type FindFriendsBuilder struct {
	userID   uint64 // Required: ID of the user to fetch friends for
	userSort uint64 // Optional: Sort order for results
	cursor   string // Optional: Cursor for pagination
	limit    uint64 // Optional: Maximum number of results to return (default: 200)
}

// NewFindFriendsBuilder creates a new FindFriendsBuilder with the given user ID.
func NewFindFriendsBuilder(userID uint64) *FindFriendsBuilder {
	return &FindFriendsBuilder{
		userID:   userID,
		userSort: 2,
		cursor:   "",
		limit:    50,
	}
}

// UserSort sets the sort order for results.
func (b *FindFriendsBuilder) UserSort(userSort uint64) *FindFriendsBuilder {
	b.userSort = userSort
	return b
}

// Cursor sets the cursor for pagination.
func (b *FindFriendsBuilder) Cursor(cursor string) *FindFriendsBuilder {
	b.cursor = cursor
	return b
}

// Limit sets the maximum number of results to return.
func (b *FindFriendsBuilder) Limit(limit uint64) *FindFriendsBuilder {
	b.limit = limit
	return b
}
