package presence

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetUserPresences fetches presence information for multiple users.
// POST https://presence.roblox.com/v1/presence/users
func (r *Resource) GetUserPresences(ctx context.Context, p UserPresencesParams) (*types.UserPresencesResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var presences types.UserPresencesResponse

	resp, err := r.client.NewRequest().
		Method(http.MethodPost).
		URL(types.PresenceEndpoint + "/v1/presence/users").
		MarshalBody(p).
		Result(&presences).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}

	defer func() { _ = resp.Body.Close() }()

	if err := r.validate.Struct(&presences); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &presences, nil
}

// UserPresencesParams holds the parameters for getting user presences.
type UserPresencesParams struct {
	UserIDs []int64 `json:"userIds" validate:"required,min=1,max=50"`
}

// UserPresencesBuilder is a builder for UserPresencesParams.
type UserPresencesBuilder struct {
	params UserPresencesParams
}

// NewUserPresencesBuilder creates a new UserPresencesBuilder.
func NewUserPresencesBuilder(userIDs ...int64) *UserPresencesBuilder {
	return &UserPresencesBuilder{
		params: UserPresencesParams{
			UserIDs: userIDs,
		},
	}
}

// WithUserIDs adds multiple user IDs to the list.
func (b *UserPresencesBuilder) WithUserIDs(userIDs ...int64) *UserPresencesBuilder {
	b.params.UserIDs = append(b.params.UserIDs, userIDs...)
	return b
}

// RemoveUserIDs removes multiple user IDs from the list.
func (b *UserPresencesBuilder) RemoveUserIDs(userIDs ...int64) *UserPresencesBuilder {
	for _, id := range userIDs {
		for i, userID := range b.params.UserIDs {
			if userID == id {
				b.params.UserIDs = append(b.params.UserIDs[:i], b.params.UserIDs[i+1:]...)
				break
			}
		}
	}

	return b
}

// Build returns the UserPresencesParams.
func (b *UserPresencesBuilder) Build() UserPresencesParams {
	return b.params
}
