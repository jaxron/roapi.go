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
	api := users.NewService(utils.NewTestEnv())

	t.Run("Fetch Known Users", func(t *testing.T) {
		userIDs := []uint64{utils.SampleUserID4, utils.SampleUserID5}
		builder := users.NewUsersByIDsBuilder(userIDs...)
		result, err := api.GetUsersByIDs(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, 2)

		for _, user := range result {
			assert.Contains(t, userIDs, user.ID)
		}
	})

	t.Run("Fetch With Non-existent User ID", func(t *testing.T) {
		userIDs := []uint64{utils.SampleUserID1, math.MaxUint64}
		builder := users.NewUsersByIDsBuilder(userIDs...)
		result, err := api.GetUsersByIDs(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, 1) // Only one user should be returned

		assert.Equal(t, utils.SampleUserID1, result[0].ID)
	})

	t.Run("Test Builder Methods", func(t *testing.T) {
		builder := users.NewUsersByIDsBuilder().
			WithUserIDs(1, 2, 3, 4).
			RemoveUserIDs(3)

		params := builder.Build()
		assert.Len(t, params.UserIDs, 3)
		assert.Contains(t, params.UserIDs, 1)
		assert.Contains(t, params.UserIDs, 2)
		assert.Contains(t, params.UserIDs, 4)
		assert.NotContains(t, params.UserIDs, 3)
	})
}
