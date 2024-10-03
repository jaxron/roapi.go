package users_test

import (
	"math"
	"testing"

	"github.com/jaxron/axonet/pkg/client"
	"github.com/jaxron/roapi.go/pkg/api/services/users"
	"github.com/stretchr/testify/assert"
)

const (
	InvalidUserID   = uint64(math.MaxUint64)
	InvalidUsername = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
)

func TestNewService(t *testing.T) {
	// Create a standard client
	client := client.NewClient()

	// Create a new service
	service := users.NewService(client)

	// Assert that the service is not nil
	assert.NotNil(t, service, "Service should not be nil")
}
