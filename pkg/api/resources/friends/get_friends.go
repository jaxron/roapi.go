package friends

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errs"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetFriends fetches the friends of a user.
// GET https://friends.roblox.com/v1/users/{userID}/friends
func (r *Resource) GetFriends(ctx context.Context, p GetFriendsParams) (*types.FriendsResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errs.ErrInvalidRequest, err)
	}

	var friends types.FriendsResponse

	req := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/friends", types.FriendsEndpoint, p.UserID)).
		Result(&friends)

	if p.UserSort != types.FriendSortDefault {
		req = req.Query("userSort", string(p.UserSort))
	}

	resp, err := req.Do(ctx)
	if err != nil {
		return nil, errs.HandleAPIError(resp, err)
	}

	defer func() { _ = resp.Body.Close() }()

	if err := r.validate.Struct(&friends); err != nil {
		return nil, fmt.Errorf("%w: %w", errs.ErrInvalidResponse, err)
	}

	return &friends, nil
}

// GetFriendsParams holds the parameters for fetching friends.
type GetFriendsParams struct {
	UserID   int64                `json:"userId"   validate:"required,gt=0"`                   // Required: ID of the user to fetch friends for
	UserSort types.FriendSortType `json:"userSort" validate:"omitempty,oneof=StatusFrequents"` // Optional: Sort order for results
}

// GetFriendsBuilder is a builder for GetFriendsParams.
type GetFriendsBuilder struct {
	params GetFriendsParams
}

// NewGetFriendsBuilder creates a new GetFriendsBuilder with default values.
func NewGetFriendsBuilder(userID int64) *GetFriendsBuilder {
	return &GetFriendsBuilder{
		params: GetFriendsParams{
			UserID:   userID,
			UserSort: types.FriendSortDefault,
		},
	}
}

// WithUserSort sets the sort order for the friends list.
func (b *GetFriendsBuilder) WithUserSort(userSort types.FriendSortType) *GetFriendsBuilder {
	b.params.UserSort = userSort
	return b
}

// Build returns the GetFriendsParams.
func (b *GetFriendsBuilder) Build() GetFriendsParams {
	return b.params
}
