package types

// OutfitResponse represents the structure of a user's outfits returned by the Roblox API.
type OutfitResponse struct {
	Data            []*Outfit `json:"data"            validate:"required,dive"` // List of outfits
	PaginationToken string    `json:"paginationToken"`                          // Token for pagination (if any)
}

// Outfit represents a single outfit in the user's outfits.
type Outfit struct {
	ID         int64  `json:"id"         validate:"required,min=1"` // Unique identifier for the outfit
	Name       string `json:"name"       validate:"omitempty"`      // Name of the outfit
	IsEditable bool   `json:"isEditable"`                           // Whether the outfit can be edited
	OutfitType string `json:"outfitType" validate:"required"`       // Type of outfit
}

// OutfitDetailsResponse represents the detailed information about a specific outfit.
type OutfitDetailsResponse struct {
	ID               int64       `json:"id"               validate:"required,min=1"` // Unique identifier for the outfit
	Name             string      `json:"name"             validate:"omitempty"`      // Name of the outfit
	Assets           []*AssetV2  `json:"assets"           validate:"required,dive"`  // List of assets in the outfit
	BodyColors       BodyColors3 `json:"bodyColor3s"      validate:"required"`       // Body part colors
	Scale            ScaleModel  `json:"scale"            validate:"required"`       // Avatar scaling information
	PlayerAvatarType string      `json:"playerAvatarType" validate:"required"`       // R6 or R15
	OutfitType       string      `json:"outfitType"       validate:"required"`       // Type of outfit
	IsEditable       bool        `json:"isEditable"`                                 // Whether the outfit can be edited
	UniverseID       *int64      `json:"universeId"`                                 // Universe ID if created in-experience
	ModerationStatus *string     `json:"moderationStatus"`                           // Moderation status for in-experience outfits
	BundleID         *int64      `json:"bundleId"`                                   // Bundle ID for in-experience outfits
}

// AssetV2 represents an asset with additional version information.
type AssetV2 struct {
	ID               int64        `json:"id"               validate:"required,min=1"` // Asset ID
	Name             string       `json:"name"             validate:"required"`       // Asset name
	AssetType        AssetType    `json:"assetType"        validate:"required"`       // Asset type information
	CurrentVersionID int64        `json:"currentVersionId" validate:"required,min=1"` // Current version ID
	Meta             *AssetMetaV1 `json:"meta,omitempty"`                             // Optional metadata
}

// AssetType represents the type information for an asset.
type AssetType struct {
	ID   ItemAssetType `json:"id"   validate:"required,min=1"` // Asset type ID
	Name string        `json:"name" validate:"required"`       // Asset type name
}

// AssetMetaV1 represents metadata for an asset.
type AssetMetaV1 struct {
	Order     *int32        `json:"order,omitempty"`     // Layered clothing order
	Puffiness *float32      `json:"puffiness,omitempty"` // Layered clothing puffiness
	Position  *AssetVector3 `json:"position,omitempty"`  // Asset position
	Rotation  *AssetVector3 `json:"rotation,omitempty"`  // Asset rotation
	Scale     *AssetVector3 `json:"scale,omitempty"`     // Asset scale
	Version   int32         `json:"version"`             // Meta model version
}

// AssetVector3 represents a 3D vector for asset positioning.
type AssetVector3 struct {
	X float32 `json:"X"` // X coordinate
	Y float32 `json:"Y"` // Y coordinate
	Z float32 `json:"Z"` // Z coordinate
}

// BodyColors3 represents the RGB hex colors for each body part.
type BodyColors3 struct {
	HeadColor3     string `json:"headColor3"     validate:"required"` // Head color in hex
	TorsoColor3    string `json:"torsoColor3"    validate:"required"` // Torso color in hex
	RightArmColor3 string `json:"rightArmColor3" validate:"required"` // Right arm color in hex
	LeftArmColor3  string `json:"leftArmColor3"  validate:"required"` // Left arm color in hex
	RightLegColor3 string `json:"rightLegColor3" validate:"required"` // Right leg color in hex
	LeftLegColor3  string `json:"leftLegColor3"  validate:"required"` // Left leg color in hex
}

// ScaleModel represents the scaling information for an avatar.
type ScaleModel struct {
	Height     float64 `json:"height"     validate:"required"` // Height scale
	Width      float64 `json:"width"      validate:"required"` // Width scale
	Head       float64 `json:"head"       validate:"required"` // Head scale
	Depth      float64 `json:"depth"      validate:"required"` // Depth scale
	Proportion float64 `json:"proportion"`                     // Proportion scale
	BodyType   float64 `json:"bodyType"`                       // Body type scale
}

// UserAvatarResponse represents the structure of a user's avatar details.
type UserAvatarResponse struct {
	Scales              ScaleModel   `json:"scales"              validate:"required"`      // Avatar scaling information
	PlayerAvatarType    string       `json:"playerAvatarType"    validate:"required"`      // R6 or R15
	BodyColors          BodyColors3  `json:"bodyColor3s"         validate:"required"`      // Body part colors
	Assets              []*AssetV2   `json:"assets"              validate:"required,dive"` // List of assets worn on the character
	DefaultShirtApplied bool         `json:"defaultShirtApplied"`                          // Whether default shirt is applied
	DefaultPantsApplied bool         `json:"defaultPantsApplied"`                          // Whether default pants are applied
	Emotes              []EmoteModel `json:"emotes"              validate:"required,dive"` // List of equipped emotes
}

// EmoteModel represents an emote equipped on a user's avatar.
type EmoteModel struct {
	AssetID   int64  `json:"assetId"   validate:"required,min=1"` // The asset ID of the emote
	AssetName string `json:"assetName" validate:"required"`       // The name of the emote
	Position  int32  `json:"position"  validate:"required"`       // The position the emote is equipped to
}
