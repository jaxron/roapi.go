package friends

import (
	"context"

	"github.com/jaxron/roapi.go/pkg/api/models"
	"github.com/jaxron/roapi.go/pkg/client"
)

const (
	SortOrderAsc  = "Asc"
	SortOrderDesc = "Desc"
)

const FriendsEndpoint = "https://friends.roblox.com"

// ServiceInterface defines the interface for friend-related operations.
type ServiceInterface interface {
	GetFriends(ctx context.Context, userID uint64) ([]models.UserResponse, error)
	GetFriendCount(ctx context.Context, userID uint64) (uint64, error)
	FindFriends(ctx context.Context, b *FindFriendsBuilder) (*models.FriendPageResponse, error)
}

// Ensure Service implements the ServiceInterface.
var _ ServiceInterface = (*Service)(nil)

// Service provides methods for interacting with friend-related endpoints.
type Service struct {
	Client *client.Client
}

// NewService creates a new Service with the specified version.
func NewService(client *client.Client) *Service {
	return &Service{
		Client: client,
	}
}
