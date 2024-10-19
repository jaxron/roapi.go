package friends

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetFriends fetches the friends of a user.
// GET https://friends.roblox.com/v1/users/{userID}/friends
func (r *Resource) GetFriends(ctx context.Context, userID uint64) ([]types.Friend, error) {
	var friends struct {
		Data []types.Friend `json:"data"` // List of friend information
	}
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/friends", types.FriendsEndpoint, userID)).
		Result(&friends).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return friends.Data, nil
}
