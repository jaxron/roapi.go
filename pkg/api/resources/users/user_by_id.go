package users

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errs"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetUserByID fetches information for a user with the given ID.
// GET https://users.roblox.com/v1/users/{userID}
func (r *Resource) GetUserByID(ctx context.Context, userID int64) (*types.UserByIDResponse, error) {
	if err := r.validate.Var(userID, "required,gt=0"); err != nil {
		return nil, fmt.Errorf("%w: %w", errs.ErrInvalidRequest, err)
	}

	var user types.UserByIDResponse

	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d", types.UsersEndpoint, userID)).
		Result(&user).
		Do(ctx)
	if err != nil {
		return nil, errs.HandleAPIError(resp, err)
	}

	defer func() { _ = resp.Body.Close() }()

	if err := r.validate.Struct(&user); err != nil {
		return nil, fmt.Errorf("%w: %w", errs.ErrInvalidResponse, err)
	}

	return &user, nil
}
