package friends

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetFollowers fetches the paginated followers of a user.
// GET https://friends.roblox.com/v1/users/{userID}/followers
func (r *Resource) GetFollowers(ctx context.Context, p GetFollowersParams) (*types.FollowerPageResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, err
	}

	var followers types.FollowerPageResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/followers", types.FriendsEndpoint, p.UserID)).
		Query("limit", strconv.FormatUint(p.Limit, 10)).
		Query("cursor", p.Cursor).
		Query("sortOrder", string(p.SortOrder)).
		Result(&followers).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return &followers, nil
}

// GetFollowersParams holds the parameters for getting followers.
type GetFollowersParams struct {
	UserID    uint64          `json:"userId"    validate:"required"`                 // Required: ID of the user to fetch followers for
	Limit     uint64          `json:"limit"     validate:"oneof=10 18 25 50 100"`    // Optional: Maximum number of results to return (default: 10)
	Cursor    string          `json:"cursor"    validate:"omitempty,base64"`         // Optional: Cursor for pagination
	SortOrder types.SortOrder `json:"sortOrder" validate:"omitempty,oneof=Asc Desc"` // Optional: Sort order for results
}

// GetFollowersBuilder is a builder for GetFollowersParams.
type GetFollowersBuilder struct {
	params GetFollowersParams
}

// NewGetFollowersBuilder creates a new GetFollowersBuilder with default values.
func NewGetFollowersBuilder(userID uint64) *GetFollowersBuilder {
	return &GetFollowersBuilder{
		params: GetFollowersParams{
			UserID:    userID,
			Limit:     10,
			Cursor:    "",
			SortOrder: "",
		},
	}
}

// WithLimit sets the limit.
func (b *GetFollowersBuilder) WithLimit(limit uint64) *GetFollowersBuilder {
	b.params.Limit = limit
	return b
}

// WithCursor sets the cursor.
func (b *GetFollowersBuilder) WithCursor(cursor string) *GetFollowersBuilder {
	b.params.Cursor = cursor
	return b
}

// WithSortOrderAsc sets the sort order to ascending.
func (b *GetFollowersBuilder) WithSortOrderAsc() *GetFollowersBuilder {
	b.params.SortOrder = types.SortOrderAsc
	return b
}

// WithSortOrderDesc sets the sort order to descending.
func (b *GetFollowersBuilder) WithSortOrderDesc() *GetFollowersBuilder {
	b.params.SortOrder = types.SortOrderDesc
	return b
}

// Build returns the GetFollowersParams.
func (b *GetFollowersBuilder) Build() GetFollowersParams {
	return b.params
}
