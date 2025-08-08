package users

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/middleware/auth"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetAuthUserInfo fetches information for the authenticated user.
// GET https://users.roblox.com/v1/users/authenticated
func (r *Resource) GetAuthUserInfo(ctx context.Context) (*types.AuthUserResponse, error) {
	ctx = context.WithValue(ctx, auth.KeyAddCookie, true)

	var user types.AuthUserResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(types.UsersEndpoint + "/v1/users/authenticated").
		Result(&user).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if err := r.validate.Struct(&user); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &user, nil
}
