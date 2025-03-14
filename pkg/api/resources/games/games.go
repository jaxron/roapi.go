package games

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetGamesByUniverseIDs fetches game details for the specified universe IDs.
// GET https://games.roblox.com/v1/games?universeIds={universeIds}
func (r *Resource) GetGamesByUniverseIDs(ctx context.Context, universeIDs []uint64) (*types.GameDetailsResponse, error) {
	if err := r.validate.Var(universeIDs, "required,min=1,max=100,dive,gt=0"); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	// Convert universe IDs to strings and join them
	ids := make([]string, len(universeIDs))
	for i, id := range universeIDs {
		ids[i] = strconv.FormatUint(id, 10)
	}

	var result types.GameDetailsResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(types.GamesEndpoint+"/v1/games").
		Query("universeIds", strings.Join(ids, ",")).
		Result(&result).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	if err := r.validate.Struct(&result); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &result, nil
}
