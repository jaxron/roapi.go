package games

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jaxron/axonet/pkg/client"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// ResourceInterface defines the methods available for game-related operations.
type ResourceInterface interface {
	GetUserGames(ctx context.Context, p UserGamesParams) (*types.GameResponse, error)
	GetGameFavoritesCount(ctx context.Context, universeID uint64) (*types.GameFavoritesCountResponse, error)
}

// Ensure Resource implements ResourceInterface.
var _ ResourceInterface = (*Resource)(nil)

// Resource handles game-related API operations.
type Resource struct {
	client   *client.Client
	validate *validator.Validate
}

// New creates a new games resource instance.
func New(client *client.Client, validate *validator.Validate) *Resource {
	return &Resource{
		client:   client,
		validate: validate,
	}
}
