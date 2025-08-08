package games

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// ServerType represents the type of game server.
type ServerType int32

const (
	ServerTypePublic  ServerType = 0 // Public game server
	ServerTypePrivate ServerType = 1 // Private game server
)

// SortOrder represents the sort order for server listings.
type SortOrder int32

const (
	SortOrderAsc  SortOrder = 1 // Ascending sort order
	SortOrderDesc SortOrder = 2 // Descending sort order
)

// GetGameServers fetches the list of servers for a specific place.
// GET https://games.roblox.com/v1/games/{placeId}/servers/{serverType}
func (r *Resource) GetGameServers(ctx context.Context, p GameServersParams) (*types.ServerResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var result types.ServerResponse

	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/games/%d/servers/%d", types.GamesEndpoint, p.PlaceID, p.ServerType)).
		Query("sortOrder", strconv.Itoa(int(p.SortOrder))).
		Query("excludeFullGames", strconv.FormatBool(p.ExcludeFullGames)).
		Query("limit", strconv.Itoa(int(p.Limit))).
		Query("cursor", p.Cursor).
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

// GameServersParams holds the parameters for fetching game servers.
type GameServersParams struct {
	PlaceID          int64      `validate:"required,gt=0"`      // Required: ID of the place to fetch servers for
	ServerType       ServerType `validate:"oneof=0 1"`          // Required: Type of servers to fetch (public/private)
	SortOrder        SortOrder  `validate:"required,oneof=1 2"` // Required: Sort order for the server list
	ExcludeFullGames bool       `validate:""`                   // Optional: Whether to exclude full games
	Limit            int32      `validate:"min=10,max=100"`     // Required: Number of results per request
	Cursor           string     `validate:"omitempty"`          // Optional: Cursor for pagination
}

// GameServersBuilder helps build parameters for the GetGameServers API call.
type GameServersBuilder struct {
	params GameServersParams
}

// NewGameServersBuilder creates a new GameServersBuilder with default values.
func NewGameServersBuilder(placeID int64) *GameServersBuilder {
	return &GameServersBuilder{
		params: GameServersParams{
			PlaceID:          placeID,
			ServerType:       ServerTypePublic,
			SortOrder:        SortOrderAsc,
			ExcludeFullGames: false,
			Limit:            10,
			Cursor:           "",
		},
	}
}

// WithServerType sets the server type.
func (b *GameServersBuilder) WithServerType(serverType ServerType) *GameServersBuilder {
	b.params.ServerType = serverType
	return b
}

// WithSortOrder sets the sort order.
func (b *GameServersBuilder) WithSortOrder(sortOrder SortOrder) *GameServersBuilder {
	b.params.SortOrder = sortOrder
	return b
}

// WithExcludeFullGames sets whether to exclude full games.
func (b *GameServersBuilder) WithExcludeFullGames(exclude bool) *GameServersBuilder {
	b.params.ExcludeFullGames = exclude
	return b
}

// WithLimit sets the maximum number of results to return.
func (b *GameServersBuilder) WithLimit(limit int32) *GameServersBuilder {
	b.params.Limit = limit
	return b
}

// WithCursor sets the cursor for pagination.
func (b *GameServersBuilder) WithCursor(cursor string) *GameServersBuilder {
	b.params.Cursor = cursor
	return b
}

// Build returns the GameServersParams.
func (b *GameServersBuilder) Build() GameServersParams {
	return b.params
}
