package inventory

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetUserAssetsParams holds the parameters for fetching user assets.
type GetUserAssetsParams struct {
	UserID                  uint64                `json:"-" validate:"required,min=1"`
	AssetTypes              []types.ItemAssetType `json:"-" validate:"required,min=1,dive,required"`
	FilterDisapprovedAssets bool                  `json:"-"`
	ShowApprovedOnly        bool                  `json:"-"`
	Limit                   uint64                `json:"-" validate:"required,oneof=10 25 50 100"`
	SortOrder               types.SortOrder       `json:"-" validate:"required,oneof=Asc Desc"`
	Cursor                  string                `json:"-" validate:"omitempty,base64"`
}

// GetUserAssets fetches assets from a user's inventory.
// GET https://inventory.roblox.com/v2/users/{userId}/inventory
func (r *Resource) GetUserAssets(ctx context.Context, p GetUserAssetsParams) (*types.InventoryAssetResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	// Convert asset types to comma-separated string of numeric IDs
	assetTypeIDs := make([]string, len(p.AssetTypes))
	for i, t := range p.AssetTypes {
		assetTypeIDs[i] = t.String()
	}

	var result types.InventoryAssetResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v2/users/%d/inventory", types.InventoryEndpoint, p.UserID)).
		Query("assetTypes", strings.Join(assetTypeIDs, ",")).
		Query("filterDisapprovedAssets", strconv.FormatBool(p.FilterDisapprovedAssets)).
		Query("showApprovedOnly", strconv.FormatBool(p.ShowApprovedOnly)).
		Query("limit", strconv.FormatUint(p.Limit, 10)).
		Query("sortOrder", string(p.SortOrder)).
		Query("cursor", p.Cursor).
		Result(&result).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if err := r.validate.Struct(&result); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &result, nil
}

// GetUserAssetsBuilder is a builder for GetUserAssetsParams.
type GetUserAssetsBuilder struct {
	params GetUserAssetsParams
}

// NewGetUserAssetsBuilder creates a new GetUserAssetsBuilder with default values.
func NewGetUserAssetsBuilder(userID uint64, assetTypes ...types.ItemAssetType) *GetUserAssetsBuilder {
	return &GetUserAssetsBuilder{
		params: GetUserAssetsParams{
			UserID:                  userID,
			AssetTypes:              assetTypes,
			FilterDisapprovedAssets: false,
			ShowApprovedOnly:        false,
			Limit:                   10,
			SortOrder:               types.SortOrderAsc,
			Cursor:                  "",
		},
	}
}

// WithFilterDisapprovedAssets sets whether to filter disapproved assets.
func (b *GetUserAssetsBuilder) WithFilterDisapprovedAssets(filter bool) *GetUserAssetsBuilder {
	b.params.FilterDisapprovedAssets = filter
	return b
}

// WithShowApprovedOnly sets whether to show only approved assets.
func (b *GetUserAssetsBuilder) WithShowApprovedOnly(show bool) *GetUserAssetsBuilder {
	b.params.ShowApprovedOnly = show
	return b
}

// WithLimit sets the maximum number of results to return.
func (b *GetUserAssetsBuilder) WithLimit(limit uint64) *GetUserAssetsBuilder {
	b.params.Limit = limit
	return b
}

// WithSortOrder sets the sort order of the results.
func (b *GetUserAssetsBuilder) WithSortOrder(order types.SortOrder) *GetUserAssetsBuilder {
	b.params.SortOrder = order
	return b
}

// WithCursor sets the cursor for pagination.
func (b *GetUserAssetsBuilder) WithCursor(cursor string) *GetUserAssetsBuilder {
	b.params.Cursor = cursor
	return b
}

// Build returns the GetUserAssetsParams.
func (b *GetUserAssetsBuilder) Build() GetUserAssetsParams {
	return b.params
}
