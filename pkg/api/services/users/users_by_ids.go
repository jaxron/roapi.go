package users

import (
	"context"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/jaxron/roapi.go/pkg/api/client"
	"github.com/jaxron/roapi.go/pkg/api/models"
)

// GetUsersByIDs fetches information for users with the given IDs.
// POST https://users.roblox.com/v1/users
func (s *Service) GetUsersByIDs(ctx context.Context, b *UsersByIDsBuilder) ([]models.VerifiedBadgeUserResponse, error) {
	var users struct {
		Data []models.VerifiedBadgeUserResponse `json:"data"` // List of users fetched by user IDs
	}
	req, err := client.NewRequest().
		Method(http.MethodPost).
		URL(UsersEndpoint + "/v1/users").
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

	return users.Data, nil
}

// UsersByIDsBuilder builds parameters for GetUsersByIDs API call.
type UsersByIDsBuilder struct {
	userIds            []uint64 // Required: List of user IDs to fetch information for
	excludeBannedUsers bool     // Optional: Whether to exclude banned users from the result
}

// NewUsersByIDsBuilder creates a new UsersByIDsBuilder with the given user IDs.
func NewUsersByIDsBuilder(userIds []uint64) *UsersByIDsBuilder {
	return &UsersByIDsBuilder{
		userIds:            userIds,
		excludeBannedUsers: false, // Default: include banned users
	}
}

// ExcludeBannedUsers sets whether to exclude banned users from the result.
func (b *UsersByIDsBuilder) ExcludeBannedUsers(excludeBannedUsers bool) *UsersByIDsBuilder {
	b.excludeBannedUsers = excludeBannedUsers
	return b
}

// MarshalJSON converts the UsersByIDsBuilder to JSON for API requests.
func (b *UsersByIDsBuilder) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(struct {
		UserIds            []uint64 `json:"userIds"`
		ExcludeBannedUsers bool     `json:"excludeBannedUsers"`
	}{
		UserIds:            b.userIds,
		ExcludeBannedUsers: b.excludeBannedUsers,
	})
}
