package groups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetGroupRoles fetches the roles for a specific group.
// GET https://groups.roblox.com/v1/groups/{groupID}/roles
func (r *Resource) GetGroupRoles(ctx context.Context, groupID int64) (*types.GroupRolesResponse, error) {
	if err := r.validate.Var(groupID, "required,gt=0"); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var groupRoles types.GroupRolesResponse

	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/groups/%d/roles", types.GroupsEndpoint, groupID)).
		Result(&groupRoles).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}

	defer func() { _ = resp.Body.Close() }()

	if err := r.validate.Struct(&groupRoles); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &groupRoles, nil
}
