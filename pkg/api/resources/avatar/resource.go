package avatar

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jaxron/axonet/pkg/client"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// ResourceInterface defines the interface for avatar-related operations.
type ResourceInterface interface {
	GetUserOutfits(ctx context.Context, p UserOutfitsParams) (*types.OutfitResponse, error)
	GetOutfitDetails(ctx context.Context, outfitID int64) (*types.OutfitDetailsResponse, error)
	GetUserAvatar(ctx context.Context, userID int64) (*types.UserAvatarResponse, error)
}

// Ensure Resource implements the ResourceInterface.
var _ ResourceInterface = (*Resource)(nil)

// Resource provides methods for interacting with avatar-related endpoints.
type Resource struct {
	client   *client.Client
	validate *validator.Validate
}

// New creates a new Resource with the specified client and validator.
func New(client *client.Client, validate *validator.Validate) *Resource {
	return &Resource{
		client:   client,
		validate: validate,
	}
}
