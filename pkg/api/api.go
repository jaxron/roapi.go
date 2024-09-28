package api

import (
	"github.com/jaxron/roapi.go/pkg/api/services/user"
	"github.com/jaxron/roapi.go/pkg/client"
)

// API represents the main struct for interacting with the Roblox API.
// It contains a client for making HTTP requests and services for different API endpoints.
type API struct {
	client *client.Client // HTTP client for making API requests
	user   *user.Service  // Service for user-related API operations
}

// New creates a new instance of API with the provided options.
// It initializes the client and sets up the services.
func New(opts ...client.Option) *API {
	// Create a new client with the provided options
	c := client.NewClient(opts...)

	// Initialize and return the API instance
	return &API{
		client: c,
		user:   user.NewService(c),
	}
}

// GetClient returns the Client instance used by the API.
// This can be useful for advanced users who need direct access to the client.
func (api *API) GetClient() *client.Client {
	return api.client
}

// User returns the Service instance for user-related operations.
// This provides access to methods for interacting with user data via the Roblox API.
func (api *API) User() *user.Service {
	return api.user
}
