package types

import "time"

// InventoryAssetResponse represents the structure of an inventory asset list response.
type InventoryAssetResponse struct {
	PreviousPageCursor *string          `json:"previousPageCursor" validate:"omitempty,base64"` // Cursor for the previous page of results (if any)
	NextPageCursor     *string          `json:"nextPageCursor"     validate:"omitempty,base64"` // Cursor for the next page of results (if any)
	Data               []InventoryAsset `json:"data"               validate:"required,dive"`    // List of inventory assets
}

// InventoryAsset represents a single asset in a user's inventory.
type InventoryAsset struct {
	AssetID   int64     `json:"assetId"   validate:"required,min=1"` // Unique identifier for the asset
	Name      string    `json:"name"      validate:"required,min=1"` // Name of the asset
	AssetType string    `json:"assetType" validate:"required,min=1"` // Type of the asset (e.g., "Hat", "Shirt")
	Created   time.Time `json:"created"   validate:"required"`       // When the asset was created/acquired
}

// ItemAssetType represents the type of asset in the inventory.
//
//go:generate go tool enumer -type=ItemAssetType -trimprefix=ItemAssetType
type ItemAssetType int64

const (
	ItemAssetTypeImage               ItemAssetType = 1
	ItemAssetTypeTShirt              ItemAssetType = 2
	ItemAssetTypeAudio               ItemAssetType = 3
	ItemAssetTypeMesh                ItemAssetType = 4
	ItemAssetTypeLua                 ItemAssetType = 5
	ItemAssetTypeHat                 ItemAssetType = 8
	ItemAssetTypePlace               ItemAssetType = 9
	ItemAssetTypeModel               ItemAssetType = 10
	ItemAssetTypeShirt               ItemAssetType = 11
	ItemAssetTypePants               ItemAssetType = 12
	ItemAssetTypeDecal               ItemAssetType = 13
	ItemAssetTypeHead                ItemAssetType = 17
	ItemAssetTypeFace                ItemAssetType = 18
	ItemAssetTypeGear                ItemAssetType = 19
	ItemAssetTypeBadge               ItemAssetType = 21
	ItemAssetTypeAnimation           ItemAssetType = 24
	ItemAssetTypeTorso               ItemAssetType = 27
	ItemAssetTypeRightArm            ItemAssetType = 28
	ItemAssetTypeLeftArm             ItemAssetType = 29
	ItemAssetTypeLeftLeg             ItemAssetType = 30
	ItemAssetTypeRightLeg            ItemAssetType = 31
	ItemAssetTypePackage             ItemAssetType = 32
	ItemAssetTypeGamePass            ItemAssetType = 34
	ItemAssetTypePlugin              ItemAssetType = 38
	ItemAssetTypeMeshPart            ItemAssetType = 40
	ItemAssetTypeHairAccessory       ItemAssetType = 41
	ItemAssetTypeFaceAccessory       ItemAssetType = 42
	ItemAssetTypeNeckAccessory       ItemAssetType = 43
	ItemAssetTypeShoulderAccessory   ItemAssetType = 44
	ItemAssetTypeFrontAccessory      ItemAssetType = 45
	ItemAssetTypeBackAccessory       ItemAssetType = 46
	ItemAssetTypeWaistAccessory      ItemAssetType = 47
	ItemAssetTypeClimbAnimation      ItemAssetType = 48
	ItemAssetTypeDeathAnimation      ItemAssetType = 49
	ItemAssetTypeFallAnimation       ItemAssetType = 50
	ItemAssetTypeIdleAnimation       ItemAssetType = 51
	ItemAssetTypeJumpAnimation       ItemAssetType = 52
	ItemAssetTypeRunAnimation        ItemAssetType = 53
	ItemAssetTypeSwimAnimation       ItemAssetType = 54
	ItemAssetTypeWalkAnimation       ItemAssetType = 55
	ItemAssetTypePoseAnimation       ItemAssetType = 56
	ItemAssetTypeEarAccessory        ItemAssetType = 57
	ItemAssetTypeEyeAccessory        ItemAssetType = 58
	ItemAssetTypeEmoteAnimation      ItemAssetType = 61
	ItemAssetTypeVideo               ItemAssetType = 62
	ItemAssetTypeTShirtAccessory     ItemAssetType = 64
	ItemAssetTypeShirtAccessory      ItemAssetType = 65
	ItemAssetTypePantsAccessory      ItemAssetType = 66
	ItemAssetTypeJacketAccessory     ItemAssetType = 67
	ItemAssetTypeSweaterAccessory    ItemAssetType = 68
	ItemAssetTypeShortsAccessory     ItemAssetType = 69
	ItemAssetTypeLeftShoeAccessory   ItemAssetType = 70
	ItemAssetTypeRightShoeAccessory  ItemAssetType = 71
	ItemAssetTypeDressSkirtAccessory ItemAssetType = 72
	ItemAssetTypeFontFamily          ItemAssetType = 73
	ItemAssetTypeEyebrowAccessory    ItemAssetType = 76
	ItemAssetTypeEyelashAccessory    ItemAssetType = 77
	ItemAssetTypeMoodAnimation       ItemAssetType = 78
	ItemAssetTypeDynamicHead         ItemAssetType = 79
)
