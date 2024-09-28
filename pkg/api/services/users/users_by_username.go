package users

import (
	"context"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/jaxron/roapi.go/pkg/api/models"
	"github.com/jaxron/roapi.go/pkg/client"
)

// GetUsersByUsernames fetches information for users with the given usernames.
// POST https://users.roblox.com/v1/usernames/users
func (s *Service) GetUsersByUsernames(ctx context.Context, b *UsersByUsernamesBuilder) (*models.UsersByUsernames, error) {
	var users models.UsersByUsernames
	req, err := client.NewRequest().
		Method(http.MethodPost).
		URL(UsersEndpoint + "/v1/usernames/users").
		Result(&users).
		JSONBody(b.MarshalJSON)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req.Build())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &users, nil
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

// MarshalJSON converts the UsersByUsernamesBuilder to JSON for API requests.
func (b *UsersByUsernamesBuilder) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(struct {
		Usernames          []string `json:"usernames"`
		ExcludeBannedUsers bool     `json:"excludeBannedUsers"`
	}{
		Usernames:          b.usernames,
		ExcludeBannedUsers: b.excludeBannedUsers,
	})
}
