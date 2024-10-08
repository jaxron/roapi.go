package friends

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jaxron/axonet/pkg/client"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// ServiceInterface defines the interface for friend-related operations.
type ServiceInterface interface {
	GetFriends(ctx context.Context, userID uint64) ([]types.UserResponse, error)
	GetFriendCount(ctx context.Context, userID uint64) (uint64, error)
	FindFriends(ctx context.Context, params FindFriendsParams) (*types.FriendPageResponse, error)
	SearchFriends(ctx context.Context, params SearchFriendsParams) (*types.FriendPageResponse, error)
	GetFollowers(ctx context.Context, params GetFollowersParams) (*types.FollowerPageResponse, error)
	GetFollowerCount(ctx context.Context, userID uint64) (uint64, error)
	GetFollowings(ctx context.Context, params GetFollowingsParams) (*types.FollowingPageResponse, error)
	GetFollowingCount(ctx context.Context, userID uint64) (uint64, error)
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
