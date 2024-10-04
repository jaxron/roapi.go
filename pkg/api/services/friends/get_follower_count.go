package friends

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
)

// GetFollowerCount fetches the count of followers for a user.
// GET https://friends.roblox.com/v1/users/{userID}/followers/count
func (s *Service) GetFollowerCount(ctx context.Context, userID uint64) (uint64, error) {
	var count struct {
		Count uint64 `json:"count"` // The number of followers
	}
	resp, err := s.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/followers/count", FriendsEndpoint, userID)).
		Result(&count).
		JSONHeaders().
		Do(ctx)
	if err != nil {
		return 0, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return count.Count, nil
}
