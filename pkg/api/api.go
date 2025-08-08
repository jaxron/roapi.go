package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/jaxron/axonet/pkg/client"
	"github.com/jaxron/roapi.go/pkg/api/middleware/auth"
	"github.com/jaxron/roapi.go/pkg/api/middleware/jsonheader"
	"github.com/jaxron/roapi.go/pkg/api/resources/avatar"
	"github.com/jaxron/roapi.go/pkg/api/resources/friends"
	"github.com/jaxron/roapi.go/pkg/api/resources/games"
	"github.com/jaxron/roapi.go/pkg/api/resources/groups"
	"github.com/jaxron/roapi.go/pkg/api/resources/inventory"
	"github.com/jaxron/roapi.go/pkg/api/resources/presence"
	"github.com/jaxron/roapi.go/pkg/api/resources/thumbnails"
	"github.com/jaxron/roapi.go/pkg/api/resources/users"
)

// API represents the main struct for interacting with the Roblox API.
// It contains a client for making HTTP requests and services for different API endpoints.
type API struct {
	client     *client.Client       // Axonet client for making API requests
	users      *users.Resource      // Resource for user-related API operations
	friends    *friends.Resource    // Resource for friend-related API operations
	groups     *groups.Resource     // Resource for group-related API operations
	thumbnails *thumbnails.Resource // Resource for thumbnail-related API operations
	avatar     *avatar.Resource     // Resource for avatar-related API operations
	presence   *presence.Resource   // Resource for presence-related API operations
	games      *games.Resource      // Resource for game-related API operations
	inventory  *inventory.Resource  // Resource for inventory-related API operations
}

// New creates a new instance of API with the provided options.
// It initializes the client and sets up the services.
//
//goland:noinspection GoUnusedExportedFunction
func New(cookies []string, opts ...client.Option) *API {
	// Initialize the client with custom options and middleware
	authMiddleware := auth.New(cookies)
	c := client.NewClient(append(
		[]client.Option{
			client.WithMiddleware(authMiddleware),
			client.WithMiddleware(jsonheader.New()),
		},
		opts...,
	)...)

	// Randomize the order of cookies for balancing
	authMiddleware.Shuffle()

	// Return a new API instance with initialized client and resources
	v := validator.New(validator.WithRequiredStructEnabled())

	return &API{
		client:     c,
		users:      users.New(c, v),
		friends:    friends.New(c, v),
		groups:     groups.New(c, v),
		thumbnails: thumbnails.New(c, v),
		avatar:     avatar.New(c, v),
		presence:   presence.New(c, v),
		games:      games.New(c, v),
		inventory:  inventory.New(c, v),
	}
}

// GetClient returns the Client instance used by the API.
// This can be useful for advanced users who need direct access to the client.
func (api *API) GetClient() *client.Client {
	return api.client
}

// Users returns the Resource instance for user-related operations.
// This provides access to methods for interacting with user data via the Roblox API.
func (api *API) Users() *users.Resource {
	return api.users
}

// Friends returns the Resource instance for friend-related operations.
// This provides access to methods for interacting with friend data via the Roblox API.
func (api *API) Friends() *friends.Resource {
	return api.friends
}

// Groups returns the Resource instance for group-related operations.
// This provides access to methods for interacting with group data via the Roblox API.
func (api *API) Groups() *groups.Resource {
	return api.groups
}

// Thumbnails returns the Resource instance for thumbnail-related operations.
// This provides access to methods for interacting with thumbnail data via the Roblox API.
func (api *API) Thumbnails() *thumbnails.Resource {
	return api.thumbnails
}

// Avatar returns the Resource instance for avatar-related operations.
// This provides access to methods for interacting with avatar data via the Roblox API.
func (api *API) Avatar() *avatar.Resource {
	return api.avatar
}

// Presence returns the Resource instance for presence-related operations.
// This provides access to methods for interacting with presence data via the Roblox API.
func (api *API) Presence() *presence.Resource {
	return api.presence
}

// Games returns the Resource instance for game-related operations.
// This provides access to methods for interacting with game data via the Roblox API.
func (api *API) Games() *games.Resource {
	return api.games
}

// Inventory returns the Resource instance for inventory-related operations.
// This provides access to methods for interacting with inventory data via the Roblox API.
func (api *API) Inventory() *inventory.Resource {
	return api.inventory
}
