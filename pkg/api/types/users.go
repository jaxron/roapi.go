package types

import "time"

// UserByIDResponse represents the structure of user information returned by the Roblox API.
type UserByIDResponse struct {
	ID                     int64     `json:"id"                     validate:"required,min=1"`  // Unique identifier for the user
	Name                   string    `json:"name"                   validate:"required,min=1"`  // Username of the user
	DisplayName            string    `json:"displayName"            validate:"required,min=1"`  // Display name of the user
	Description            string    `json:"description"`                                       // User's profile description
	Created                time.Time `json:"created"                validate:"required"`        // Date when the user account was created
	IsBanned               bool      `json:"isBanned"`                                          // Whether the user is banned
	ExternalAppDisplayName *string   `json:"externalAppDisplayName" validate:"omitempty,min=1"` // Display name for external apps (if any)
}

// AuthUserResponse represents the minimal structure of authenticated user information returned by the Roblox API.
type AuthUserResponse struct {
	ID          int64  `json:"id"          validate:"required,min=1"` // Unique identifier for the user
	Name        string `json:"name"        validate:"required,min=1"` // Username of the user
	DisplayName string `json:"displayName" validate:"required,min=1"` // Display name of the user
}

// UsersByUsernameResponse represents the structure of users fetched by username.
type UsersByUsernameResponse struct {
	Data []UserByUsername `json:"data" validate:"required,dive"` // List of users fetched by username
}

// UserByUsername represents a single user fetched by username.
type UserByUsername struct {
	ID                int64  `json:"id"                validate:"required,min=1"` // Unique identifier for the user
	Name              string `json:"name"              validate:"required,min=1"` // Current username of the user
	DisplayName       string `json:"displayName"       validate:"required,min=1"` // Display name of the user
	RequestedUsername string `json:"requestedUsername" validate:"required,min=1"` // The username that was requested in the API call
	HasVerifiedBadge  bool   `json:"hasVerifiedBadge"`                            // Whether the user has a verified badge
}

// UsersByIDsResponse represents the structure of users fetched by ID.
type UsersByIDsResponse struct {
	Data []VerifiedBadgeUser `json:"data" validate:"required,dive"` // List of users fetched by ID
}

// VerifiedBadgeUser represents a single user with a verified badge.
type VerifiedBadgeUser struct {
	ID               int64  `json:"id"               validate:"required,min=1"` // Unique identifier for the user
	Name             string `json:"name"             validate:"required,min=1"` // Current username of the user
	DisplayName      string `json:"displayName"      validate:"required,min=1"` // Display name of the user
	HasVerifiedBadge bool   `json:"hasVerifiedBadge"`                           // Whether the user has a verified badge
}

// UsernameHistoryPageResponse represents the structure of a user's username history returned by the Roblox API.
type UsernameHistoryPageResponse struct {
	PreviousPageCursor *string                   `json:"previousPageCursor" validate:"omitempty,base64"` // Cursor for the previous page of results (if any)
	NextPageCursor     *string                   `json:"nextPageCursor"     validate:"omitempty,base64"` // Cursor for the next page of results (if any)
	Data               []UsernameHistoryResponse `json:"data"               validate:"required,dive"`    // List of previous usernames
}

// UsernameHistoryResponse represents a single username in a user's username history.
type UsernameHistoryResponse struct {
	Name string `json:"name" validate:"required,min=1"` // A previous username of the user
}

// UserSearchPageResponse represents the structure of a user search result returned by the Roblox API.
type UserSearchPageResponse struct {
	PreviousPageCursor *string              `json:"previousPageCursor" validate:"omitempty,base64"` // Cursor for the previous page of results (if any)
	NextPageCursor     *string              `json:"nextPageCursor"     validate:"omitempty,base64"` // Cursor for the next page of results (if any)
	Data               []UserSearchResponse `json:"data"               validate:"required,dive"`    // List of users matching the search criteria
}

// UserSearchResponse represents a single user in a username search result.
type UserSearchResponse struct {
	ID                int64    `json:"id"                validate:"required,min=1"`      // Unique identifier for the user
	Name              string   `json:"name"              validate:"required,min=1"`      // Current username of the user
	DisplayName       string   `json:"displayName"       validate:"required,min=1"`      // Display name of the user
	HasVerifiedBadge  bool     `json:"hasVerifiedBadge"`                                 // Whether the user has a verified badge
	PreviousUsernames []string `json:"previousUsernames" validate:"required,dive,min=1"` // List of previous usernames for this user
}
