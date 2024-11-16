package groups

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetRoleUsers fetches the paginated list of users in a specific role of a group.
// GET https://groups.roblox.com/v1/groups/{groupID}/roles/{roleID}/users
func (r *Resource) GetRoleUsers(ctx context.Context, p RoleUsersParams) (*types.RoleUsersResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var roleUsers types.RoleUsersResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/groups/%d/roles/%d/users", types.GroupsEndpoint, p.GroupID, p.RoleID)).
		Query("limit", strconv.FormatUint(p.Limit, 10)).
		Query("cursor", p.Cursor).
		Query("sortOrder", string(p.SortOrder)).
		Result(&roleUsers).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	if err := r.validate.Struct(&roleUsers); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &roleUsers, nil
}

// RoleUsersParams holds the parameters for getting role users.
type RoleUsersParams struct {
	GroupID   uint64          `json:"groupId"   validate:"required,gt=0"`
	RoleID    uint64          `json:"roleId"    validate:"required,gt=0"`
	Limit     uint64          `json:"limit"     validate:"omitempty,oneof=10 25 50 100"`
	Cursor    string          `json:"cursor"    validate:"omitempty"`
	SortOrder types.SortOrder `json:"sortOrder" validate:"omitempty,oneof=Asc Desc"`
}

// RoleUsersBuilder is a builder for RoleUsersParams.
type RoleUsersBuilder struct {
	params RoleUsersParams
}

// NewRoleUsersBuilder creates a new RoleUsersBuilder with default values.
func NewRoleUsersBuilder(groupID, roleID uint64) *RoleUsersBuilder {
	return &RoleUsersBuilder{
		params: RoleUsersParams{
			GroupID:   groupID,
			RoleID:    roleID,
			Limit:     10,
			Cursor:    "",
			SortOrder: "",
		},
	}
}

// WithLimit sets the limit.
func (b *RoleUsersBuilder) WithLimit(limit uint64) *RoleUsersBuilder {
	b.params.Limit = limit
	return b
}

// WithCursor sets the cursor.
func (b *RoleUsersBuilder) WithCursor(cursor string) *RoleUsersBuilder {
	b.params.Cursor = cursor
	return b
}

// WithSortOrderAsc sets the sort order to ascending.
func (b *RoleUsersBuilder) WithSortOrderAsc() *RoleUsersBuilder {
	b.params.SortOrder = types.SortOrderAsc
	return b
}

// WithSortOrderDesc sets the sort order to descending.
func (b *RoleUsersBuilder) WithSortOrderDesc() *RoleUsersBuilder {
	b.params.SortOrder = types.SortOrderDesc
	return b
}

// Build returns the RoleUsersParams.
func (b *RoleUsersBuilder) Build() RoleUsersParams {
	return b.params
}
