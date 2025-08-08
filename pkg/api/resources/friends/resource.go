package friends

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jaxron/axonet/pkg/client"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// ResourceInterface defines the interface for friend-related operations.
type ResourceInterface interface {
	GetFriends(ctx context.Context, userID int64) (*types.FriendsResponse, error)
	GetFriendCount(ctx context.Context, userID int64) (int64, error)
	FindFriends(ctx context.Context, params FindFriendsParams) (*types.FriendPageResponse, error)
	SearchFriends(ctx context.Context, params SearchFriendsParams) (*types.FriendPageResponse, error)
	GetFollowers(ctx context.Context, params GetFollowersParams) (*types.FollowerPageResponse, error)
	GetFollowerCount(ctx context.Context, userID int64) (int64, error)
	GetFollowings(ctx context.Context, params GetFollowingsParams) (*types.FollowingPageResponse, error)
	GetFollowingCount(ctx context.Context, userID int64) (int64, error)
}

// Ensure Resource implements the ResourceInterface.
var _ ResourceInterface = (*Resource)(nil)

// Resource provides methods for interacting with friend-related endpoints.
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
