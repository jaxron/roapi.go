package avatar

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetUserAvatar fetches the avatar details for a specific user.
// GET https://avatar.roblox.com/v2/avatar/users/{userId}/avatar
func (r *Resource) GetUserAvatar(ctx context.Context, userID uint64) (*types.UserAvatarResponse, error) {
	if err := r.validate.Var(userID, "required,gt=0"); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var userAvatar types.UserAvatarResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v2/avatar/users/%d/avatar", types.AvatarEndpoint, userID)).
		Result(&userAvatar).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	if err := r.validate.Struct(&userAvatar); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &userAvatar, nil
}
