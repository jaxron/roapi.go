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
	api := users.NewService(utils.NewTestClient(true, false))

	// Test case: Fetch information for a known user
	t.Run("Fetch Known User", func(t *testing.T) {
		userID := uint64(1)
		user, err := api.GetUserByID(context.Background(), 1)
		require.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, userID, user.ID)
		assert.Equal(t, "Roblox", user.Name)
	})

	// Test case: Attempt to fetch information for a non-existent user
	t.Run("Fetch Non-existent User", func(t *testing.T) {
		userID := InvalidUserID
		user, err := api.GetUserByID(context.Background(), userID)
		require.Error(t, err)
		assert.Nil(t, user)
	})
}
