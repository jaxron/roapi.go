package users

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetUsernameHistory fetches the username history for a user.
// GET https://users.roblox.com/v1/users/{userID}/username-history
func (r *Resource) GetUsernameHistory(ctx context.Context, p UsernameHistoryParams) (*types.UsernameHistoryPageResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var history types.UsernameHistoryPageResponse

	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/username-history", types.UsersEndpoint, p.UserID)).
		Query("limit", strconv.FormatInt(p.Limit, 10)).
		Query("sortOrder", string(p.SortOrder)).
		Query("cursor", p.Cursor).
		Result(&history).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}

	defer func() { _ = resp.Body.Close() }()

	if err := r.validate.Struct(&history); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &history, nil
}

// UsernameHistoryParams holds the parameters for fetching username history.
type UsernameHistoryParams struct {
	UserID    int64           `json:"userId"    validate:"required,gt=0"`
	Limit     int64           `json:"limit"     validate:"oneof=10 25 50 100"`
	SortOrder types.SortOrder `json:"sortOrder" validate:"omitempty,oneof=Asc Desc"`
	Cursor    string          `json:"cursor"    validate:"omitempty,base64"`
}

// UsernameHistoryBuilder builds parameters for GetUsernameHistory API call.
type UsernameHistoryBuilder struct {
	params UsernameHistoryParams
}

// NewUsernameHistoryBuilder creates a new UsernameHistoryBuilder with the given user ID.
func NewUsernameHistoryBuilder(userID int64) *UsernameHistoryBuilder {
	return &UsernameHistoryBuilder{
		params: UsernameHistoryParams{
			UserID:    userID,
			Limit:     10,
			SortOrder: "",
			Cursor:    "",
		},
	}
}

// WithLimit sets the maximum number of results to return.
func (b *UsernameHistoryBuilder) WithLimit(limit int64) *UsernameHistoryBuilder {
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
	b.params.SortOrder = types.SortOrderAsc
	return b
}

// WithSortOrderDesc sets the sort order for results to descending.
func (b *UsernameHistoryBuilder) WithSortOrderDesc() *UsernameHistoryBuilder {
	b.params.SortOrder = types.SortOrderDesc
	return b
}

// Build returns the UsernameHistoryParams.
func (b *UsernameHistoryBuilder) Build() UsernameHistoryParams {
	return b.params
}
