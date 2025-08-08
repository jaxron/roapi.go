package types

// PresenceType represents the type of user presence.
type PresenceType string

const (
	UserPresenceTypeOnline PresenceType = "Online"
	UserPresenceTypeInGame PresenceType = "InGame"
)

// LocationType represents the type of user location.
type LocationType string

const (
	UserLocationTypePage LocationType = "Page"
	UserLocationTypeGame LocationType = "Game"
)

// FollowerPageResponse represents the structure of a user's follower list returned by the Roblox API.
type FollowerPageResponse struct {
	PreviousPageCursor *string  `json:"previousPageCursor" validate:"omitempty,base64"` // Cursor for the previous page of results (if any)
	NextPageCursor     *string  `json:"nextPageCursor"     validate:"omitempty,base64"` // Cursor for the next page of results (if any)
	Data               []Friend `json:"data"               validate:"required,dive"`    // List of followers
}

// FollowingPageResponse represents the structure of a user's following list returned by the Roblox API.
type FollowingPageResponse struct {
	PreviousPageCursor *string  `json:"previousPageCursor" validate:"omitempty,base64"` // Cursor for the previous page of results (if any)
	NextPageCursor     *string  `json:"nextPageCursor"     validate:"omitempty,base64"` // Cursor for the next page of results (if any)
	Data               []Friend `json:"data"               validate:"required,dive"`    // List of users being followed
}

// FriendsResponse represents the structure of a user's friends list returned by the Roblox API.
type FriendsResponse struct {
	Data []ExtendedFriend `json:"data" validate:"required,dive"` // List of friends
}

// Friend represents a single friend in a user's friend list.
type Friend struct {
	ID int64 `json:"id" validate:"required"` // Unique identifier for the friend
}

// ExtendedFriend represents a friend with additional information.
type ExtendedFriend struct {
	Friend

	Name        string `json:"name"        validate:"omitempty,min=1"` // Current username of the user
	DisplayName string `json:"displayName" validate:"omitempty,min=1"` // Display name of the user
}

// FriendPageResponse represents the structure of a user's friend list returned by the Roblox API.
type FriendPageResponse struct {
	PreviousCursor *string          `json:"previousCursor" validate:"omitempty"`     // Cursor for the previous page of results (if any)
	NextCursor     *string          `json:"nextCursor"     validate:"omitempty"`     // Cursor for the next page of results (if any)
	PageItems      []FriendResponse `json:"pageItems"      validate:"required,dive"` // List of friends
	HasMore        bool             `json:"hasMore"`                                 // Whether there are more friends to fetch
}

// FriendResponse represents the structure of friend information returned by the Roblox API.
type FriendResponse struct {
	ID               int64 `json:"id"               validate:"required"` // Unique identifier for the friend
	HasVerifiedBadge bool  `json:"hasVerifiedBadge"`                     // Whether the friend has a verified badge
}

// OnlineFriend represents the structure of friend information returned by the Roblox API.
type OnlineFriend struct {
	ID           int64                `json:"id"           validate:"required"` // Unique identifier for the friend
	UserPresence UserPresenceResponse `json:"userPresence" validate:"required"` // User presence information
}
