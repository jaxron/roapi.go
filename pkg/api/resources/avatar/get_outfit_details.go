package avatar

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetOutfitDetails fetches the details of a specific outfit.
// GET https://avatar.roblox.com/v3/outfits/{outfitId}/details
func (r *Resource) GetOutfitDetails(ctx context.Context, outfitID uint64) (*types.OutfitDetailsResponse, error) {
	if err := r.validate.Var(outfitID, "required,gt=0"); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var outfitDetails types.OutfitDetailsResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v3/outfits/%d/details", types.AvatarEndpoint, outfitID)).
		Result(&outfitDetails).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if err := r.validate.Struct(&outfitDetails); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &outfitDetails, nil
}
