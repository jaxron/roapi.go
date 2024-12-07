package types

import "time"

// UserPresenceType represents the type of presence for a user.
type UserPresenceType int

const (
	Offline  UserPresenceType = 0
	Website  UserPresenceType = 1
	InGame   UserPresenceType = 2
	InStudio UserPresenceType = 3
)

// UserPresencesResponse represents the structure of user presences information returned by the Roblox API.
type UserPresencesResponse struct {
	UserPresences []UserPresenceResponse `json:"userPresences" validate:"required,dive"` // List of user presences
}

// UserPresenceResponse represents the presence information for a single user.
type UserPresenceResponse struct {
	UserPresenceType UserPresenceType `json:"userPresenceType" validate:"oneof=0 1 2 3"`   // Type of presence (Offline, Website, InGame, InStudio)
	LastLocation     string           `json:"lastLocation"     validate:"omitempty"`       // Last known location of the user
	PlaceID          *uint64          `json:"placeId"          validate:"omitempty,min=1"` // ID of the place if user is in game
	RootPlaceID      *uint64          `json:"rootPlaceId"      validate:"omitempty,min=1"` // ID of the root place if user is in game
	GameID           *string          `json:"gameId"           validate:"omitempty,min=1"` // ID of the game instance if user is in game
	UniverseID       *uint64          `json:"universeId"       validate:"omitempty,min=1"` // ID of the universe if user is in game
	UserID           uint64           `json:"userId"           validate:"required,min=1"`  // ID of the user
	LastOnline       time.Time        `json:"lastOnline"       validate:"required"`        // Last time the user was online
}
