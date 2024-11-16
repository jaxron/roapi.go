package friends

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetFollowings fetches the paginated followings of a user.
// GET https://friends.roblox.com/v1/users/{userID}/followings
func (r *Resource) GetFollowings(ctx context.Context, p GetFollowingsParams) (*types.FollowingPageResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var followings types.FollowingPageResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/followings", types.FriendsEndpoint, p.UserID)).
		Query("limit", strconv.FormatUint(p.Limit, 10)).
		Query("cursor", p.Cursor).
		Query("sortOrder", string(p.SortOrder)).
		Result(&followings).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	if err := r.validate.Struct(&followings); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &followings, nil
}

// GetFollowingsParams holds the parameters for getting followings.
type GetFollowingsParams struct {
	UserID    uint64          `json:"userId"    validate:"required,gt=0"`                   // Required: ID of the user to fetch followings for
	Limit     uint64          `json:"limit"     validate:"omitempty,oneof=10 18 25 50 100"` // Optional: Maximum number of results to return (default: 10)
	Cursor    string          `json:"cursor"    validate:"omitempty,base64"`                // Optional: Cursor for pagination
	SortOrder types.SortOrder `json:"sortOrder" validate:"omitempty,oneof=Asc Desc"`        // Optional: Sort order for results
}

// GetFollowingsBuilder is a builder for GetFollowingsParams.
type GetFollowingsBuilder struct {
	params GetFollowingsParams
}

// NewGetFollowingsBuilder creates a new GetFollowingsBuilder with default values.
func NewGetFollowingsBuilder(userID uint64) *GetFollowingsBuilder {
	return &GetFollowingsBuilder{
		params: GetFollowingsParams{
			UserID:    userID,
			Limit:     10,
			Cursor:    "",
			SortOrder: "",
		},
	}
}

// WithLimit sets the limit.
func (b *GetFollowingsBuilder) WithLimit(limit uint64) *GetFollowingsBuilder {
	b.params.Limit = limit
	return b
}

// WithCursor sets the cursor.
func (b *GetFollowingsBuilder) WithCursor(cursor string) *GetFollowingsBuilder {
	b.params.Cursor = cursor
	return b
}

// WithSortOrderAsc sets the sort order to ascending.
func (b *GetFollowingsBuilder) WithSortOrderAsc() *GetFollowingsBuilder {
	b.params.SortOrder = types.SortOrderAsc
	return b
}

// WithSortOrderDesc sets the sort order to descending.
func (b *GetFollowingsBuilder) WithSortOrderDesc() *GetFollowingsBuilder {
	b.params.SortOrder = types.SortOrderDesc
	return b
}

// Build returns the GetFollowingsParams.
func (b *GetFollowingsBuilder) Build() GetFollowingsParams {
	return b.params
}
