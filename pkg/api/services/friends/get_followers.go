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

// GetFollowers fetches the paginated followers of a user.
// GET https://friends.roblox.com/v1/users/{userID}/followers
func (s *Service) GetFollowers(ctx context.Context, b *FollowersBuilder) (*models.FollowerPageResponse, error) {
	ctx = context.WithValue(ctx, auth.KeyAddCookie, true)

	var followers models.FollowerPageResponse
	resp, err := s.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/followers", FriendsEndpoint, b.userID)).
		Query("limit", strconv.FormatUint(b.limit, 10)).
		Query("cursor", b.cursor).
		Query("sortOrder", b.sortOrder).
		Result(&followers).
		JSONHeaders().
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return &followers, nil
}

// FollowersBuilder builds parameters for GetFollowers API call.
type FollowersBuilder struct {
	userID    uint64 // Required: ID of the user to fetch followers for
	limit     uint64 // Optional: Maximum number of results to return (default: 10)
	cursor    string // Optional: Cursor for pagination
	sortOrder string // Optional: Sort order for results
}

// NewFollowersBuilder creates a new FollowersBuilder with the given user ID.
func NewFollowersBuilder(userID uint64) *FollowersBuilder {
	return &FollowersBuilder{
		userID:    userID,
		limit:     10,
		cursor:    "",
		sortOrder: SortOrderAsc,
	}
}

// Limit sets the maximum number of results to return.
func (b *FollowersBuilder) Limit(limit uint64) *FollowersBuilder {
	b.limit = limit
	return b
}

// Cursor sets the cursor for pagination.
func (b *FollowersBuilder) Cursor(cursor string) *FollowersBuilder {
	b.cursor = cursor
	return b
}

// SortOrderAsc sets the sort order for results to ascending.
func (b *FollowersBuilder) SortOrderAsc(sortOrder string) *FollowersBuilder {
	b.sortOrder = SortOrderAsc
	return b
}

// SortOrderDesc sets the sort order for results to descending.
func (b *FollowersBuilder) SortOrderDesc(sortOrder string) *FollowersBuilder {
	b.sortOrder = SortOrderDesc
	return b
}
