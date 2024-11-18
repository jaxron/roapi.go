package users

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetUsersByIDs fetches information for users with the given IDs.
// POST https://users.roblox.com/v1/users
func (r *Resource) GetUsersByIDs(ctx context.Context, p UsersByIDsParams) (*types.UsersByIDsResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var users types.UsersByIDsResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodPost).
		URL(types.UsersEndpoint + "/v1/users").
		Result(&users).
		MarshalBody(p).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	if err := r.validate.Struct(&users); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &users, nil
}

// UsersByIDsParams holds the parameters for fetching users by IDs.
type UsersByIDsParams struct {
	UserIDs            []uint64 `json:"userIds"            validate:"required,min=1,max=100"`
	ExcludeBannedUsers bool     `json:"excludeBannedUsers"`
}

// UsersByIDsBuilder builds parameters for GetUsersByIDs API call.
type UsersByIDsBuilder struct {
	params UsersByIDsParams
}

// NewUsersByIDsBuilder creates a new UsersByIDsBuilder with the given user IDs.
func NewUsersByIDsBuilder(userIDs ...uint64) *UsersByIDsBuilder {
	return &UsersByIDsBuilder{
		params: UsersByIDsParams{
			UserIDs:            userIDs,
			ExcludeBannedUsers: false, // Default: include banned users
		},
	}
}

// ExcludeBannedUsers sets whether to exclude banned users from the result.
func (b *UsersByIDsBuilder) ExcludeBannedUsers(exclude bool) *UsersByIDsBuilder {
	b.params.ExcludeBannedUsers = exclude
	return b
}

// WithUserIDs adds multiple user IDs to the list.
func (b *UsersByIDsBuilder) WithUserIDs(userIDs ...uint64) *UsersByIDsBuilder {
	b.params.UserIDs = append(b.params.UserIDs, userIDs...)
	return b
}

// RemoveUserIDs removes multiple user IDs from the list.
func (b *UsersByIDsBuilder) RemoveUserIDs(userIDs ...uint64) *UsersByIDsBuilder {
	for _, id := range userIDs {
		for i, userID := range b.params.UserIDs {
			if userID == id {
				b.params.UserIDs = append(b.params.UserIDs[:i], b.params.UserIDs[i+1:]...)
				break
			}
		}
	}
	return b
}

// Build returns the UsersByIDsParams.
func (b *UsersByIDsBuilder) Build() UsersByIDsParams {
	return b.params
}
