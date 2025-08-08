package groups

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetGroupUsers fetches the paginated list of users in a group.
// GET https://groups.roblox.com/v1/groups/{groupID}/users
func (r *Resource) GetGroupUsers(ctx context.Context, p GroupUsersParams) (*types.GroupUsersResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var groupUsers types.GroupUsersResponse

	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/groups/%d/users", types.GroupsEndpoint, p.GroupID)).
		Query("limit", strconv.FormatInt(p.Limit, 10)).
		Query("cursor", p.Cursor).
		Query("sortOrder", string(p.SortOrder)).
		Result(&groupUsers).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}

	defer func() { _ = resp.Body.Close() }()

	if err := r.validate.Struct(&groupUsers); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &groupUsers, nil
}

// GroupUsersParams holds the parameters for getting group users.
type GroupUsersParams struct {
	GroupID   int64           `json:"groupId"   validate:"required,gt=0"`
	Limit     int64           `json:"limit"     validate:"omitempty,oneof=10 25 50 100"`
	Cursor    string          `json:"cursor"    validate:"omitempty"`
	SortOrder types.SortOrder `json:"sortOrder" validate:"omitempty,oneof=Asc Desc"`
}

// GroupUsersBuilder is a builder for GroupUsersParams.
type GroupUsersBuilder struct {
	params GroupUsersParams
}

// NewGroupUsersBuilder creates a new GroupUsersBuilder with default values.
func NewGroupUsersBuilder(groupID int64) *GroupUsersBuilder {
	return &GroupUsersBuilder{
		params: GroupUsersParams{
			GroupID:   groupID,
			Limit:     10,
			Cursor:    "",
			SortOrder: "",
		},
	}
}

// WithLimit sets the limit.
func (b *GroupUsersBuilder) WithLimit(limit int64) *GroupUsersBuilder {
	b.params.Limit = limit
	return b
}

// WithCursor sets the cursor.
func (b *GroupUsersBuilder) WithCursor(cursor string) *GroupUsersBuilder {
	b.params.Cursor = cursor
	return b
}

// WithSortOrderAsc sets the sort order to ascending.
func (b *GroupUsersBuilder) WithSortOrderAsc() *GroupUsersBuilder {
	b.params.SortOrder = types.SortOrderAsc
	return b
}

// WithSortOrderDesc sets the sort order to descending.
func (b *GroupUsersBuilder) WithSortOrderDesc() *GroupUsersBuilder {
	b.params.SortOrder = types.SortOrderDesc
	return b
}

// Build returns the GroupUsersParams.
func (b *GroupUsersBuilder) Build() GroupUsersParams {
	return b.params
}
