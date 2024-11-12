package presence_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/presence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserPresences(t *testing.T) {
	// Create a new test resource
	api := presence.New(utils.NewTestEnv())

	t.Run("Fetch Known Users Presence", func(t *testing.T) {
		userIDs := []uint64{utils.SampleUserID1, utils.SampleUserID2}
		builder := presence.NewUserPresencesBuilder(userIDs...)
		result, err := api.GetUserPresences(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, 2)

		for _, presence := range result {
			assert.Contains(t, userIDs, presence.UserID)
			assert.NotZero(t, presence.LastOnline)
			assert.NotEmpty(t, presence.LastLocation)
		}
	})

	t.Run("Fetch With Non-existent User ID", func(t *testing.T) {
		builder := presence.NewUserPresencesBuilder(utils.InvalidUserID)
		result, err := api.GetUserPresences(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.Len(t, result, 0)
	})

	t.Run("Test Builder Methods", func(t *testing.T) {
		builder := presence.NewUserPresencesBuilder().
			WithUserIDs(1, 2, 3, 4).
			RemoveUserIDs(2, 3)

		params := builder.Build()
		assert.Len(t, params.UserIDs, 2)
		assert.Contains(t, params.UserIDs, uint64(1))
		assert.Contains(t, params.UserIDs, uint64(4))
		assert.NotContains(t, params.UserIDs, uint64(2))
		assert.NotContains(t, params.UserIDs, uint64(3))
	})
}
