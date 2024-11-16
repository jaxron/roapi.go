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
	PreviousPageCursor *string  `json:"previousPageCursor"` // Cursor for the previous page of results (if any)
	NextPageCursor     *string  `json:"nextPageCursor"`     // Cursor for the next page of results (if any)
	Data               []Friend `json:"data"`               // List of followers
}

// FollowingPageResponse represents the structure of a user's following list returned by the Roblox API.
type FollowingPageResponse struct {
	PreviousPageCursor *string  `json:"previousPageCursor"` // Cursor for the previous page of results (if any)
	NextPageCursor     *string  `json:"nextPageCursor"`     // Cursor for the next page of results (if any)
	Data               []Friend `json:"data"`               // List of users being followed
}

// Friend represents a single friend in a user's friend list.
type Friend struct {
	ID uint64 `json:"id"` // Unique identifier for the friend
}

// FriendPageResponse represents the structure of a user's friend list returned by the Roblox API.
type FriendPageResponse struct {
	PreviousCursor *string          `json:"previousCursor"` // Cursor for the previous page of results (if any)
	NextCursor     *string          `json:"nextCursor"`     // Cursor for the next page of results (if any)
	PageItems      []FriendResponse `json:"pageItems"`      // List of friends
	HasMore        bool             `json:"hasMore"`        // Whether there are more friends to fetch
}

// FriendResponse represents the structure of friend information returned by the Roblox API.
type FriendResponse struct {
	ID               uint64 `json:"id"`               // Unique identifier for the friend
	HasVerifiedBadge bool   `json:"hasVerifiedBadge"` // Whether the friend has a verified badge
}

// OnlineFriend represents the structure of friend information returned by the Roblox API.
type OnlineFriend struct {
	ID           uint64               `json:"id"`           // Unique identifier for the friend
	UserPresence UserPresenceResponse `json:"userPresence"` // User presence information
}
