package friends_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/services/friends"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetFollowingCount tests the GetFollowingCount method of the friends.Service.
func TestGetFollowingCount(t *testing.T) {
	// Create a new test service
	api := friends.NewService(utils.NewTestEnv())

	// Test case: Fetch following count for a known user
	t.Run("Fetch Known User Following Count", func(t *testing.T) {
		count, err := api.GetFollowingCount(context.Background(), utils.SampleUserID1)
		require.NoError(t, err)
		assert.NotZero(t, count)
	})

	// Test case: Attempt to fetch following count for a non-existent user
	t.Run("Fetch Non-existent User Following Count", func(t *testing.T) {
		count, err := api.GetFollowingCount(context.Background(), utils.InvalidUserID)
		require.Error(t, err)
		assert.Zero(t, count)
	})
}
