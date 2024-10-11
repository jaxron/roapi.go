package groups

import (
	"context"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// LookupGroup searches for groups based on the provided group name.
// GET https://groups.roblox.com/v1/groups/search/lookup
func (r *Resource) LookupGroup(ctx context.Context, groupName string) ([]types.GroupLookupResponse, error) {
	var lookupResults struct {
		Data []types.GroupLookupResponse `json:"data"`
	}
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(types.GroupsEndpoint+"/v1/groups/search/lookup").
		Query("groupName", groupName).
		Result(&lookupResults).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return lookupResults.Data, nil
}
