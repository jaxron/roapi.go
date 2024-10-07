package users_test

import (
	"math"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/jaxron/axonet/pkg/client"
	"github.com/jaxron/roapi.go/pkg/api/services/users"
	"github.com/stretchr/testify/assert"
)

const (
	SampleUserID    = uint64(339310190)
	InvalidUserID   = uint64(math.MaxUint64)
	InvalidUsername = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
)

func TestNewService(t *testing.T) {
	// Create a new service
	service := users.NewService(client.NewClient(), validator.New())

	// Assert that the service is not nil
	assert.NotNil(t, service, "Service should not be nil")
}
