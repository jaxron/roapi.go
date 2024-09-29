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
	// Create a new test service
	api := friends.NewService(utils.NewTestClient(true, true))

	// Test case: Search for friends for a known user
	t.Run("Search Known User Friends", func(t *testing.T) {
		friends, err := api.SearchFriends(context.Background(), friends.NewSearchFriendsBuilder(SampleUserID).Query("laugh"))
		require.NoError(t, err)
		assert.NotNil(t, friends)
		assert.NotEmpty(t, friends)
	})

	// Test case: Attempt to search for friends for a non-existent user
	t.Run("Search Non-existent User Friends", func(t *testing.T) {
		friends, err := api.SearchFriends(context.Background(), friends.NewSearchFriendsBuilder(InvalidUserID))
		require.Error(t, err)
		assert.Nil(t, friends)
	})
}
