package users

import (
	"context"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/models"
)

// GetUsersByIDs fetches information for users with the given IDs.
// POST https://users.roblox.com/v1/users
func (s *Service) GetUsersByIDs(ctx context.Context, params UsersByIDsParams) ([]models.VerifiedBadgeUserResponse, error) {
	if err := s.validate.Struct(params); err != nil {
		return nil, err
	}

	var users struct {
		Data []models.VerifiedBadgeUserResponse `json:"data"` // List of users fetched by user IDs
	}
	resp, err := s.client.NewRequest().
		Method(http.MethodPost).
		URL(UsersEndpoint + "/v1/users").
		Result(&users).
		MarshalBody(params).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return users.Data, nil
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
func NewUsersByIDsBuilder(userIDs []uint64) *UsersByIDsBuilder {
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

// Build returns the UsersByIDsParams.
func (b *UsersByIDsBuilder) Build() UsersByIDsParams {
	return b.params
}
