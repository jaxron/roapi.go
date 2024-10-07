package users

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/models"
)

// GetUsernameHistory fetches the username history for a user.
// GET https://users.roblox.com/v1/users/{userID}/username-history
func (s *Service) GetUsernameHistory(ctx context.Context, params UsernameHistoryParams) (*models.UsernameHistoryPageResponse, error) {
	if err := s.validate.Struct(params); err != nil {
		return nil, err
	}

	var history models.UsernameHistoryPageResponse
	resp, err := s.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/username-history", UsersEndpoint, params.UserID)).
		Query("limit", strconv.FormatUint(params.Limit, 10)).
		Query("sortOrder", params.SortOrder).
		Query("cursor", params.Cursor).
		Result(&history).
		JSONHeaders().
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return &history, nil
}

// UsernameHistoryParams holds the parameters for fetching username history.
type UsernameHistoryParams struct {
	UserID    uint64 `json:"userId"    validate:"required"`
	Limit     uint64 `json:"limit"     validate:"oneof=10 25 50 100"`
	SortOrder string `json:"sortOrder" validate:"oneof=Asc Desc"`
	Cursor    string `json:"cursor"    validate:"omitempty,base64"`
}

// UsernameHistoryBuilder builds parameters for GetUsernameHistory API call.
type UsernameHistoryBuilder struct {
	params UsernameHistoryParams
}

// NewUsernameHistoryBuilder creates a new UsernameHistoryBuilder with the given user ID.
func NewUsernameHistoryBuilder(userID uint64) *UsernameHistoryBuilder {
	return &UsernameHistoryBuilder{
		params: UsernameHistoryParams{
			UserID:    userID,
			Limit:     10,
			SortOrder: SortOrderAsc,
			Cursor:    "",
		},
	}
}

// WithLimit sets the maximum number of results to return.
func (b *UsernameHistoryBuilder) WithLimit(limit uint64) *UsernameHistoryBuilder {
	b.params.Limit = limit
	return b
}

// WithCursor sets the cursor for pagination.
func (b *UsernameHistoryBuilder) WithCursor(cursor string) *UsernameHistoryBuilder {
	b.params.Cursor = cursor
	return b
}

// WithSortOrderAsc sets the sort order for results to ascending.
func (b *UsernameHistoryBuilder) WithSortOrderAsc() *UsernameHistoryBuilder {
	b.params.SortOrder = SortOrderAsc
	return b
}

// WithSortOrderDesc sets the sort order for results to descending.
func (b *UsernameHistoryBuilder) WithSortOrderDesc() *UsernameHistoryBuilder {
	b.params.SortOrder = SortOrderDesc
	return b
}

// Build returns the UsernameHistoryParams.
func (b *UsernameHistoryBuilder) Build() UsernameHistoryParams {
	return b.params
}
