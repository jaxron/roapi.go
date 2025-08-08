package friends

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/middleware/auth"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetOnlineFriends fetches the online friends of a user.
// GET https://friends.roblox.com/v1/users/{userID}/friends/online
func (r *Resource) GetOnlineFriends(ctx context.Context, p GetOnlineFriendsParams) ([]*types.OnlineFriend, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	ctx = context.WithValue(ctx, auth.KeyAddCookie, true)

	var friends struct {
		Data []*types.OnlineFriend `json:"data" validate:"required,dive"` // List of online friends
	}
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/friends/online", types.FriendsEndpoint, p.UserID)).
		Query("userSort", strconv.FormatUint(p.UserSort, 10)).
		Result(&friends).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if err := r.validate.Struct(&friends); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return friends.Data, nil
}

// GetOnlineFriendsParams holds the parameters for fetching online friends.
type GetOnlineFriendsParams struct {
	UserID   uint64 `json:"userId"   validate:"required,gt=0"` // Required: ID of the user to fetch online friends for
	UserSort uint64 `json:"userSort" validate:"oneof=0 1 2"`   // Optional: Sort order for results (default: 0)
}

// GetOnlineFriendsBuilder is a builder for GetOnlineFriendsParams.
type GetOnlineFriendsBuilder struct {
	params GetOnlineFriendsParams
}

// NewGetOnlineFriendsBuilder creates a new GetOnlineFriendsBuilder with default values.
func NewGetOnlineFriendsBuilder(userID uint64) *GetOnlineFriendsBuilder {
	return &GetOnlineFriendsBuilder{
		params: GetOnlineFriendsParams{
			UserID:   userID,
			UserSort: 0,
		},
	}
}

// WithUserSort sets the user sort for the request.
func (b *GetOnlineFriendsBuilder) WithUserSort(userSort uint64) *GetOnlineFriendsBuilder {
	b.params.UserSort = userSort
	return b
}

// Build returns the GetOnlineFriendsParams.
func (b *GetOnlineFriendsBuilder) Build() GetOnlineFriendsParams {
	return b.params
}
