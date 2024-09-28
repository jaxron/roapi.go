package friends_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/services/friends"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFriends(t *testing.T) {
	// Create a new test service
	api := friends.NewService(utils.NewTestClient(true, false))

	// Test case: Fetch friends for a known user
	t.Run("Fetch Known User Friends", func(t *testing.T) {
		friends, err := api.GetFriends(context.Background(), SampleUserID)
		require.NoError(t, err)
		assert.NotNil(t, friends)
		assert.NotEmpty(t, friends.Data)
	})

	// Test case: Attempt to fetch friends for a non-existent user
	t.Run("Fetch Non-existent User Friends", func(t *testing.T) {
		friends, err := api.GetFriends(context.Background(), InvalidUserID)
		require.Error(t, err)
		assert.Nil(t, friends)
	})
}
