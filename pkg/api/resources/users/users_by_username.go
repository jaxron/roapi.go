package users

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetUsersByUsernames fetches information for users with the given usernames.
// POST https://users.roblox.com/v1/usernames/users
func (r *Resource) GetUsersByUsernames(ctx context.Context, p GetUsersByUsernamesParams) (*types.UsersByUsernameResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var users types.UsersByUsernameResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodPost).
		URL(types.UsersEndpoint + "/v1/usernames/users").
		Result(&users).
		MarshalBody(struct {
			Usernames          []string `json:"usernames"`
			ExcludeBannedUsers bool     `json:"excludeBannedUsers"`
		}{
			Usernames:          p.Usernames,
			ExcludeBannedUsers: p.ExcludeBannedUsers,
		}).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	if err := r.validate.Struct(&users); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &users, nil
}

// GetUsersByUsernamesParams holds the parameters for fetching users by usernames.
type GetUsersByUsernamesParams struct {
	Usernames          []string `json:"usernames"          validate:"required,min=1,max=100"`
	ExcludeBannedUsers bool     `json:"excludeBannedUsers"`
}

// GetUsersByUsernamesBuilder is a builder for GetUsersByUsernamesParams.
type GetUsersByUsernamesBuilder struct {
	params GetUsersByUsernamesParams
}

// NewGetUsersByUsernamesBuilder creates a new GetUsersByUsernamesBuilder with default values.
func NewGetUsersByUsernamesBuilder(usernames ...string) *GetUsersByUsernamesBuilder {
	return &GetUsersByUsernamesBuilder{
		params: GetUsersByUsernamesParams{
			Usernames:          usernames,
			ExcludeBannedUsers: false,
		},
	}
}

// WithUsernames adds multiple usernames to the list.
func (b *GetUsersByUsernamesBuilder) WithUsernames(usernames ...string) *GetUsersByUsernamesBuilder {
	b.params.Usernames = append(b.params.Usernames, usernames...)
	return b
}

// RemoveUsernames removes multiple usernames from the list.
func (b *GetUsersByUsernamesBuilder) RemoveUsernames(usernames ...string) *GetUsersByUsernamesBuilder {
	for _, username := range usernames {
		for i, u := range b.params.Usernames {
			if u == username {
				b.params.Usernames = append(b.params.Usernames[:i], b.params.Usernames[i+1:]...)
				break
			}
		}
	}
	return b
}

// ExcludeBannedUsers sets whether to exclude banned users from the result.
func (b *GetUsersByUsernamesBuilder) ExcludeBannedUsers(exclude bool) *GetUsersByUsernamesBuilder {
	b.params.ExcludeBannedUsers = exclude
	return b
}

// Build returns the GetUsersByUsernamesParams.
func (b *GetUsersByUsernamesBuilder) Build() GetUsersByUsernamesParams {
	return b.params
}
