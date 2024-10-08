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
	// Create a new test service and validator
	api := users.NewService(utils.NewTestEnv())

	t.Run("Fetch Known Users", func(t *testing.T) {
		usernames := []string{utils.SampleUsername4, utils.SampleUsername5}
		builder := users.NewGetUsersByUsernamesBuilder(usernames)
		result, err := api.GetUsersByUsernames(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, 2)

		// Check if the returned users match the requested usernames
		for _, user := range result {
			assert.Contains(t, usernames, user.Name)
		}
	})

	t.Run("Fetch With Non-existent Username", func(t *testing.T) {
		usernames := []string{utils.SampleUsername4, utils.InvalidUsername}
		builder := users.NewGetUsersByUsernamesBuilder(usernames)
		result, err := api.GetUsersByUsernames(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, 1) // Only one user should be returned

		assert.Equal(t, utils.SampleUsername4, result[0].Name)
	})
}
