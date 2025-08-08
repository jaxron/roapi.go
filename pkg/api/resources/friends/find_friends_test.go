package friends_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/friends"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindFriends(t *testing.T) {
	// Create a new test resource and validator
	api := friends.New(utils.NewTestEnv())

	// Test case: Find friends for a known user
	t.Run("Find Known User Friends", func(t *testing.T) {
		builder := friends.NewFindFriendsBuilder(utils.SampleUserID1)
		result, err := api.FindFriends(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.PageItems, 2)
	})

	// Test case: Attempt to find friends for a non-existent user
	t.Run("Find Non-existent User Friends", func(t *testing.T) {
		builder := friends.NewFindFriendsBuilder(utils.InvalidUserID)
		result, err := api.FindFriends(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Nil(t, result)
	})

	// Test case: Validate with invalid UserSort
	t.Run("Invalid UserSort", func(t *testing.T) {
		builder := friends.NewFindFriendsBuilder(utils.SampleUserID1).WithUserSort(3)
		_, err := api.FindFriends(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "UserSort")
	})

	// Test case: Validate with invalid Limit
	t.Run("Invalid Limit", func(t *testing.T) {
		builder := friends.NewFindFriendsBuilder(utils.SampleUserID1).WithLimit(100)
		_, err := api.FindFriends(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Limit")
	})

	// Test case: Valid parameters with all fields set
	t.Run("Valid Parameters", func(t *testing.T) {
		builder := friends.NewFindFriendsBuilder(utils.SampleUserID1).
			WithUserSort(1).
			WithLimit(20).
			WithCursor("someCursor")

		params := builder.Build()
		assert.Equal(t, utils.SampleUserID1, params.UserID)
		assert.Equal(t, uint64(1), params.UserSort)
		assert.Equal(t, uint64(20), params.Limit)
		assert.Equal(t, "someCursor", params.Cursor)
	})
}
