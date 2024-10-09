package types

import "time"

// GroupResponse represents the structure of group information returned by the Roblox API.
type GroupResponse struct {
	ID                 uint64     `json:"id"`
	Name               string     `json:"name"`
	Description        string     `json:"description"`
	Owner              GroupUser  `json:"owner"`
	Shout              GroupShout `json:"shout"`
	MemberCount        uint64     `json:"memberCount"`
	IsBuildersClubOnly bool       `json:"isBuildersClubOnly"`
	PublicEntryAllowed bool       `json:"publicEntryAllowed"`
	HasVerifiedBadge   bool       `json:"hasVerifiedBadge"`
}

// GroupUser represents a user in the context of a group.
type GroupUser struct {
	UserID           uint64 `json:"userId"`
	Username         string `json:"username"`
	DisplayName      string `json:"displayName"`
	HasVerifiedBadge bool   `json:"hasVerifiedBadge"`
}

// GroupShout represents a group shout.
type GroupShout struct {
	Body    string    `json:"body"`
	Poster  GroupUser `json:"poster"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// GroupUsersResponse represents the structure of group users information returned by the Roblox API.
type GroupUsersResponse struct {
	PreviousPageCursor *string         `json:"previousPageCursor"`
	NextPageCursor     *string         `json:"nextPageCursor"`
	Data               []GroupUserData `json:"data"`
}

// GroupUserData represents a single user in a group's user list.
type GroupUserData struct {
	User GroupUser `json:"user"`
	Role GroupRole `json:"role"`
}

// GroupRole represents a role in a group.
type GroupRole struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Rank        uint64 `json:"rank"`
	MemberCount uint64 `json:"memberCount"`
}

// GroupRolesResponse represents the structure of group roles information returned by the Roblox API.
type GroupRolesResponse struct {
	GroupID uint64      `json:"groupId"`
	Roles   []GroupRole `json:"roles"`
}

// RoleUsersResponse represents the structure of role users information returned by the Roblox API.
type RoleUsersResponse struct {
	PreviousPageCursor *string     `json:"previousPageCursor"`
	NextPageCursor     *string     `json:"nextPageCursor"`
	Data               []GroupUser `json:"data"`
}
