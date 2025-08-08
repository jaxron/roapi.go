package users_test

import (
	"context"
	"math"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUsersByIDs(t *testing.T) {
	// Create a new test resource
	api := users.New(utils.NewTestEnv())

	t.Run("Fetch Known Users", func(t *testing.T) {
		userIDs := []int64{utils.SampleUserID4, utils.SampleUserID5}
		builder := users.NewUsersByIDsBuilder(userIDs...)
		result, err := api.GetUsersByIDs(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 2)

		for _, user := range result.Data {
			assert.Contains(t, userIDs, user.ID)
		}
	})

	t.Run("Fetch With Non-existent User ID", func(t *testing.T) {
		userIDs := []int64{utils.SampleUserID1, math.MaxInt64}
		builder := users.NewUsersByIDsBuilder(userIDs...)
		result, err := api.GetUsersByIDs(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 1) // Only one user should be returned

		assert.Equal(t, utils.SampleUserID1, result.Data[0].ID)
	})

	t.Run("Test Builder Methods", func(t *testing.T) {
		builder := users.NewUsersByIDsBuilder().
			WithUserIDs(1, 2, 3, 4).
			RemoveUserIDs(3)

		params := builder.Build()
		assert.Len(t, params.UserIDs, 3)
		assert.Contains(t, params.UserIDs, int64(1))
		assert.Contains(t, params.UserIDs, int64(2))
		assert.Contains(t, params.UserIDs, int64(4))
		assert.NotContains(t, params.UserIDs, int64(3))
	})
}
