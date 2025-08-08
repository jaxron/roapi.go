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
func (r *Resource) GetFriends(ctx context.Context, userID uint64) (*types.FriendsResponse, error) {
	if err := r.validate.Var(userID, "required,gt=0"); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var friends types.FriendsResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/friends", types.FriendsEndpoint, userID)).
		Result(&friends).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if err := r.validate.Struct(&friends); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &friends, nil
}
