package users

import (
	"context"
	"net/http"

	"github.com/jaxron/roapi.go/internal/middleware/auth"
	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetAuthUserInfo fetches information for the authenticated user.
// GET https://users.roblox.com/v1/users/authenticated
func (s *Service) GetAuthUserInfo(ctx context.Context) (*types.AuthUserResponse, error) {
	ctx = context.WithValue(ctx, auth.KeyAddCookie, true)

	var user types.AuthUserResponse
	resp, err := s.client.NewRequest().
		Method(http.MethodGet).
		URL(types.UsersEndpoint + "/v1/users/authenticated").
		Result(&user).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return &user, nil
}
