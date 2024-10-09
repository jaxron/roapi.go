package groups

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetGroupsInfo fetches information about multiple groups.
// GET https://groups.roblox.com/v2/groups
func (s *Service) GetGroupsInfo(ctx context.Context, p GetGroupsInfoParams) (*types.GroupsInfoResponse, error) {
	if err := s.validate.Struct(p); err != nil {
		return nil, err
	}

	var groupsInfo types.GroupsInfoResponse
	resp, err := s.client.NewRequest().
		Method(http.MethodGet).
		URL(types.GroupsEndpoint+"/v2/groups").
		Query("groupIds", strings.Join(p.GroupIDs, ",")).
		Result(&groupsInfo).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return &groupsInfo, nil
}

// GetGroupsInfoParams holds the parameters for getting information about multiple groups.
type GetGroupsInfoParams struct {
	GroupIDs []string `json:"groupIds" validate:"required,min=1,max=100,dive,numeric"`
}

// GetGroupsInfoBuilder is a builder for GetGroupsInfoParams.
type GetGroupsInfoBuilder struct {
	params GetGroupsInfoParams
}

// NewGetGroupsInfoBuilder creates a new GetGroupsInfoBuilder with default values.
func NewGetGroupsInfoBuilder(groupIDs ...uint64) *GetGroupsInfoBuilder {
	stringGroupIDs := make([]string, len(groupIDs))
	for i, id := range groupIDs {
		stringGroupIDs[i] = strconv.FormatUint(id, 10)
	}

	return &GetGroupsInfoBuilder{
		params: GetGroupsInfoParams{
			GroupIDs: stringGroupIDs,
		},
	}
}

// WithGroupIDs adds multiple group IDs to the list.
func (b *GetGroupsInfoBuilder) WithGroupIDs(groupIDs ...uint64) *GetGroupsInfoBuilder {
	for _, id := range groupIDs {
		b.params.GroupIDs = append(b.params.GroupIDs, strconv.FormatUint(id, 10))
	}
	return b
}

// RemoveGroupIDs removes a group ID from the list.
func (b *GetGroupsInfoBuilder) RemoveGroupIDs(groupIDs ...uint64) *GetGroupsInfoBuilder {
	for _, id := range groupIDs {
		for i, groupID := range b.params.GroupIDs {
			if groupID == strconv.FormatUint(id, 10) {
				b.params.GroupIDs = append(b.params.GroupIDs[:i], b.params.GroupIDs[i+1:]...)
				break
			}
		}
	}
	return b
}

// Build returns the GetGroupsInfoParams.
func (b *GetGroupsInfoBuilder) Build() GetGroupsInfoParams {
	return b.params
}
