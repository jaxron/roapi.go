package users

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/models"
	"github.com/jaxron/roapi.go/pkg/client"
)

// GetUserByID fetches information for a user with the given ID.
// GET https://users.roblox.com/v1/users/{userID}
func (s *Service) GetUserByID(ctx context.Context, userID uint64) (*models.UserByIDResponse, error) {
	var user models.UserByIDResponse
	req := client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d", UsersEndpoint, userID)).
		Result(&user)

	resp, err := s.client.Do(ctx, req.Build())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &user, nil
}
