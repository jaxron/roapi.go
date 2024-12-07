package types

// OutfitResponse represents the structure of a user's outfits returned by the Roblox API.
type OutfitResponse struct {
	Data            []Outfit `json:"data"            validate:"required,dive"` // List of outfits
	PaginationToken string   `json:"paginationToken"`                          // Token for pagination (if any)
}

// Outfit represents a single outfit in the user's outfits.
type Outfit struct {
	ID         uint64 `json:"id"         validate:"required,min=1"` // Unique identifier for the outfit
	Name       string `json:"name"       validate:"omitempty"`      // Name of the outfit
	IsEditable bool   `json:"isEditable"`                           // Whether the outfit can be edited
	OutfitType string `json:"outfitType" validate:"required"`       // Type of outfit
}
