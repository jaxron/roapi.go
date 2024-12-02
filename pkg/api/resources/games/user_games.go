package games

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetUserGames fetches games for a specific user.
// GET https://games.roblox.com/v2/users/{userId}/games
func (r *Resource) GetUserGames(ctx context.Context, p UserGamesParams) (*types.GameResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var result types.GameResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v2/users/%d/games", types.GamesEndpoint, p.UserID)).
		Query("accessFilter", strconv.FormatUint(uint64(p.AccessFilter), 10)).
		Query("limit", strconv.FormatUint(p.Limit, 10)).
		Query("cursor", p.Cursor).
		Query("sortOrder", string(p.SortOrder)).
		Result(&result).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	if err := r.validate.Struct(&result); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &result, nil
}

// AccessFilter represents the filter type for game access.
type AccessFilter uint8

const (
	AccessFilterUnknown1 AccessFilter = 1
	AccessFilterPublic   AccessFilter = 2
	AccessFilterUnknown4 AccessFilter = 4
)

// UserGamesParams holds the parameters for fetching user games.
type UserGamesParams struct {
	UserID       uint64          `validate:"required,gt=0"`
	AccessFilter AccessFilter    `validate:"oneof=1 2 4"`
	Limit        uint64          `validate:"oneof=10 25 50"`
	Cursor       string          `validate:"omitempty"`
	SortOrder    types.SortOrder `validate:"oneof=Asc Desc"`
}

// UserGamesBuilder helps build parameters for the GetUserGames API call.
type UserGamesBuilder struct {
	params UserGamesParams
}

// NewUserGamesBuilder creates a new UserGamesBuilder with default values.
func NewUserGamesBuilder(userID uint64) *UserGamesBuilder {
	return &UserGamesBuilder{
		params: UserGamesParams{
			UserID:       userID,
			AccessFilter: AccessFilterPublic,
			Limit:        50,
			Cursor:       "",
			SortOrder:    types.SortOrderAsc,
		},
	}
}

// WithAccessFilter sets the access filter.
func (b *UserGamesBuilder) WithAccessFilter(filter AccessFilter) *UserGamesBuilder {
	b.params.AccessFilter = filter
	return b
}

// WithLimit sets the maximum number of results to return.
func (b *UserGamesBuilder) WithLimit(limit uint64) *UserGamesBuilder {
	b.params.Limit = limit
	return b
}

// WithCursor sets the cursor for pagination.
func (b *UserGamesBuilder) WithCursor(cursor string) *UserGamesBuilder {
	b.params.Cursor = cursor
	return b
}

// WithSortOrder sets the sort order for results.
func (b *UserGamesBuilder) WithSortOrder(order types.SortOrder) *UserGamesBuilder {
	b.params.SortOrder = order
	return b
}

// Build returns the UserGamesParams.
func (b *UserGamesBuilder) Build() UserGamesParams {
	return b.params
}
