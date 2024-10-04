package friends_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/services/friends"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetFollowerCount tests the GetFollowerCount method of the friends.Service.
func TestGetFollowerCount(t *testing.T) {
	// Create a new test service
	api := friends.NewService(utils.NewTestClient())

	// Test case: Fetch follower count for a known user
	t.Run("Fetch Known User Follower Count", func(t *testing.T) {
		userID := uint64(SampleUserID)
		count, err := api.GetFollowerCount(context.Background(), userID)
		require.NoError(t, err)
		assert.NotZero(t, count)
	})

	// Test case: Attempt to fetch follower count for a non-existent user
	t.Run("Fetch Non-existent User Follower Count", func(t *testing.T) {
		userID := InvalidUserID
		count, err := api.GetFollowerCount(context.Background(), userID)
		require.Error(t, err)
		assert.Zero(t, count)
	})
}
