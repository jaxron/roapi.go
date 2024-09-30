package models

import "time"

// FollowerPageResponse represents the structure of a user's follower list returned by the Roblox API.
type FollowerPageResponse struct {
	PreviousPageCursor *string        `json:"previousPageCursor"` // Cursor for the previous page of results (if any)
	NextPageCursor     *string        `json:"nextPageCursor"`     // Cursor for the next page of results (if any)
	Data               []UserResponse `json:"data"`               // List of followers
}

// UserResponse represents a single user in a user's friend list.
type UserResponse struct {
	ID                     uint64    `json:"id"`                     // Unique identifier for the friend
	Name                   string    `json:"name"`                   // Username of the friend
	DisplayName            string    `json:"displayName"`            // Display name of the friend
	Description            *string   `json:"description"`            // Friend's profile description
	Created                time.Time `json:"created"`                // Date when the friend account was created
	IsBanned               bool      `json:"isBanned"`               // Whether the friend is banned
	ExternalAppDisplayName *string   `json:"externalAppDisplayName"` // Display name for external apps (if any)
	HasVerifiedBadge       bool      `json:"hasVerifiedBadge"`       // Whether the friend has a verified badge
	FriendFrequentScore    int       `json:"friendFrequentScore"`    // Friend frequent score
	FriendFrequentRank     int       `json:"friendFrequentRank"`     // Friend frequent rank
	PresenceType           int       `json:"presenceType,omitempty"` // Type of presence
	IsOnline               bool      `json:"isOnline"`               // Whether the friend is online
	IsDeleted              bool      `json:"isDeleted"`              // Whether the friend account is deleted
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
