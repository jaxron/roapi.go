package groups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetUserGroupRoles fetches the group roles for a specific user.
// GET https://groups.roblox.com/v1/users/{userId}/groups/roles
func (r *Resource) GetUserGroupRoles(ctx context.Context, userID uint64) (*types.UserGroupRolesResponse, error) {
	var userGroupRoles types.UserGroupRolesResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/groups/roles", types.GroupsEndpoint, userID)).
		Result(&userGroupRoles).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return &userGroupRoles, nil
}
