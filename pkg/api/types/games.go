package types

import "time"

// GameResponse represents the structure of a game response returned by the Roblox API.
type GameResponse struct {
	PreviousPageCursor *string `json:"previousPageCursor" validate:"omitempty"`     // Cursor for the previous page of results (if any)
	NextPageCursor     *string `json:"nextPageCursor"     validate:"omitempty"`     // Cursor for the next page of results (if any)
	Data               []Game  `json:"data"               validate:"required,dive"` // List of games
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

// UniverseIDResponse represents the response containing a universe ID.
type UniverseIDResponse struct {
	UniverseID uint64 `json:"universeId" validate:"required,min=1"` // Universe ID associated with the place
}

// GameCreator represents the extended information about the creator of a game.
type GameCreator struct {
	ID               uint64 `json:"id"               validate:"required,min=1"` // The game creator id
	Name             string `json:"name"             validate:"required"`       // The game creator name
	Type             string `json:"type"             validate:"required"`       // The game creator type
	IsRNVAccount     bool   `json:"isRNVAccount"`                               // The game creator account is Luobu Real Name Verified
	HasVerifiedBadge bool   `json:"hasVerifiedBadge"`                           // Builder verified badge status
}

// GameDetailResponse represents detailed information about a game.
type GameDetailResponse struct {
	ID                        uint64      `json:"id"                        validate:"required,min=1"` // The game universe id
	RootPlaceID               uint64      `json:"rootPlaceId"               validate:"required,min=1"` // The game root place id
	Name                      string      `json:"name"                      validate:"required"`       // The game name
	Description               string      `json:"description"`                                         // The game description
	SourceName                string      `json:"sourceName"`                                          // The game name in the source language
	SourceDescription         string      `json:"sourceDescription"`                                   // The game description in the source language
	Creator                   GameCreator `json:"creator"                   validate:"required"`       // Information about the game creator
	Price                     *uint64     `json:"price"`                                               // The game paid access price
	AllowedGearGenres         []string    `json:"allowedGearGenres"         validate:"required"`       // List of allowed gear genres
	AllowedGearCategories     []string    `json:"allowedGearCategories"     validate:"required"`       // List of allowed gear categories
	IsGenreEnforced           bool        `json:"isGenreEnforced"`                                     // Whether the game must specify a genre
	CopyingAllowed            bool        `json:"copyingAllowed"`                                      // Whether the game allows place to be copied
	Playing                   uint64      `json:"playing"`                                             // Current player count of the game
	Visits                    uint64      `json:"visits"`                                              // The total visits to the game
	MaxPlayers                uint32      `json:"maxPlayers"`                                          // The game max players
	Created                   time.Time   `json:"created"                   validate:"required"`       // The game created time
	Updated                   time.Time   `json:"updated"                   validate:"required"`       // The game updated time
	StudioAccessToApisAllowed bool        `json:"studioAccessToApisAllowed"`                           // The setting of IsStudioAccessToApisAllowed of the universe
	CreateVipServersAllowed   bool        `json:"createVipServersAllowed"`                             // Whether the VIP servers are allowed to be created
	UniverseAvatarType        string      `json:"universeAvatarType"        validate:"required"`       // Avatar type. Possible values are MorphToR6, MorphToR15, and PlayerChoice
	Genre                     string      `json:"genre"`                                               // The game genre display name
	GenreL1                   string      `json:"genre_l1"`                                            // The game genre from experience-genres-service
	GenreL2                   string      `json:"genre_l2"`                                            // The game subgenre from experience-genres-service
	IsAllGenre                bool        `json:"isAllGenre"`                                          // Is this game all genre
	IsFavoritedByUser         bool        `json:"isFavoritedByUser"`                                   // Is this game favorited by user
	FavoritedCount            uint64      `json:"favoritedCount"`                                      // Game number of favorites
}

// GameDetailsResponse represents a response containing multiple game details.
type GameDetailsResponse struct {
	Data []GameDetailResponse `json:"data" validate:"required,dive"` // List of game details
}

// ServerResponse represents the response structure for game server queries.
type ServerResponse struct {
	PreviousPageCursor *string  `json:"previousPageCursor" validate:"omitempty"`     // Cursor for the previous page of results (if any)
	NextPageCursor     *string  `json:"nextPageCursor"     validate:"omitempty"`     // Cursor for the next page of results (if any)
	Data               []Server `json:"data"               validate:"required,dive"` // List of servers
}

// Server represents a single game server instance.
type Server struct {
	ID           string   `json:"id"           validate:"required"`       // Unique identifier for the server
	MaxPlayers   int32    `json:"maxPlayers"   validate:"required,gt=0"`  // Maximum number of players allowed
	Playing      int32    `json:"playing"      validate:"required,gte=0"` // Current number of players
	PlayerTokens []string `json:"playerTokens" validate:"required"`       // List of player tokens
	Players      []string `json:"players"`                                // List of players (may be empty)
	FPS          float64  `json:"fps"          validate:"required"`       // Current server FPS
	Ping         int32    `json:"ping"         validate:"gte=0"`          // Server ping in milliseconds (may be omitted)
}

// PlaceDetailResponse represents detailed information about a place.
type PlaceDetailResponse struct {
	PlaceID             uint64 `json:"placeId"             validate:"required,min=1"` // The place ID
	Name                string `json:"name"                validate:"required"`       // The place name
	Description         string `json:"description"`                                   // The place description
	SourceName          string `json:"sourceName"`                                    // The place name in source language
	SourceDescription   string `json:"sourceDescription"`                             // The place description in source language
	URL                 string `json:"url"                 validate:"required"`       // URL to the place
	Builder             string `json:"builder"             validate:"required"`       // Builder's username
	BuilderID           uint64 `json:"builderId"           validate:"required,min=1"` // Builder's ID
	HasVerifiedBadge    bool   `json:"hasVerifiedBadge"`                              // Whether the builder has verified badge
	IsPlayable          bool   `json:"isPlayable"`                                    // Whether the place is playable
	ReasonProhibited    string `json:"reasonProhibited"`                              // Reason why the place is not playable (if applicable)
	UniverseID          uint64 `json:"universeId"          validate:"required,min=1"` // Associated universe ID
	UniverseRootPlaceID uint64 `json:"universeRootPlaceId" validate:"required,min=1"` // Root place ID of the universe
	Price               uint64 `json:"price"               validate:"min=0"`          // Price to access the place
	ImageToken          string `json:"imageToken"          validate:"required"`       // Token for the place thumbnail
}
