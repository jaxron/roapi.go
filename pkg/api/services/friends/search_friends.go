package friends

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/client"
	"github.com/jaxron/roapi.go/pkg/api/models"
	"github.com/jaxron/roapi.go/pkg/errors"
)

// SearchFriends fetches the paginated list of friends for a user.
// GET https://friends.roblox.com/v1/users/{userID}/friends/search
func (s *Service) SearchFriends(ctx context.Context, b *SearchFriendsBuilder) (*models.FriendPageResponse, error) {
	if s.Client.Handler.Auth.GetCookieCount() == 0 {
		return nil, errors.ErrNoCookie
	}

	var friends models.FriendPageResponse
	req := client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/friends/search", FriendsEndpoint, b.userID)).
		Query("query", b.query).
		Query("cursor", b.cursor).
		Query("limit", strconv.FormatUint(b.limit, 10)).
		UseCookie(true).
		Result(&friends)

	resp, err := s.Client.Do(ctx, req.Build())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &friends, nil
}

// SearchFriendsBuilder builds parameters for SearchFriends API call.
type SearchFriendsBuilder struct {
	userID uint64 // Required: ID of the user to fetch friends for
	query  string // Optional: Search keyword
	cursor string // Optional: Cursor for pagination
	limit  uint64 // Optional: Maximum number of results to return (default: 20)
}

// NewSearchFriendsBuilder creates a new SearchFriendsBuilder with the given user ID.
func NewSearchFriendsBuilder(userID uint64) *SearchFriendsBuilder {
	return &SearchFriendsBuilder{
		userID: userID,
		query:  "",
		cursor: "",
		limit:  20,
	}
}

// Query sets the search keyword.
func (b *SearchFriendsBuilder) Query(query string) *SearchFriendsBuilder {
	b.query = query
	return b
}

// Cursor sets the cursor for pagination.
func (b *SearchFriendsBuilder) Cursor(cursor string) *SearchFriendsBuilder {
	b.cursor = cursor
	return b
}

// Limit sets the maximum number of results to return.
func (b *SearchFriendsBuilder) Limit(limit uint64) *SearchFriendsBuilder {
	b.limit = limit
	return b
}
