package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/jaxron/axonet/pkg/client"
	"github.com/jaxron/roapi.go/internal/middleware/auth"
	"github.com/jaxron/roapi.go/internal/middleware/jsonheader"
	"github.com/jaxron/roapi.go/pkg/api/resources/friends"
	"github.com/jaxron/roapi.go/pkg/api/resources/groups"
	"github.com/jaxron/roapi.go/pkg/api/resources/users"
)

// API represents the main struct for interacting with the Roblox API.
// It contains a client for making HTTP requests and services for different API endpoints.
type API struct {
	client  *client.Client    // Axonet client for making API requests
	users   *users.Resource   // Resource for user-related API operations
	friends *friends.Resource // Resource for friend-related API operations
	groups  *groups.Resource  // Resource for group-related API operations
}

// New creates a new instance of API with the provided options.
// It initializes the client and sets up the services.
func New(cookies []string, opts ...client.Option) *API {
	c := client.NewClient(append(
		opts,
		client.WithMiddleware(auth.New(cookies)),
		client.WithMiddleware(jsonheader.New()),
	)...)
	v := validator.New(validator.WithRequiredStructEnabled())

	return &API{
		client:  c,
		users:   users.New(c, v),
		friends: friends.New(c, v),
		groups:  groups.New(c, v),
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
