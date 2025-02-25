package friends

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetFollowerCount fetches the count of followers for a user.
// GET https://friends.roblox.com/v1/users/{userID}/followers/count
func (r *Resource) GetFollowerCount(ctx context.Context, userID uint64) (uint64, error) {
	if err := r.validate.Var(userID, "required,gt=0"); err != nil {
		return 0, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var count struct {
		Count uint64 `json:"count"` // The number of followers
	}
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/followers/count", types.FriendsEndpoint, userID)).
		Result(&count).
		Do(ctx)
	if err != nil {
		return 0, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return count.Count, nil
}
