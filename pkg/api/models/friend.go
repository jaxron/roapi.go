package models

// FriendInfos represents the structure of a list of friend information returned by the Roblox API.
type FriendInfos struct {
	Data []FriendInfo `json:"data"` // List of friend information
}

// FriendInfo represents a single friend in a list of friend information.
type FriendInfo struct {
	ID                     uint64  `json:"id"`                     // Unique identifier for the friend
	Name                   string  `json:"name"`                   // Username of the friend
	DisplayName            string  `json:"displayName"`            // Display name of the friend
	Description            string  `json:"description"`            // Friend's profile description
	Created                string  `json:"created"`                // Date when the friend account was created
	IsBanned               bool    `json:"isBanned"`               // Whether the friend is banned
	ExternalAppDisplayName *string `json:"externalAppDisplayName"` // Display name for external apps (if any)
	HasVerifiedBadge       bool    `json:"hasVerifiedBadge"`       // Whether the friend has a verified badge
	FriendFrequentScore    int     `json:"friendFrequentScore"`    // Friend frequent score
	FriendFrequentRank     int     `json:"friendFrequentRank"`     // Friend frequent rank
	IsOnline               bool    `json:"isOnline"`               // Whether the friend is online
	IsDeleted              bool    `json:"isDeleted"`              // Whether the friend account is deleted
}
