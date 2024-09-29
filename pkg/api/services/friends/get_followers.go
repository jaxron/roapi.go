package friends

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/models"
	"github.com/jaxron/roapi.go/pkg/client"
)

// GetFollowers fetches the paginated followers of a user.
// GET https://friends.roblox.com/v1/users/{userID}/followers
func (s *Service) GetFollowers(ctx context.Context, b *FollowersBuilder) (*models.FollowerPageResponse, error) {
	var followers models.FollowerPageResponse
	req := client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/followers", FriendsEndpoint, b.userID)).
		Query("limit", strconv.FormatUint(b.limit, 10)).
		Query("cursor", b.cursor).
		Query("sortOrder", b.sortOrder).
		UseCookie(true).
		Result(&followers)

	resp, err := s.Client.Do(ctx, req.Build())
	if err != nil {
		return nil, err
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
