package friends_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/services/friends"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindFriends(t *testing.T) {
	// Create a new test service
	api := friends.NewService(utils.NewTestClient())

	// Test case: Find friends for a known user
	t.Run("Find Known User Friends", func(t *testing.T) {
		friends, err := api.FindFriends(context.Background(), friends.NewFindFriendsBuilder(SampleUserID))
		require.NoError(t, err)
		assert.NotNil(t, friends)
		assert.NotEmpty(t, friends.PageItems)
	})

	// Test case: Attempt to find friends for a non-existent user
	t.Run("Find Non-existent User Friends", func(t *testing.T) {
		friends, err := api.FindFriends(context.Background(), friends.NewFindFriendsBuilder(InvalidUserID))
		require.Error(t, err)
		assert.Nil(t, friends)
	})
}
