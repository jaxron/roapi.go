package users

import (
	"context"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/models"
)

// GetUsersByUsernames fetches information for users with the given usernames.
// POST https://users.roblox.com/v1/usernames/users
func (s *Service) GetUsersByUsernames(ctx context.Context, params GetUsersByUsernamesParams) ([]models.UserByUsernameResponse, error) {
	if err := s.validate.Struct(params); err != nil {
		return nil, err
	}

	var users struct {
		Data []models.UserByUsernameResponse `json:"data"` // List of users fetched by usernames
	}
	resp, err := s.client.NewRequest().
		Method(http.MethodPost).
		URL(UsersEndpoint + "/v1/usernames/users").
		Result(&users).
		MarshalBody(struct {
			Usernames          []string `json:"usernames"`
			ExcludeBannedUsers bool     `json:"excludeBannedUsers"`
		}{
			Usernames:          params.Usernames,
			ExcludeBannedUsers: params.ExcludeBannedUsers,
		}).
		JSONHeaders().
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return users.Data, nil
}

// GetUsersByUsernamesParams holds the parameters for fetching users by usernames.
type GetUsersByUsernamesParams struct {
	Usernames          []string `json:"usernames"          validate:"required,min=1,max=100"`
	ExcludeBannedUsers bool     `json:"excludeBannedUsers"`
}

// GetUsersByUsernamesBuilder is a builder for GetUsersByUsernamesParams.
type GetUsersByUsernamesBuilder struct {
	params GetUsersByUsernamesParams
}

// NewGetUsersByUsernamesBuilder creates a new GetUsersByUsernamesBuilder with default values.
func NewGetUsersByUsernamesBuilder(usernames []string) *GetUsersByUsernamesBuilder {
	return &GetUsersByUsernamesBuilder{
		params: GetUsersByUsernamesParams{
			Usernames:          usernames,
			ExcludeBannedUsers: false,
		},
	}
}

// ExcludeBannedUsers sets whether to exclude banned users from the result.
func (b *GetUsersByUsernamesBuilder) ExcludeBannedUsers(exclude bool) *GetUsersByUsernamesBuilder {
	b.params.ExcludeBannedUsers = exclude
	return b
}

// Build returns the GetUsersByUsernamesParams.
func (b *GetUsersByUsernamesBuilder) Build() GetUsersByUsernamesParams {
	return b.params
}
