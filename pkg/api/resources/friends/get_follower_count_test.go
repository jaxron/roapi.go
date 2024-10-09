package friends_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/friends"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFollowerCount(t *testing.T) {
	// Create a new test resource
	api := friends.New(utils.NewTestEnv())

	// Test case: Fetch follower count for a known user
	t.Run("Fetch Known User Follower Count", func(t *testing.T) {
		count, err := api.GetFollowerCount(context.Background(), utils.SampleUserID1)
		require.NoError(t, err)
		assert.NotZero(t, count)
	})

	// Test case: Attempt to fetch follower count for a non-existent user
	t.Run("Fetch Non-existent User Follower Count", func(t *testing.T) {
		count, err := api.GetFollowerCount(context.Background(), utils.InvalidUserID)
		require.Error(t, err)
		assert.Zero(t, count)
	})
}
