package types

import "time"

// GameResponse represents the structure of a game response returned by the Roblox API.
type GameResponse struct {
	PreviousPageCursor *string `json:"previousPageCursor" validate:"omitempty,base64"` // Cursor for the previous page of results (if any)
	NextPageCursor     *string `json:"nextPageCursor"     validate:"omitempty,base64"` // Cursor for the next page of results (if any)
	Data               []Game  `json:"data"               validate:"required,dive"`    // List of games
}

// Game represents a single game returned by the Roblox API.
type Game struct {
	ID          uint64    `json:"id"          validate:"required,min=1"` // Unique identifier for the game
	Name        string    `json:"name"        validate:"required,min=1"` // Name of the game
	Description string    `json:"description"`                           // Description of the game
	Creator     Creator   `json:"creator"     validate:"required"`       // Creator information
	RootPlace   Place     `json:"rootPlace"   validate:"required"`       // Root place information
	Created     time.Time `json:"created"     validate:"required"`       // When the game was created
	Updated     time.Time `json:"updated"     validate:"required"`       // When the game was last updated
	PlaceVisits uint64    `json:"placeVisits" validate:"min=0"`          // Number of visits to the game
}

// Creator represents the creator of a game.
type Creator struct {
	ID   uint64 `json:"id"   validate:"required,min=1"`            // Creator's unique identifier
	Type string `json:"type" validate:"required,oneof=User Group"` // Type of creator (User/Group)
}

// Place represents a place within a game.
type Place struct {
	ID   uint64 `json:"id"   validate:"required,min=1"` // Place's unique identifier
	Type string `json:"type" validate:"required"`       // Type of place
}

// GameFavoritesCountResponse represents the favorites count for a game.
type GameFavoritesCountResponse struct {
	FavoritesCount uint64 `json:"favoritesCount" validate:"min=0"` // Number of times the game has been favorited
}
