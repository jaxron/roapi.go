package games

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errs"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetUniverseIDFromPlace fetches the universe ID for a specific place.
// GET https://apis.roblox.com/universes/v1/places/{placeId}/universe
func (r *Resource) GetUniverseIDFromPlace(ctx context.Context, placeID int64) (*types.UniverseIDResponse, error) {
	if err := r.validate.Var(placeID, "required,gt=0"); err != nil {
		return nil, fmt.Errorf("%w: %w", errs.ErrInvalidRequest, err)
	}

	var result types.UniverseIDResponse

	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("https://apis.roblox.com/universes/v1/places/%d/universe", placeID)).
		Result(&result).
		Do(ctx)
	if err != nil {
		return nil, errs.HandleAPIError(resp, err)
	}

	defer func() { _ = resp.Body.Close() }()

	if err := r.validate.Struct(&result); err != nil {
		return nil, fmt.Errorf("%w: %w", errs.ErrInvalidResponse, err)
	}

	return &result, nil
}
