package friends_test

import (
	"math"
	"testing"

	"github.com/jaxron/roapi.go/pkg/api/services/friends"
	"github.com/jaxron/roapi.go/pkg/client"
	"github.com/stretchr/testify/assert"
)

const (
	SampleUserID  = 339310190
	InvalidUserID = uint64(math.MaxUint64)
)

func TestNewService(t *testing.T) {
	// Create a standard client
	client := client.NewClient()

	// Create a new service
	service := friends.NewService(client)

	// Assert that the service is not nil
	assert.NotNil(t, service, "Service should not be nil")

	// Assert that the client in the service is the same as the mock client
	assert.Equal(t, client, service.Client, "Service client should be the same as the provided client")
}
