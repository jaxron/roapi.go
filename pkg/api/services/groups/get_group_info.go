package groups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetGroupInfo fetches information about a specific group.
// GET https://groups.roblox.com/v1/groups/{groupID}
func (s *Service) GetGroupInfo(ctx context.Context, groupID uint64) (*types.GroupResponse, error) {
	var groupInfo types.GroupResponse
	resp, err := s.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/groups/%d", types.GroupsEndpoint, groupID)).
		Result(&groupInfo).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return &groupInfo, nil
}
