package models

import "time"

// UserByIDResponse represents the structure of user information returned by the Roblox API.
type UserByIDResponse struct {
	ID                     uint64    `json:"id"`                     // Unique identifier for the user
	Name                   string    `json:"name"`                   // Username of the user
	DisplayName            string    `json:"displayName"`            // Display name of the user
	Description            string    `json:"description"`            // User's profile description
	Created                time.Time `json:"created"`                // Date when the user account was created
	IsBanned               bool      `json:"isBanned"`               // Whether the user is banned
	ExternalAppDisplayName *string   `json:"externalAppDisplayName"` // Display name for external apps (if any)
}

// AuthUserResponse represents the minimal structure of authenticated user information returned by the Roblox API.
type AuthUserResponse struct {
	ID          uint64 `json:"id"`          // Unique identifier for the user
	Name        string `json:"name"`        // Username of the user
	DisplayName string `json:"displayName"` // Display name of the user
}

// UserByUsernameResponse represents a single user fetched by username.
type UserByUsernameResponse struct {
	ID                uint64 `json:"id"`                // Unique identifier for the user
	Name              string `json:"name"`              // Current username of the user
	DisplayName       string `json:"displayName"`       // Display name of the user
	RequestedUsername string `json:"requestedUsername"` // The username that was requested in the API call
	HasVerifiedBadge  bool   `json:"hasVerifiedBadge"`  // Whether the user has a verified badge
}

// VerifiedBadgeUserResponse represents a single user with a verified badge.
type VerifiedBadgeUserResponse struct {
	ID               uint64 `json:"id"`               // Unique identifier for the user
	Name             string `json:"name"`             // Current username of the user
	DisplayName      string `json:"displayName"`      // Display name of the user
	HasVerifiedBadge bool   `json:"hasVerifiedBadge"` // Whether the user has a verified badge
}

// UsernameHistoryPageResponse represents the structure of a user's username history returned by the Roblox API.
type UsernameHistoryPageResponse struct {
	PreviousPageCursor *string                   `json:"previousPageCursor"` // Cursor for the previous page of results (if any)
	NextPageCursor     *string                   `json:"nextPageCursor"`     // Cursor for the next page of results (if any)
	Data               []UsernameHistoryResponse `json:"data"`               // List of previous usernames
}

// UsernameHistoryResponse represents a single username in a user's username history.
type UsernameHistoryResponse struct {
	Name string `json:"name"` // A previous username of the user
}

// UserSearchPageResponse represents the structure of a user search result returned by the Roblox API.
type UserSearchPageResponse struct {
	PreviousPageCursor *string              `json:"previousPageCursor"` // Cursor for the previous page of results (if any)
	NextPageCursor     *string              `json:"nextPageCursor"`     // Cursor for the next page of results (if any)
	Data               []UserSearchResponse `json:"data"`               // List of users matching the search criteria
}

// UserSearchResponse represents a single user in a username search result.
type UserSearchResponse struct {
	ID                uint64   `json:"id"`                // Unique identifier for the user
	Name              string   `json:"name"`              // Current username of the user
	DisplayName       string   `json:"displayName"`       // Display name of the user
	HasVerifiedBadge  bool     `json:"hasVerifiedBadge"`  // Whether the user has a verified badge
	PreviousUsernames []string `json:"previousUsernames"` // List of previous usernames for this user
}
