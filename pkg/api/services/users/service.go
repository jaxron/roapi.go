package users

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jaxron/axonet/pkg/client"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// ServiceInterface defines the interface for user-related operations.
type ServiceInterface interface {
	GetUserByID(ctx context.Context, userID uint64) (*types.UserByIDResponse, error)
	GetAuthUserInfo(ctx context.Context) (*types.AuthUserResponse, error)
	GetUsersByUsernames(ctx context.Context, params GetUsersByUsernamesParams) ([]types.UserByUsernameResponse, error)
	GetUsersByIDs(ctx context.Context, params UsersByIDsParams) ([]types.VerifiedBadgeUserResponse, error)
	GetUsernameHistory(ctx context.Context, params UsernameHistoryParams) (*types.UsernameHistoryPageResponse, error)
	SearchUsers(ctx context.Context, params SearchUsersParams) (*types.UserSearchPageResponse, error)
}

// Ensure Service implements the ServiceInterface.
var _ ServiceInterface = (*Service)(nil)

// Service provides methods for interacting with user-related endpoints.
type Service struct {
	client   *client.Client
	validate *validator.Validate
}

// NewService creates a new Service with the specified version.
func NewService(client *client.Client, validate *validator.Validate) *Service {
	return &Service{
		client:   client,
		validate: validate,
	}
}
