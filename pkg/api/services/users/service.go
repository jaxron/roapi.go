package users

import (
	"context"

	"github.com/jaxron/roapi.go/pkg/api/client"
	"github.com/jaxron/roapi.go/pkg/api/models"
)

const UsersEndpoint = "https://users.roblox.com"

const (
	SortOrderAsc  = "Asc"
	SortOrderDesc = "Desc"

	DefaultLimit = 10
)

// ServiceInterface defines the interface for user-related operations.
type ServiceInterface interface {
	GetUserByID(ctx context.Context, userID uint64) (*models.UserByIDResponse, error)
	GetAuthUserInfo(ctx context.Context) (*models.AuthUserResponse, error)
	GetUsersByUsernames(ctx context.Context, b *UsersByUsernamesBuilder) ([]models.UserByUsernameResponse, error)
	GetUsersByIDs(ctx context.Context, b *UsersByIDsBuilder) ([]models.VerifiedBadgeUserResponse, error)
	GetUsernameHistory(ctx context.Context, b *UsernameHistoryBuilder) (*models.UsernameHistoryPageResponse, error)
	SearchUsers(ctx context.Context, b *SearchUsersBuilder) (*models.UserSearchPageResponse, error)
}

// Ensure Service implements the ServiceInterface.
var _ ServiceInterface = (*Service)(nil)

// Service provides methods for interacting with user-related endpoints.
type Service struct {
	Client *client.Client
}

// NewService creates a new Service with the specified version.
func NewService(client *client.Client) *Service {
	return &Service{
		Client: client,
	}
}
