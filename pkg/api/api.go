package api

import (
	"github.com/jaxron/axonet/pkg/client"
	"github.com/jaxron/roapi.go/internal/middleware/auth"
	"github.com/jaxron/roapi.go/pkg/api/services/friends"
	"github.com/jaxron/roapi.go/pkg/api/services/users"
)

// API represents the main struct for interacting with the Roblox API.
// It contains a client for making HTTP requests and services for different API endpoints.
type API struct {
	client  *client.Client   // Axonet client for making API requests
	users   *users.Service   // Service for user-related API operations
	friends *friends.Service // Service for friend-related API operations
}

// New creates a new instance of API with the provided options.
// It initializes the client and sets up the services.
func New(cookies []string, opts ...client.Option) *API {
	c := client.NewClient(append(
		opts,
		client.WithMiddleware(auth.New(cookies)),
	)...)

	return &API{
		client:  c,
		users:   users.NewService(c),
		friends: friends.NewService(c),
	}
}

// GetClient returns the Client instance used by the API.
// This can be useful for advanced users who need direct access to the client.
func (api *API) GetClient() *client.Client {
	return api.client
}

// Users returns the Service instance for user-related operations.
// This provides access to methods for interacting with user data via the Roblox API.
func (api *API) Users() *users.Service {
	return api.users
}

// Friends returns the Service instance for friend-related operations.
// This provides access to methods for interacting with friend data via the Roblox API.
func (api *API) Friends() *friends.Service {
	return api.friends
}
