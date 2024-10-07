package users_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/services/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetUserByID tests the GetUserByID method of the user.Service.
func TestGetUserByID(t *testing.T) {
	// Create a new test service
	api := users.NewService(utils.NewTestEnv())

	// Test case: Fetch information for a known user
	t.Run("Fetch Known User", func(t *testing.T) {
		user, err := api.GetUserByID(context.Background(), SampleUserID)
		require.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, SampleUserID, user.ID)
	})

	// Test case: Attempt to fetch information for a non-existent user
	t.Run("Fetch Non-existent User", func(t *testing.T) {
		user, err := api.GetUserByID(context.Background(), InvalidUserID)
		require.Error(t, err)
		assert.Nil(t, user)
	})
}
