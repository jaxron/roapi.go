package users

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/models"
)

// GetUserByID fetches information for a user with the given ID.
// GET https://users.roblox.com/v1/users/{userID}
func (s *Service) GetUserByID(ctx context.Context, userID uint64) (*models.UserByIDResponse, error) {
	var user models.UserByIDResponse
	resp, err := s.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d", UsersEndpoint, userID)).
		Result(&user).
		JSONHeaders().
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return &user, nil
}
