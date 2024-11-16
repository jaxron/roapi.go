package presence

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jaxron/axonet/pkg/client"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// ResourceInterface defines the interface for presence-related operations.
type ResourceInterface interface {
	GetUserPresences(ctx context.Context, p UserPresencesParams) ([]types.UserPresenceResponse, error)
}

// Ensure Resource implements the ResourceInterface.
var _ ResourceInterface = (*Resource)(nil)

// Resource provides methods for interacting with presence-related endpoints.
type Resource struct {
	client   *client.Client
	validate *validator.Validate
}

// New creates a new Resource with the specified client and validator.
func New(client *client.Client, validate *validator.Validate) *Resource {
	return &Resource{
		client:   client,
		validate: validate,
	}
}
