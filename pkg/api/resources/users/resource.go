package users

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jaxron/axonet/pkg/client"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// ResourceInterface defines the interface for user-related operations.
type ResourceInterface interface {
	GetUserByID(ctx context.Context, userID uint64) (*types.UserByIDResponse, error)
	GetAuthUserInfo(ctx context.Context) (*types.AuthUserResponse, error)
	GetUsersByUsernames(ctx context.Context, params GetUsersByUsernamesParams) (*types.UsersByUsernameResponse, error)
	GetUsersByIDs(ctx context.Context, params UsersByIDsParams) (*types.UsersByIDsResponse, error)
	GetUsernameHistory(ctx context.Context, params UsernameHistoryParams) (*types.UsernameHistoryPageResponse, error)
	SearchUsers(ctx context.Context, params SearchUsersParams) (*types.UserSearchPageResponse, error)
}

// Ensure Resource implements the ResourceInterface.
var _ ResourceInterface = (*Resource)(nil)

// Resource provides methods for interacting with user-related endpoints.
type Resource struct {
	client   *client.Client
	validate *validator.Validate
}

// New creates a new Resource with the specified version.
func New(client *client.Client, validate *validator.Validate) *Resource {
	return &Resource{
		client:   client,
		validate: validate,
	}
}
