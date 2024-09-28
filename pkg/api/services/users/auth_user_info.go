package users

import (
	"context"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/models"
	"github.com/jaxron/roapi.go/pkg/client"
	"github.com/jaxron/roapi.go/pkg/errors"
)

// GetAuthUserInfo fetches information for the authenticated user.
// GET https://users.roblox.com/v1/users/authenticated
func (s *Service) GetAuthUserInfo(ctx context.Context) (*models.AuthUserInfo, error) {
	if s.Client.Handler.Auth.GetCookieCount() == 0 {
		return nil, errors.ErrNoCookie
	}

	var user models.AuthUserInfo
	req := client.NewRequest().
		Method(http.MethodGet).
		URL(UsersEndpoint + "/v1/users/authenticated").
		Result(&user).
		UseCookie(true)

	resp, err := s.Client.Do(ctx, req.Build())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &user, nil
}
