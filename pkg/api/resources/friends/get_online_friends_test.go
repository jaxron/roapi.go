package friends_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/friends"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetOnlineFriends(t *testing.T) {
	// Create a new test resource
	api := friends.New(utils.NewTestEnv())

	// Test case: Fetch online friends for a known user
	t.Run("Fetch Known User Online Friends", func(t *testing.T) {
		builder := friends.NewGetOnlineFriendsBuilder(utils.SampleUserID1)
		onlineFriends, err := api.GetOnlineFriends(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, onlineFriends)
	})

	// Test case: Attempt to fetch online friends for a non-existent user
	t.Run("Fetch Non-existent User Online Friends", func(t *testing.T) {
		builder := friends.NewGetOnlineFriendsBuilder(utils.InvalidUserID)
		onlineFriends, err := api.GetOnlineFriends(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Nil(t, onlineFriends)
	})

	// Test case: Validate with invalid UserSort
	t.Run("Invalid UserSort", func(t *testing.T) {
		builder := friends.NewGetOnlineFriendsBuilder(utils.SampleUserID1).WithUserSort(3)
		_, err := api.GetOnlineFriends(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "UserSort")
	})

	// Test case: Valid parameters with all fields set
	t.Run("Valid Parameters", func(t *testing.T) {
		builder := friends.NewGetOnlineFriendsBuilder(utils.SampleUserID1).
			WithUserSort(2)

		params := builder.Build()
		assert.Equal(t, utils.SampleUserID1, params.UserID)
		assert.Equal(t, uint64(2), params.UserSort)
	})
}
