package groups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// LookupGroup searches for groups based on the provided group name.
// GET https://groups.roblox.com/v1/groups/search/lookup
func (r *Resource) LookupGroup(ctx context.Context, groupName string) (*types.GroupLookupResponse, error) {
	if err := r.validate.Var(groupName, "required"); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var lookupResults types.GroupLookupResponse

	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(types.GroupsEndpoint+"/v1/groups/search/lookup").
		Query("groupName", groupName).
		Result(&lookupResults).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}

	defer func() { _ = resp.Body.Close() }()

	if err := r.validate.Struct(&lookupResults); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &lookupResults, nil
}
