package users

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetUserByID fetches information for a user with the given ID.
// GET https://users.roblox.com/v1/users/{userID}
func (r *Resource) GetUserByID(ctx context.Context, userID uint64) (*types.UserByIDResponse, error) {
	var user types.UserByIDResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d", types.UsersEndpoint, userID)).
		Result(&user).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return &user, nil
}
