package catalog

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errs"
	"github.com/jaxron/roapi.go/pkg/api/middleware/auth"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetItemDetails fetches catalog details for a batch of items.
// POST https://catalog.roblox.com/v1/catalog/items/details
func (r *Resource) GetItemDetails(ctx context.Context, p GetItemDetailsParams) (*types.ItemDetailsResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errs.ErrInvalidRequest, err)
	}

	ctx = context.WithValue(ctx, auth.KeyAddCookie, true)
	ctx = context.WithValue(ctx, auth.KeyAddToken, true)

	var result types.ItemDetailsResponse

	resp, err := r.client.NewRequest().
		Method(http.MethodPost).
		URL(types.CatalogEndpoint + "/v1/catalog/items/details").
		Result(&result).
		MarshalBody(struct {
			Items []CatalogItemRequest `json:"items"`
		}{
			Items: p.Items,
		}).
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

// CatalogItemRequest represents a single item in a catalog details request.
type CatalogItemRequest struct {
	ItemType types.CatalogItemType `json:"itemType" validate:"required,oneof=Asset Bundle"` // Item type ("Asset" or "Bundle")
	ID       int64                 `json:"id"       validate:"required,min=1"`              // Item ID
}

// GetItemDetailsParams holds the parameters for fetching catalog item details.
type GetItemDetailsParams struct {
	Items []CatalogItemRequest `json:"items" validate:"required,min=1,max=120,dive"`
}

// GetItemDetailsBuilder is a builder for GetItemDetailsParams.
type GetItemDetailsBuilder struct {
	params GetItemDetailsParams
}

// NewGetItemDetailsBuilder creates a new GetItemDetailsBuilder with the given items.
func NewGetItemDetailsBuilder(items ...CatalogItemRequest) *GetItemDetailsBuilder {
	return &GetItemDetailsBuilder{
		params: GetItemDetailsParams{
			Items: items,
		},
	}
}

// WithItems adds multiple items to the request.
func (b *GetItemDetailsBuilder) WithItems(items ...CatalogItemRequest) *GetItemDetailsBuilder {
	b.params.Items = append(b.params.Items, items...)
	return b
}

// RemoveItems removes items by ID from the request.
func (b *GetItemDetailsBuilder) RemoveItems(ids ...int64) *GetItemDetailsBuilder {
	idSet := make(map[int64]struct{}, len(ids))
	for _, id := range ids {
		idSet[id] = struct{}{}
	}

	filtered := b.params.Items[:0]
	for _, item := range b.params.Items {
		if _, found := idSet[item.ID]; !found {
			filtered = append(filtered, item)
		}
	}

	b.params.Items = filtered

	return b
}

// Build returns the GetItemDetailsParams.
func (b *GetItemDetailsBuilder) Build() GetItemDetailsParams {
	return b.params
}
