package friends

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetFollowingCount fetches the count of users a user is following.
// GET https://friends.roblox.com/v1/users/{userID}/followings/count
func (r *Resource) GetFollowingCount(ctx context.Context, userID uint64) (uint64, error) {
	if err := r.validate.Var(userID, "required,gt=0"); err != nil {
		return 0, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var count struct {
		Count uint64 `json:"count"` // The number of users being followed
	}
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/followings/count", types.FriendsEndpoint, userID)).
		Result(&count).
		Do(ctx)
	if err != nil {
		return 0, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return count.Count, nil
}
