package friends

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jaxron/axonet/pkg/client"
	"github.com/jaxron/roapi.go/pkg/api/models"
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
	FindFriends(ctx context.Context, params FindFriendsParams) (*models.FriendPageResponse, error)
	SearchFriends(ctx context.Context, params SearchFriendsParams) (*models.FriendPageResponse, error)
	GetFollowers(ctx context.Context, params GetFollowersParams) (*models.FollowerPageResponse, error)
	GetFollowerCount(ctx context.Context, userID uint64) (uint64, error)
}

// Ensure Service implements the ServiceInterface.
var _ ServiceInterface = (*Service)(nil)

// Service provides methods for interacting with friend-related endpoints.
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
