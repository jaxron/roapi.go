package types

// Constants for Roblox API endpoints.
const (
	UsersEndpoint      = "https://users.roblox.com"
	FriendsEndpoint    = "https://friends.roblox.com"
	GroupsEndpoint     = "https://groups.roblox.com"
	ThumbnailsEndpoint = "https://thumbnails.roblox.com"
	AvatarEndpoint     = "https://avatar.roblox.com"
)

// SortOrder represents the sort order of the results.
type SortOrder string

const (
	SortOrderAsc  SortOrder = "Asc"
	SortOrderDesc SortOrder = "Desc"
)
