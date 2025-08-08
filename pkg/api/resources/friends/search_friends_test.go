package friends_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/friends"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSearchFriends(t *testing.T) {
	// Create a new test resource and validator
	api := friends.New(utils.NewTestEnv())

	// Test case: Search for friends for a known user
	t.Run("Search Known User Friends", func(t *testing.T) {
		builder := friends.NewSearchFriendsBuilder(utils.SampleUserID1).WithQuery(utils.SampleUsername2)
		result, err := api.SearchFriends(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.PageItems, 1)
	})

	// Test case: Attempt to search for friends for a non-existent user
	t.Run("Search Non-existent User Friends", func(t *testing.T) {
		builder := friends.NewSearchFriendsBuilder(utils.InvalidUserID)
		result, err := api.SearchFriends(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Nil(t, result)
	})

	// Test case: Validate with invalid Cursor
	t.Run("Invalid Cursor", func(t *testing.T) {
		builder := friends.NewSearchFriendsBuilder(utils.SampleUserID1).WithCursor("invalidCursor")
		_, err := api.SearchFriends(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Cursor")
	})

	// Test case: Validate with invalid Limit
	t.Run("Invalid Limit", func(t *testing.T) {
		builder := friends.NewSearchFriendsBuilder(utils.SampleUserID1).WithLimit(51)
		_, err := api.SearchFriends(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Limit")
	})

	// Test case: Valid parameters with all fields set
	t.Run("Valid Parameters", func(t *testing.T) {
		builder := friends.NewSearchFriendsBuilder(utils.SampleUserID1).
			WithQuery("test").
			WithLimit(50).
			WithCursor("someCursor")

		params := builder.Build()
		assert.Equal(t, utils.SampleUserID1, params.UserID)
		assert.Equal(t, "test", params.Query)
		assert.Equal(t, int64(50), params.Limit)
		assert.Equal(t, "someCursor", params.Cursor)
	})
}
