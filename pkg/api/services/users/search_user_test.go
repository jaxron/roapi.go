package users_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/services/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSearchUser tests the SearchUser method of the user.Service.
func TestSearchUser(t *testing.T) {
	// Create a new test service
	api := users.NewService(utils.NewTestClient(true, false))

	// Test case: Search for a known user
	t.Run("Search Known User", func(t *testing.T) {
		username := "Roblox"
		res, err := api.SearchUser(context.Background(), users.NewSearchUserBuilder(username))
		require.NoError(t, err)
		assert.NotNil(t, res)
		assert.Nil(t, res.PreviousPageCursor)
		assert.NotNil(t, res.NextPageCursor)
		assert.Len(t, res.Data, 10)

		user := res.Data[0]
		assert.Equal(t, uint64(1), user.ID)
		assert.Equal(t, username, user.Name)
	})

	// Test case: Search for a non-existent user
	t.Run("Search Non-existent User", func(t *testing.T) {
		username := InvalidUsername
		res, err := api.SearchUser(context.Background(), users.NewSearchUserBuilder(username))
		require.NoError(t, err)
		assert.NotNil(t, res)
		assert.Nil(t, res.PreviousPageCursor)
		assert.Nil(t, res.NextPageCursor)
		assert.Empty(t, res.Data)
	})
}
