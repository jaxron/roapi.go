package games

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetUserFavoriteGames fetches favorite games for a specific user.
// GET https://games.roblox.com/v2/users/{userId}/favorite/games
func (r *Resource) GetUserFavoriteGames(ctx context.Context, p UserFavoriteGamesParams) (*types.GameResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var result types.GameResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v2/users/%d/favorite/games", types.GamesEndpoint, p.UserID)).
		Query("accessFilter", strconv.FormatUint(uint64(p.AccessFilter), 10)).
		Query("limit", strconv.FormatUint(p.Limit, 10)).
		Query("cursor", p.Cursor).
		Query("sortOrder", string(types.SortOrderDesc)).
		Result(&result).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if err := r.validate.Struct(&result); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &result, nil
}

// UserFavoriteGamesParams holds the parameters for fetching user's favorite games.
type UserFavoriteGamesParams struct {
	UserID       uint64       `validate:"required,gt=0"`
	AccessFilter AccessFilter `validate:"oneof=1 2 4"`
	Limit        uint64       `validate:"oneof=10 25 50 100"`
	Cursor       string       `validate:"omitempty"`
}

// UserFavoriteGamesBuilder helps build parameters for the GetUserFavoriteGames API call.
type UserFavoriteGamesBuilder struct {
	params UserFavoriteGamesParams
}

// NewUserFavoriteGamesBuilder creates a new UserFavoriteGamesBuilder with default values.
func NewUserFavoriteGamesBuilder(userID uint64) *UserFavoriteGamesBuilder {
	return &UserFavoriteGamesBuilder{
		params: UserFavoriteGamesParams{
			UserID:       userID,
			AccessFilter: AccessFilterPublic,
			Limit:        10,
			Cursor:       "",
		},
	}
}

// WithAccessFilter sets the access filter.
func (b *UserFavoriteGamesBuilder) WithAccessFilter(filter AccessFilter) *UserFavoriteGamesBuilder {
	b.params.AccessFilter = filter
	return b
}

// WithLimit sets the maximum number of results to return.
func (b *UserFavoriteGamesBuilder) WithLimit(limit uint64) *UserFavoriteGamesBuilder {
	b.params.Limit = limit
	return b
}

// WithCursor sets the cursor for pagination.
func (b *UserFavoriteGamesBuilder) WithCursor(cursor string) *UserFavoriteGamesBuilder {
	b.params.Cursor = cursor
	return b
}

// Build returns the UserFavoriteGamesParams.
func (b *UserFavoriteGamesBuilder) Build() UserFavoriteGamesParams {
	return b.params
}
