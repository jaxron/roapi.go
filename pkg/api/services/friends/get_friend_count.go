package friends

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetFriendCount fetches the count of friends for a user.
// GET https://friends.roblox.com/v1/users/{userID}/friends/count
func (s *Service) GetFriendCount(ctx context.Context, userID uint64) (uint64, error) {
	var count struct {
		Count uint64 `json:"count"` // The number of friends
	}
	resp, err := s.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/friends/count", types.FriendsEndpoint, userID)).
		Result(&count).
		Do(ctx)
	if err != nil {
		return 0, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return count.Count, nil
}
