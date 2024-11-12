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

// UserPresence represents the presence information for a single user.
type UserPresence struct {
	UserPresenceType UserPresenceType `json:"userPresenceType"`
	LastLocation     string           `json:"lastLocation"`
	PlaceID          *uint64          `json:"placeId"`
	RootPlaceID      *uint64          `json:"rootPlaceId"`
	GameID           *string          `json:"gameId"`
	UniverseID       *uint64          `json:"universeId"`
	UserID           uint64           `json:"userId"`
	LastOnline       time.Time        `json:"lastOnline"`
}
