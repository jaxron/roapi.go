package types

// CatalogItemType represents the type of item in a catalog request.
type CatalogItemType string

const (
	CatalogItemTypeAsset  CatalogItemType = "Asset"
	CatalogItemTypeBundle CatalogItemType = "Bundle"
)

// ItemDetailsResponse represents the structure of catalog item details returned by the Roblox API.
type ItemDetailsResponse struct {
	Data []*CatalogItem `json:"data" validate:"required,dive"` // List of catalog items
}

// CatalogItem represents a single item returned by the catalog details endpoint.
type CatalogItem struct {
	ID                           int64         `json:"id"                                     validate:"required,min=1"` // Unique identifier for the item
	ItemType                     string        `json:"itemType"                               validate:"required"`       // Item type ("Asset" or "Bundle")
	AssetType                    *int64        `json:"assetType,omitempty"`                                              // Asset type ID (present for assets)
	BundleType                   *int64        `json:"bundleType,omitempty"`                                             // Bundle type ID (present for bundles)
	Name                         string        `json:"name"                                   validate:"required"`       // Item name
	Description                  string        `json:"description"`                                                      // Item description
	ProductID                    *int64        `json:"productId,omitempty"`                                              // Product ID (absent for off-sale items)
	BundledItems                 []BundledItem `json:"bundledItems"`                                                     // Items included in a bundle
	Taxonomy                     []Taxonomy    `json:"taxonomy"`                                                         // Category taxonomy
	ItemStatus                   []string      `json:"itemStatus"`                                                       // Status flags
	ItemRestrictions             []string      `json:"itemRestrictions"`                                                 // Restriction flags
	CreatorHasVerifiedBadge      bool          `json:"creatorHasVerifiedBadge"`                                          // Whether the creator has a verified badge
	CreatorType                  string        `json:"creatorType"                            validate:"required"`       // Creator type ("User" or "Group")
	CreatorTargetID              int64         `json:"creatorTargetId"                        validate:"required,min=1"` // Creator's user or group ID
	CreatorName                  string        `json:"creatorName"                            validate:"required"`       // Creator's name
	Price                        *int64        `json:"price,omitempty"`                                                  // Current price in Robux
	LowestPrice                  *int64        `json:"lowestPrice,omitempty"`                                            // Lowest available price
	LowestResalePrice            *int64        `json:"lowestResalePrice,omitempty"`                                      // Lowest resale price
	PriceStatus                  *string       `json:"priceStatus,omitempty"`                                            // Price status (e.g., "Free")
	UnitsAvailableForConsumption *int64        `json:"unitsAvailableForConsumption,omitempty"`                           // Remaining units available
	FavoriteCount                int64         `json:"favoriteCount"`                                                    // Number of favorites
	OffSaleDeadline              *string       `json:"offSaleDeadline"`                                                  // Off-sale deadline timestamp
	CollectibleItemID            *string       `json:"collectibleItemId,omitempty"`                                      // Collectible item UUID
	TotalQuantity                *int64        `json:"totalQuantity,omitempty"`                                          // Total quantity produced
	SaleLocationType             *string       `json:"saleLocationType,omitempty"`                                       // Where the item can be purchased
	HasResellers                 *bool         `json:"hasResellers,omitempty"`                                           // Whether resellers exist
	IsOffSale                    *bool         `json:"isOffSale,omitempty"`                                              // Whether the item is off sale
	IsRecolorable                *bool         `json:"isRecolorable,omitempty"`                                          // Whether the item can be recolored
	QuantityLimitPerUser         *int64        `json:"quantityLimitPerUser,omitempty"`                                   // Per-user purchase limit
	SupportsHeadShapes           *bool         `json:"supportsHeadShapes,omitempty"`                                     // Whether the item supports head shapes
	TimedOptions                 []TimedOption `json:"timedOptions,omitempty"`                                           // Time-limited purchase options
}

// BundledItem represents a single item within a bundle.
type BundledItem struct {
	Owned              bool   `json:"owned"`                                        // Whether the user owns this bundled item
	ID                 int64  `json:"id"                 validate:"required,min=1"` // Bundled item ID
	Name               string `json:"name"               validate:"required"`       // Bundled item name
	Type               string `json:"type"               validate:"required"`       // Bundled item type
	SupportsHeadShapes bool   `json:"supportsHeadShapes"`                           // Whether the bundled item supports head shapes
	AssetType          int64  `json:"assetType"          validate:"required,min=1"` // Asset type ID
}

// Taxonomy represents a category classification for a catalog item.
type Taxonomy struct {
	TaxonomyID   string `json:"taxonomyId"   validate:"required"` // Unique taxonomy identifier
	TaxonomyName string `json:"taxonomyName" validate:"required"` // Human-readable taxonomy name
}

// TimedOption represents a time-limited purchase option for a catalog item.
type TimedOption struct {
	Days     int64 `json:"days"`     // Duration in days
	Price    int64 `json:"price"`    // Price for this duration
	Selected bool  `json:"selected"` // Whether this option is selected
}
