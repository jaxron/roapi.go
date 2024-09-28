package models

// UserInfo represents the structure of user information returned by the Roblox API.
type UserInfo struct {
	ID                     uint64  `json:"id"`                     // Unique identifier for the user
	Name                   string  `json:"name"`                   // Username of the user
	DisplayName            string  `json:"displayName"`            // Display name of the user
	Description            string  `json:"description"`            // User's profile description
	Created                string  `json:"created"`                // Date when the user account was created
	IsBanned               bool    `json:"isBanned"`               // Whether the user is banned
	ExternalAppDisplayName *string `json:"externalAppDisplayName"` // Display name for external apps (if any)
}

// AuthUserInfo represents the structure of authenticated user information returned by the Roblox API.
type AuthUserInfo struct {
	ID          uint64 `json:"id"`          // Unique identifier for the authenticated user
	Name        string `json:"name"`        // Username of the authenticated user
	DisplayName string `json:"displayName"` // Display name of the authenticated user
}

// UsersByUsernames represents the structure of a list of users returned by the Roblox API.
type UsersByUsernames struct {
	Data []UserByUsername `json:"data"` // List of users fetched by usernames
}

// UserByUsername represents a single user in a list of users fetched by username.
type UserByUsername struct {
	ID                uint64 `json:"id"`                // Unique identifier for the user
	Name              string `json:"name"`              // Current username of the user
	DisplayName       string `json:"displayName"`       // Display name of the user
	RequestedUsername string `json:"requestedUsername"` // The username that was requested in the API call
	HasVerifiedBadge  bool   `json:"hasVerifiedBadge"`  // Whether the user has a verified badge
}

// UsersByIDs represents the structure of a list of users returned by the Roblox API.
type UsersByIDs struct {
	Data []UserByID `json:"data"` // List of users fetched by user IDs
}

// UserByID represents a single user in a list of users fetched by user ID.
type UserByID struct {
	ID               uint64 `json:"id"`               // Unique identifier for the user
	Name             string `json:"name"`             // Current username of the user
	DisplayName      string `json:"displayName"`      // Display name of the user
	HasVerifiedBadge bool   `json:"hasVerifiedBadge"` // Whether the user has a verified badge
}

// UsernameHistory represents the structure of a user's username history returned by the Roblox API.
type UsernameHistory struct {
	PreviousPageCursor *string    `json:"previousPageCursor"` // Cursor for the previous page of results (if any)
	NextPageCursor     *string    `json:"nextPageCursor"`     // Cursor for the next page of results (if any)
	Data               []Username `json:"data"`               // List of previous usernames
}

// Username represents a single username in a user's username history.
type Username struct {
	Name string `json:"name"` // A previous username of the user
}

// SearchResult represents the structure of a user search result returned by the Roblox API.
type SearchResult struct {
	PreviousPageCursor *string        `json:"previousPageCursor"` // Cursor for the previous page of results (if any)
	NextPageCursor     *string        `json:"nextPageCursor"`     // Cursor for the next page of results (if any)
	Data               []SearchedUser `json:"data"`               // List of users matching the search criteria
}

// SearchedUser represents a single user in a username search result.
type SearchedUser struct {
	ID                uint64   `json:"id"`                // Unique identifier for the user
	Name              string   `json:"name"`              // Current username of the user
	DisplayName       string   `json:"displayName"`       // Display name of the user
	HasVerifiedBadge  bool     `json:"hasVerifiedBadge"`  // Whether the user has a verified badge
	PreviousUsernames []string `json:"previousUsernames"` // List of previous usernames for this user
}
