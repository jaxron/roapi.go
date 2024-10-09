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
func (s *Service) GetGroupRoles(ctx context.Context, groupID uint64) (*types.GroupRolesResponse, error) {
	var groupRoles types.GroupRolesResponse
	resp, err := s.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/groups/%d/roles", types.GroupsEndpoint, groupID)).
		Result(&groupRoles).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return &groupRoles, nil
}
