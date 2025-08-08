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
func (r *Resource) GetFriendCount(ctx context.Context, userID int64) (int64, error) {
	if err := r.validate.Var(userID, "required,gt=0"); err != nil {
		return 0, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var count struct {
		Count int64 `json:"count"` // The number of friends
	}

	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/friends/count", types.FriendsEndpoint, userID)).
		Result(&count).
		Do(ctx)
	if err != nil {
		return 0, errors.HandleAPIError(resp, err)
	}

	defer func() { _ = resp.Body.Close() }()

	return count.Count, nil
}
