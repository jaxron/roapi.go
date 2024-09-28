package users_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/services/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetUsersByUsernames tests the GetUsersByUsernames method of the user.Service.
func TestGetUsersByUsernames(t *testing.T) {
	// Create a new test service
	api := users.NewService(utils.NewTestClient(true, false))

	t.Run("Fetch Known Users", func(t *testing.T) {
		usernames := []string{"Roblox", "builderman"}
		builder := users.NewUsersByUsernamesBuilder(usernames)
		result, err := api.GetUsersByUsernames(context.Background(), builder)

		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 2)

		// Check if the returned users match the requested usernames
		for _, user := range result.Data {
			assert.Contains(t, usernames, user.Name)
		}
	})

	t.Run("Fetch With Non-existent Username", func(t *testing.T) {
		usernames := []string{"Roblox", InvalidUsername}
		builder := users.NewUsersByUsernamesBuilder(usernames)
		result, err := api.GetUsersByUsernames(context.Background(), builder)

		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 1) // Only one user should be returned

		assert.Equal(t, "Roblox", result.Data[0].Name)
	})
}
