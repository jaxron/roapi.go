package friends_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/services/friends"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetFollowers tests the GetFollowers method of the friends.Service.
func TestGetFollowers(t *testing.T) {
	// Create a new test service
	api := friends.NewService(utils.NewTestClient())

	// Test case: Fetch followers for a known user
	t.Run("Fetch Known User Followers", func(t *testing.T) {
		followers, err := api.GetFollowers(context.Background(), friends.NewFollowersBuilder(SampleUserID))
		require.NoError(t, err)
		assert.NotNil(t, followers)
		assert.NotEmpty(t, followers)
	})

	// Test case: Attempt to fetch followers for a non-existent user
	t.Run("Fetch Non-existent User Followers", func(t *testing.T) {
		followers, err := api.GetFollowers(context.Background(), friends.NewFollowersBuilder(InvalidUserID))
		require.Error(t, err)
		assert.Nil(t, followers)
	})
}
