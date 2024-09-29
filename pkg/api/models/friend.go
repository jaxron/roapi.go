package models

import "time"

// FriendResponse represents a single friend in a user's friend list.
type FriendResponse struct {
	ID                     uint64    `json:"id"`                     // Unique identifier for the friend
	Name                   string    `json:"name"`                   // Username of the friend
	DisplayName            string    `json:"displayName"`            // Display name of the friend
	Description            string    `json:"description"`            // Friend's profile description
	Created                time.Time `json:"created"`                // Date when the friend account was created
	IsBanned               bool      `json:"isBanned"`               // Whether the friend is banned
	ExternalAppDisplayName *string   `json:"externalAppDisplayName"` // Display name for external apps (if any)
	HasVerifiedBadge       bool      `json:"hasVerifiedBadge"`       // Whether the friend has a verified badge
	FriendFrequentScore    int       `json:"friendFrequentScore"`    // Friend frequent score
	FriendFrequentRank     int       `json:"friendFrequentRank"`     // Friend frequent rank
	IsOnline               bool      `json:"isOnline"`               // Whether the friend is online
	IsDeleted              bool      `json:"isDeleted"`              // Whether the friend account is deleted
}
