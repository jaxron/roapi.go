package types

// Constants for Roblox API endpoints.
const (
	UsersEndpoint      = "https://users.roblox.com"
	FriendsEndpoint    = "https://friends.roblox.com"
	GroupsEndpoint     = "https://groups.roblox.com"
	ThumbnailsEndpoint = "https://thumbnails.roblox.com"
)

// Constants for building requests.
const (
	SortOrderAsc  = "Asc"
	SortOrderDesc = "Desc"
)

// Constants for user presence types.
const (
	UserPresenceTypeOnline string = "Online"
	UserPresenceTypeInGame string = "InGame"

	UserLocationTypePage string = "Page"
	UserLocationTypeGame string = "Game"
)
