package games

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetGameFavoritesCount fetches the number of favorites for a specific game.
// GET https://games.roblox.com/v1/games/{universeId}/favorites/count
func (r *Resource) GetGameFavoritesCount(ctx context.Context, universeID uint64) (*types.GameFavoritesCountResponse, error) {
	if err := r.validate.Var(universeID, "required,gt=0"); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var result types.GameFavoritesCountResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/games/%d/favorites/count", types.GamesEndpoint, universeID)).
		Result(&result).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	if err := r.validate.Struct(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
