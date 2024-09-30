package users_test

import (
	"math"
	"testing"

	"github.com/jaxron/roapi.go/pkg/api/client"
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

	// Assert that the client in the service is the same as the mock client
	assert.Equal(t, client, service.Client, "Service client should be the same as the provided client")
}
