package groups

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetGroupWallPosts fetches the wall posts for a specific group.
// GET https://groups.roblox.com/v2/groups/{groupId}/wall/posts
func (r *Resource) GetGroupWallPosts(ctx context.Context, p GroupWallPostsParams) (*types.GroupWallPostsResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var wallPosts types.GroupWallPostsResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v2/groups/%d/wall/posts", types.GroupsEndpoint, p.GroupID)).
		Query("limit", strconv.FormatUint(p.Limit, 10)).
		Query("cursor", p.Cursor).
		Query("sortOrder", string(p.SortOrder)).
		Result(&wallPosts).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	if err := r.validate.Struct(&wallPosts); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &wallPosts, nil
}

// GroupWallPostsParams holds the parameters for getting group wall posts.
type GroupWallPostsParams struct {
	GroupID   uint64          `json:"groupId"   validate:"required,gt=0"`
	Limit     uint64          `json:"limit"     validate:"omitempty,oneof=10 25 50 100"`
	Cursor    string          `json:"cursor"    validate:"omitempty"`
	SortOrder types.SortOrder `json:"sortOrder" validate:"omitempty,oneof=Asc Desc"`
}

// GroupWallPostsBuilder is a builder for GroupWallPostsParams.
type GroupWallPostsBuilder struct {
	params GroupWallPostsParams
}

// NewGroupWallPostsBuilder creates a new GroupWallPostsBuilder with default values.
func NewGroupWallPostsBuilder(groupID uint64) *GroupWallPostsBuilder {
	return &GroupWallPostsBuilder{
		params: GroupWallPostsParams{
			GroupID:   groupID,
			Limit:     10,
			Cursor:    "",
			SortOrder: "",
		},
	}
}

// WithLimit sets the limit.
func (b *GroupWallPostsBuilder) WithLimit(limit uint64) *GroupWallPostsBuilder {
	b.params.Limit = limit
	return b
}

// WithCursor sets the cursor.
func (b *GroupWallPostsBuilder) WithCursor(cursor string) *GroupWallPostsBuilder {
	b.params.Cursor = cursor
	return b
}

// WithSortOrderAsc sets the sort order to ascending.
func (b *GroupWallPostsBuilder) WithSortOrderAsc() *GroupWallPostsBuilder {
	b.params.SortOrder = types.SortOrderAsc
	return b
}

// WithSortOrderDesc sets the sort order to descending.
func (b *GroupWallPostsBuilder) WithSortOrderDesc() *GroupWallPostsBuilder {
	b.params.SortOrder = types.SortOrderDesc
	return b
}

// Build returns the GroupWallPostsParams.
func (b *GroupWallPostsBuilder) Build() GroupWallPostsParams {
	return b.params
}
