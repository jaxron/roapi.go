package types

// Outfit represents a single outfit in the user's outfits.
type Outfit struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	IsEditable bool   `json:"isEditable"`
	OutfitType string `json:"outfitType"`
}
