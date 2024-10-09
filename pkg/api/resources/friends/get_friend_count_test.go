package friends_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/friends"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFriendCount(t *testing.T) {
	// Create a new test resource
	api := friends.New(utils.NewTestEnv())

	// Test case: Fetch friend count for a known user
	t.Run("Fetch Known User Friend Count", func(t *testing.T) {
		count, err := api.GetFriendCount(context.Background(), utils.SampleUserID1)
		require.NoError(t, err)
		assert.NotZero(t, count)
	})

	// Test case: Attempt to fetch friend count for a non-existent user
	t.Run("Fetch Non-existent User Friend Count", func(t *testing.T) {
		count, err := api.GetFriendCount(context.Background(), utils.InvalidUserID)
		require.Error(t, err)
		assert.Zero(t, count)
	})
}
