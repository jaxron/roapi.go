package users

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jaxron/axonet/pkg/client"
	"github.com/jaxron/roapi.go/pkg/api/models"
)

const UsersEndpoint = "https://users.roblox.com"

const (
	SortOrderAsc  = "Asc"
	SortOrderDesc = "Desc"
)

// ServiceInterface defines the interface for user-related operations.
type ServiceInterface interface {
	GetUserByID(ctx context.Context, userID uint64) (*models.UserByIDResponse, error)
	GetAuthUserInfo(ctx context.Context) (*models.AuthUserResponse, error)
	GetUsersByUsernames(ctx context.Context, params GetUsersByUsernamesParams) ([]models.UserByUsernameResponse, error)
	GetUsersByIDs(ctx context.Context, params UsersByIDsParams) ([]models.VerifiedBadgeUserResponse, error)
	GetUsernameHistory(ctx context.Context, params UsernameHistoryParams) (*models.UsernameHistoryPageResponse, error)
	SearchUsers(ctx context.Context, params SearchUsersParams) (*models.UserSearchPageResponse, error)
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
