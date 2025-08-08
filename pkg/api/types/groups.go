package types

import "time"

// GroupResponse represents the structure of group information returned by the Roblox API.
type GroupResponse struct {
	ID                 int64       `json:"id"                 validate:"required,min=1"` // Unique identifier for the group
	Name               string      `json:"name"               validate:"required,min=1"` // Name of the group
	Description        string      `json:"description"`                                  // Description of the group
	Owner              *GroupUser  `json:"owner"              validate:"omitempty"`      // Owner information
	Shout              *GroupShout `json:"shout"              validate:"omitempty"`      // Group shout (if any)
	MemberCount        int64       `json:"memberCount"        validate:"min=0"`          // Number of members in the group
	IsBuildersClubOnly bool        `json:"isBuildersClubOnly"`                           // Whether the group is builders club only
	PublicEntryAllowed bool        `json:"publicEntryAllowed"`                           // Whether public entry is allowed
	IsLocked           *bool       `json:"isLocked"`                                     // Whether the group is locked
	HasVerifiedBadge   bool        `json:"hasVerifiedBadge"`                             // Whether the group has a verified badge
}

// GroupUser represents a user in the context of a group.
type GroupUser struct {
	UserID           int64  `json:"userId"           validate:"required,min=1"` // User's unique identifier
	Username         string `json:"username"         validate:"required,min=1"` // Username of the user
	DisplayName      string `json:"displayName"      validate:"required,min=1"` // Display name of the user
	HasVerifiedBadge bool   `json:"hasVerifiedBadge"`                           // Whether the user has a verified badge
}

// GroupShout represents a group shout.
type GroupShout struct {
	Body    string    `json:"body"`                                // Content of the shout
	Poster  GroupUser `json:"poster"  validate:"required"`         // User who posted the shout
	Created time.Time `json:"created" validate:"required"`         // When the shout was created
	Updated time.Time `json:"updated" validate:"gtefield=Created"` // When the shout was last updated
}

// GroupUsersResponse represents the structure of group users information returned by the Roblox API.
type GroupUsersResponse struct {
	PreviousPageCursor *string         `json:"previousPageCursor" validate:"omitempty"`     // Cursor for the previous page of results (if any)
	NextPageCursor     *string         `json:"nextPageCursor"     validate:"omitempty"`     // Cursor for the next page of results (if any)
	Data               []GroupUserData `json:"data"               validate:"required,dive"` // List of group users
}

// GroupUserData represents a single user in a group's user list.
type GroupUserData struct {
	User GroupUser `json:"user" validate:"required"` // User information
	Role GroupRole `json:"role" validate:"required"` // User's role in the group
}

// GroupRole represents a role in a group.
type GroupRole struct {
	ID          int64  `json:"id"          validate:"required,min=1"` // Role's unique identifier
	Name        string `json:"name"        validate:"required,min=1"` // Name of the role
	Rank        int64  `json:"rank"`                                  // Rank of the role (0 is lowest)
	MemberCount int64  `json:"memberCount"`                           // Number of members with this role
}

// GroupRolesResponse represents the structure of group roles information returned by the Roblox API.
type GroupRolesResponse struct {
	GroupID int64       `json:"groupId" validate:"required,min=1"` // Unique identifier for the group
	Roles   []GroupRole `json:"roles"   validate:"required,dive"`  // List of roles in the group
}

// RoleUsersResponse represents the structure of role users information returned by the Roblox API.
type RoleUsersResponse struct {
	PreviousPageCursor *string     `json:"previousPageCursor" validate:"omitempty"`     // Cursor for the previous page of results (if any)
	NextPageCursor     *string     `json:"nextPageCursor"     validate:"omitempty"`     // Cursor for the next page of results (if any)
	Data               []GroupUser `json:"data"               validate:"required,dive"` // List of users with the role
}

// SearchGroupsResponse represents the structure of group search results returned by the Roblox API.
type SearchGroupsResponse struct {
	Keyword            string        `json:"keyword"            validate:"required"`      // Search keyword used
	PreviousPageCursor *string       `json:"previousPageCursor" validate:"omitempty"`     // Cursor for the previous page of results (if any)
	NextPageCursor     *string       `json:"nextPageCursor"     validate:"omitempty"`     // Cursor for the next page of results (if any)
	Data               []GroupSearch `json:"data"               validate:"required,dive"` // List of groups matching the search criteria
}

// GroupSearch represents a single group in the search results.
type GroupSearch struct {
	ID                 int64     `json:"id"                 validate:"required,min=1"`   // Unique identifier for the group
	Name               string    `json:"name"               validate:"required,min=1"`   // Name of the group
	Description        string    `json:"description"`                                    // Description of the group
	MemberCount        int64     `json:"memberCount"        validate:"min=0"`            // Number of members in the group
	PreviousName       string    `json:"previousName"`                                   // Previous name of the group (if any)
	PublicEntryAllowed bool      `json:"publicEntryAllowed"`                             // Whether public entry is allowed
	Created            time.Time `json:"created"            validate:"required"`         // When the group was created
	Updated            time.Time `json:"updated"            validate:"gtefield=Created"` // When the group was last updated
	HasVerifiedBadge   bool      `json:"hasVerifiedBadge"`                               // Whether the group has a verified badge
}

// GroupLookupResponse represents the structure of group lookup results returned by the Roblox API.
type GroupLookupResponse struct {
	Data []GroupLookup `json:"data" validate:"required,dive"` // List of groups matching the lookup criteria
}

// GroupLookup represents the structure of group lookup results returned by the Roblox API.
type GroupLookup struct {
	ID               int64  `json:"id"               validate:"required,min=1"` // Unique identifier for the group
	Name             string `json:"name"             validate:"required,min=1"` // Name of the group
	MemberCount      int64  `json:"memberCount"      validate:"min=0"`          // Number of members in the group
	HasVerifiedBadge bool   `json:"hasVerifiedBadge"`                           // Whether the group has a verified badge
}

// GroupsInfoResponse represents the structure of multiple groups information returned by the Roblox API.
type GroupsInfoResponse struct {
	Data []GroupInfo `json:"data" validate:"required,dive"` // List of group information
}

// GroupInfo represents the structure of a single group's information returned by the Roblox API.
type GroupInfo struct {
	ID               int64      `json:"id"               validate:"required,min=1"` // Unique identifier for the group
	Name             string     `json:"name"             validate:"required,min=1"` // Name of the group
	Description      string     `json:"description"`                                // Description of the group
	Owner            GroupOwner `json:"owner"`                                      // Owner information
	Created          time.Time  `json:"created"          validate:"required"`       // When the group was created
	HasVerifiedBadge bool       `json:"hasVerifiedBadge"`                           // Whether the group has a verified badge
}

// GroupOwner represents the owner of a group.
type GroupOwner struct {
	ID   int64  `json:"id"   validate:"required,min=1"` // Owner's unique identifier
	Type string `json:"type" validate:"required"`       // Type of owner
}

// UserGroupRolesResponse represents the structure of user group roles information returned by the Roblox API.
type UserGroupRolesResponse struct {
	Data []UserGroupRoles `json:"data" validate:"required,dive"` // List of user group roles
}

// UserGroupRoles represents the structure of user group roles information returned by the Roblox API.
type UserGroupRoles struct {
	Group GroupResponse `json:"group" validate:"required"` // Group information
	Role  UserGroupRole `json:"role"  validate:"required"` // User's role in the group
}

// UserGroupRole represents a single group role for a user.
type UserGroupRole struct {
	ID   int64  `json:"id"   validate:"required,min=1"` // Role's unique identifier
	Name string `json:"name" validate:"required,min=1"` // Name of the role
	Rank int64  `json:"rank"`                           // Rank of the role (0 is lowest)
}

// GroupWallPostsResponse represents the structure of group wall posts returned by the Roblox API.
type GroupWallPostsResponse struct {
	PreviousPageCursor *string         `json:"previousPageCursor" validate:"omitempty"`     // Cursor for the previous page of results (if any)
	NextPageCursor     *string         `json:"nextPageCursor"     validate:"omitempty"`     // Cursor for the next page of results (if any)
	Data               []GroupWallPost `json:"data"               validate:"required,dive"` // List of wall posts
}

// GroupWallPost represents a single wall post in a group.
type GroupWallPost struct {
	ID      int64            `json:"id"      validate:"required,min=1"` // Unique identifier for the wall post
	Poster  *GroupWallPoster `json:"poster"  validate:"omitempty"`      // User who posted the wall post
	Body    string           `json:"body"`                              // Content of the wall post
	Created time.Time        `json:"created" validate:"required"`       // When the wall post was created
	Updated time.Time        `json:"updated" validate:"required"`       // When the wall post was last updated
}

// GroupWallPoster represents the user who posted a wall post.
type GroupWallPoster struct {
	User GroupUser `json:"user" validate:"required"` // User information
	Role GroupRole `json:"role" validate:"required"` // User's role in the group
}
