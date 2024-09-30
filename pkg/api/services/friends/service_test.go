package friends_test

import (
	"math"
	"testing"

	"github.com/jaxron/roapi.go/pkg/api/client"
	"github.com/jaxron/roapi.go/pkg/api/services/friends"
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
}
