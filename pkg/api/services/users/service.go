package users

import (
	"context"

	"github.com/jaxron/roapi.go/pkg/api/models"
	"github.com/jaxron/roapi.go/pkg/client"
)

const UsersEndpoint = "https://users.roblox.com"

const (
	SortOrderAsc  = "Asc"
	SortOrderDesc = "Desc"

	DefaultLimit = 10
)

// ServiceInterface defines the interface for user-related operations.
type ServiceInterface interface {
	GetUserInfo(ctx context.Context, userID uint64) (*models.UserInfo, error)
	GetAuthUserInfo(ctx context.Context) (*models.AuthUserInfo, error)
	GetUsersByUsernames(ctx context.Context, b *UsersByUsernamesBuilder) (*models.UsersByUsernames, error)
	GetUsersByIDs(ctx context.Context, b *UsersByIDsBuilder) (*models.UsersByIDs, error)
	GetUsernameHistory(ctx context.Context, b *UsernameHistoryBuilder) (*models.UsernameHistory, error)
	SearchUser(ctx context.Context, b *SearchUserBuilder) (*models.SearchResult, error)
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
