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
	GetGameFavoritesCount(ctx context.Context, universeID int64) (*types.GameFavoritesCountResponse, error)
	GetUniverseIDFromPlace(ctx context.Context, placeID int64) (*types.UniverseIDResponse, error)
	GetGamesByUniverseIDs(ctx context.Context, universeIDs []int64) (*types.GameDetailsResponse, error)
	GetGameServers(ctx context.Context, p GameServersParams) (*types.ServerResponse, error)
	GetMultiplePlaceDetails(ctx context.Context, placeIDs []int64) ([]*types.PlaceDetailResponse, error)
	GetUserFavoriteGames(ctx context.Context, p UserFavoriteGamesParams) (*types.GameResponse, error)
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
