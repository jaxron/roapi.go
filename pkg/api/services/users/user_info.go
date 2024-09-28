package users

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/models"
	"github.com/jaxron/roapi.go/pkg/client"
)

// GetUserInfo fetches information for a user with the given ID.
// GET https://users.roblox.com/v1/users/{userID}
func (s *Service) GetUserInfo(ctx context.Context, userID uint64) (*models.UserInfo, error) {
	var user models.UserInfo
	req := client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d", UsersEndpoint, userID)).
		Result(&user)

	resp, err := s.Client.Do(ctx, req.Build())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &user, nil
}
