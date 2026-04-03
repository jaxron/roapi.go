package friends_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/friends"
	"github.com/jaxron/roapi.go/pkg/api/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFriends(t *testing.T) {
	// Create a new test resource
	api := friends.New(utils.NewTestEnv())

	// Test case: Fetch friends for a known user
	t.Run("Fetch Known User Friends", func(t *testing.T) {
		builder := friends.NewGetFriendsBuilder(utils.SampleUserID1)
		result, err := api.GetFriends(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.Data)
	})

	// Test case: Fetch friends sorted by interaction frequency
	t.Run("Fetch Friends By StatusFrequents", func(t *testing.T) {
		builder := friends.NewGetFriendsBuilder(utils.SampleUserID1).
			WithUserSort(types.FriendSortStatusFrequents)
		result, err := api.GetFriends(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.Data)
	})

	// Test case: Attempt to fetch friends for a non-existent user
	t.Run("Fetch Non-existent User Friends", func(t *testing.T) {
		builder := friends.NewGetFriendsBuilder(utils.InvalidUserID)
		_, err := api.GetFriends(context.Background(), builder.Build())
		require.Error(t, err)
	})

	// Test case: Valid parameters with all fields set
	t.Run("Valid Parameters", func(t *testing.T) {
		builder := friends.NewGetFriendsBuilder(utils.SampleUserID1).
			WithUserSort(types.FriendSortStatusFrequents)

		params := builder.Build()
		assert.Equal(t, utils.SampleUserID1, params.UserID)
		assert.Equal(t, types.FriendSortStatusFrequents, params.UserSort)
	})

	// Test case: Default parameters
	t.Run("Default Parameters", func(t *testing.T) {
		builder := friends.NewGetFriendsBuilder(utils.SampleUserID1)

		params := builder.Build()
		assert.Equal(t, utils.SampleUserID1, params.UserID)
		assert.Equal(t, types.FriendSortDefault, params.UserSort)
	})
}
