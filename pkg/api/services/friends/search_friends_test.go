package friends_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/services/friends"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSearchFriends tests the SearchFriends method of the friends.Service.
func TestSearchFriends(t *testing.T) {
	// Create a new test service and validator
	api := friends.NewService(utils.NewTestEnv())

	// Test case: Search for friends for a known user
	t.Run("Search Known User Friends", func(t *testing.T) {
		builder := friends.NewSearchFriendsBuilder(SampleUserID).WithQuery("laugh")
		friends, err := api.SearchFriends(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, friends)
		assert.NotEmpty(t, friends.PageItems)
	})

	// Test case: Attempt to search for friends for a non-existent user
	t.Run("Search Non-existent User Friends", func(t *testing.T) {
		builder := friends.NewSearchFriendsBuilder(InvalidUserID)
		friends, err := api.SearchFriends(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Nil(t, friends)
	})

	// Test case: Validate with invalid Cursor
	t.Run("Invalid Cursor", func(t *testing.T) {
		builder := friends.NewSearchFriendsBuilder(SampleUserID).WithCursor("invalidCursor")
		_, err := api.SearchFriends(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Cursor")
	})

	// Test case: Validate with invalid Limit
	t.Run("Invalid Limit", func(t *testing.T) {
		builder := friends.NewSearchFriendsBuilder(SampleUserID).WithLimit(51)
		_, err := api.SearchFriends(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Limit")
	})

	// Test case: Valid parameters with all fields set
	t.Run("Valid Parameters", func(t *testing.T) {
		builder := friends.NewSearchFriendsBuilder(SampleUserID).
			WithQuery("test").
			WithLimit(50).
			WithCursor("someCursor")

		params := builder.Build()
		assert.Equal(t, uint64(SampleUserID), params.UserID)
		assert.Equal(t, "test", params.Query)
		assert.Equal(t, uint64(50), params.Limit)
		assert.Equal(t, "someCursor", params.Cursor)
	})
}
