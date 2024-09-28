package users_test

import (
	"context"
	"math"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/services/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetUsersByIDs tests the GetUsersByIDs method of the user.Service.
func TestGetUsersByIDs(t *testing.T) {
	// Create a new test service
	api := users.NewService(utils.NewTestClient(true, false))

	t.Run("Fetch Known Users", func(t *testing.T) {
		userIDs := []uint64{1, 156} // IDs for Roblox and Builderman
		builder := users.NewUsersByIDsBuilder(userIDs)
		result, err := api.GetUsersByIDs(context.Background(), builder)

		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 2)

		// Check if the returned users match the requested IDs
		for _, user := range result.Data {
			assert.Contains(t, userIDs, user.ID)
		}
	})

	t.Run("Fetch With Non-existent User ID", func(t *testing.T) {
		userIDs := []uint64{1, math.MaxUint64}
		builder := users.NewUsersByIDsBuilder(userIDs)
		result, err := api.GetUsersByIDs(context.Background(), builder)

		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 1) // Only one user should be returned

		assert.Equal(t, uint64(1), result.Data[0].ID)
	})
}
