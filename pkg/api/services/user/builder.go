package user

import "github.com/bytedance/sonic"

const (
	SortOrderAsc  = "Asc"
	SortOrderDesc = "Desc"

	DefaultLimit = 10
)

// UsersByUsernamesBuilder builds parameters for GetUsersByUsernames API call.
type UsersByUsernamesBuilder struct {
	usernames          []string // Required: List of usernames to fetch information for
	excludeBannedUsers bool     // Optional: Whether to exclude banned users from the result
}

// NewUsersByUsernamesBuilder creates a new UsersByUsernamesBuilder with the given usernames.
func NewUsersByUsernamesBuilder(usernames []string) *UsersByUsernamesBuilder {
	return &UsersByUsernamesBuilder{
		usernames:          usernames,
		excludeBannedUsers: false, // Default: include banned users
	}
}

// ExcludeBannedUsers sets whether to exclude banned users from the result.
func (b *UsersByUsernamesBuilder) ExcludeBannedUsers(excludeBannedUsers bool) *UsersByUsernamesBuilder {
	b.excludeBannedUsers = excludeBannedUsers
	return b
}

// MarshalJSON converts the UsersByUsernamesBuilder to JSON for API requests.
func (b *UsersByUsernamesBuilder) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(struct {
		Usernames          []string `json:"usernames"`
		ExcludeBannedUsers bool     `json:"excludeBannedUsers"`
	}{
		Usernames:          b.usernames,
		ExcludeBannedUsers: b.excludeBannedUsers,
	})
}

// UsersByIDsBuilder builds parameters for GetUsersByIDs API call.
type UsersByIDsBuilder struct {
	userIds            []uint64 // Required: List of user IDs to fetch information for
	excludeBannedUsers bool     // Optional: Whether to exclude banned users from the result
}

// NewUsersByIDsBuilder creates a new UsersByIDsBuilder with the given user IDs.
func NewUsersByIDsBuilder(userIds []uint64) *UsersByIDsBuilder {
	return &UsersByIDsBuilder{
		userIds:            userIds,
		excludeBannedUsers: false, // Default: include banned users
	}
}

// ExcludeBannedUsers sets whether to exclude banned users from the result.
func (b *UsersByIDsBuilder) ExcludeBannedUsers(excludeBannedUsers bool) *UsersByIDsBuilder {
	b.excludeBannedUsers = excludeBannedUsers
	return b
}

// MarshalJSON converts the UsersByIDsBuilder to JSON for API requests.
func (b *UsersByIDsBuilder) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(struct {
		UserIds            []uint64 `json:"userIds"`
		ExcludeBannedUsers bool     `json:"excludeBannedUsers"`
	}{
		UserIds:            b.userIds,
		ExcludeBannedUsers: b.excludeBannedUsers,
	})
}

// UsernameHistoryBuilder builds parameters for GetUsernameHistory API call.
type UsernameHistoryBuilder struct {
	userID    uint64 // Required: ID of the user to fetch username history for
	limit     uint64 // Optional: Maximum number of results to return (default: 10)
	sortOrder string // Optional: Sort order for results (default: Ascending)
	cursor    string // Optional: Cursor for pagination
}

// NewUsernameHistoryBuilder creates a new UsernameHistoryBuilder with the given user ID.
func NewUsernameHistoryBuilder(userID uint64) *UsernameHistoryBuilder {
	return &UsernameHistoryBuilder{
		userID:    userID,
		limit:     DefaultLimit,
		sortOrder: SortOrderAsc,
		cursor:    "",
	}
}

// Limit sets the maximum number of results to return.
func (b *UsernameHistoryBuilder) Limit(limit uint64) *UsernameHistoryBuilder {
	b.limit = limit
	return b
}

// Cursor sets the cursor for pagination.
func (b *UsernameHistoryBuilder) Cursor(cursor string) *UsernameHistoryBuilder {
	b.cursor = cursor
	return b
}

// SortOrder sets the sort order for results.
func (b *UsernameHistoryBuilder) SortOrder(sortOrder string) *UsernameHistoryBuilder {
	b.sortOrder = sortOrder
	return b
}

// SearchUserBuilder builds parameters for SearchUser API call.
type SearchUserBuilder struct {
	username string // Required: Username to search for
	limit    uint64 // Optional: Maximum number of results to return (default: 10)
	cursor   string // Optional: Cursor for pagination
}

// NewSearchUserBuilder creates a new SearchUserBuilder with the given username.
func NewSearchUserBuilder(username string) *SearchUserBuilder {
	return &SearchUserBuilder{
		username: username,
		limit:    DefaultLimit,
		cursor:   "",
	}
}

// Limit sets the maximum number of results to return.
func (b *SearchUserBuilder) Limit(limit uint64) *SearchUserBuilder {
	b.limit = limit
	return b
}

// Cursor sets the cursor for pagination.
func (b *SearchUserBuilder) Cursor(cursor string) *SearchUserBuilder {
	b.cursor = cursor
	return b
}
