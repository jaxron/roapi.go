package types

import "time"

type GameResponse struct {
	PreviousPageCursor *string `json:"previousPageCursor"` // Cursor for the previous page of results (if any)
	NextPageCursor     *string `json:"nextPageCursor"`     // Cursor for the next page of results (if any)
	Data               []Game  `json:"data"`               // List of games
}

// Game represents a single game returned by the Roblox API.
type Game struct {
	ID          uint64    `json:"id"`          // Unique identifier for the game
	Name        string    `json:"name"`        // Name of the game
	Description string    `json:"description"` // Description of the game
	Creator     Creator   `json:"creator"`     // Creator information
	RootPlace   Place     `json:"rootPlace"`   // Root place information
	Created     time.Time `json:"created"`     // When the game was created
	Updated     time.Time `json:"updated"`     // When the game was last updated
	PlaceVisits uint64    `json:"placeVisits"` // Number of visits to the game
}

// Creator represents the creator of a game.
type Creator struct {
	ID   uint64 `json:"id"`   // Creator's unique identifier
	Type string `json:"type"` // Type of creator (User/Group)
}

// Place represents a place within a game.
type Place struct {
	ID   uint64 `json:"id"`   // Place's unique identifier
	Type string `json:"type"` // Type of place
}
