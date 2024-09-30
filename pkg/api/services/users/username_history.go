package users

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/models"
	"github.com/jaxron/roapi.go/pkg/client"
)

// GetUsernameHistory fetches the username history for a user.
// GET https://users.roblox.com/v1/users/{userID}/username-history
func (s *Service) GetUsernameHistory(ctx context.Context, b *UsernameHistoryBuilder) (*models.UsernameHistoryPageResponse, error) {
	var history models.UsernameHistoryPageResponse
	req := client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/username-history", UsersEndpoint, b.userID)).
		Query("limit", strconv.FormatUint(b.limit, 10)).
		Query("sortOrder", b.sortOrder).
		Query("cursor", b.cursor).
		Result(&history)

	resp, err := s.client.Do(ctx, req.Build())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &history, nil
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

// SortOrderAsc sets the sort order for results to ascending.
func (b *UsernameHistoryBuilder) SortOrderAsc(sortOrder string) *UsernameHistoryBuilder {
	b.sortOrder = SortOrderAsc
	return b
}

// SortOrderDesc sets the sort order for results to descending.
func (b *UsernameHistoryBuilder) SortOrderDesc(sortOrder string) *UsernameHistoryBuilder {
	b.sortOrder = SortOrderDesc
	return b
}
