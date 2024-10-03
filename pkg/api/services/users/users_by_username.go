package users

import (
	"context"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/models"
)

// GetUsersByUsernames fetches information for users with the given usernames.
// POST https://users.roblox.com/v1/usernames/users
func (s *Service) GetUsersByUsernames(ctx context.Context, b *UsersByUsernamesBuilder) ([]models.UserByUsernameResponse, error) {
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
			Usernames:          b.usernames,
			ExcludeBannedUsers: b.excludeBannedUsers,
		}).
		JSONHeaders().
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return users.Data, nil
}

// UsersByUsernamesBuilder builds parameters for GetUsersByUsernames API call.
type UsersByUsernamesBuilder struct {
	usernames          []string // Required: List of usernames to fetch information for
	excludeBannedUsers bool     // Optional: Whether to exclude banned users from the result
}

// NewUsersByUsernamesBuilder creates a new UsersByUsernamesBuilder with the given usernames.
func NewUsersByUsernamesBuilder(usernames []string) *UsersByUsernamesBuilder {
	return &UsersByUsernamesBuilder{
		usernames:          usernames,
		excludeBannedUsers: false, // Default: include banned users
	}
}

// ExcludeBannedUsers sets whether to exclude banned users from the result.
func (b *UsersByUsernamesBuilder) ExcludeBannedUsers(excludeBannedUsers bool) *UsersByUsernamesBuilder {
	b.excludeBannedUsers = excludeBannedUsers
	return b
}
