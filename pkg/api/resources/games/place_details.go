package games

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/middleware/auth"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetMultiplePlaceDetails fetches details for multiple places simultaneously.
// GET https://games.roblox.com/v1/games/multiget-place-details?placeIds={placeIds}
func (r *Resource) GetMultiplePlaceDetails(ctx context.Context, placeIDs []int64) ([]*types.PlaceDetailResponse, error) {
	if err := r.validate.Var(placeIDs, "required,min=1,max=100,dive,gt=0"); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	ctx = context.WithValue(ctx, auth.KeyAddCookie, true)

	// Convert place IDs to strings for query parameters
	ids := make([]string, len(placeIDs))
	for i, id := range placeIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}

	// Create request with multiple placeIds query parameters
	req := r.client.NewRequest().
		Method(http.MethodGet).
		URL(types.GamesEndpoint + "/v1/games/multiget-place-details")

	// Add each placeId as a separate query parameter
	for _, id := range ids {
		req.Query("placeIds", id)
	}

	var result []*types.PlaceDetailResponse

	resp, err := req.Result(&result).Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}

	defer func() { _ = resp.Body.Close() }()

	if err := r.validate.Var(result, "required,dive"); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return result, nil
}
