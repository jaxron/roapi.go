package users_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/services/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetUsernameHistory tests the GetUsernameHistory method of the user.Service.
func TestGetUsernameHistory(t *testing.T) {
	// Create a new test service
	api := users.NewService(utils.NewTestClient())

	// Test case: Fetch username history for a known user
	t.Run("Fetch Known User Username History", func(t *testing.T) {
		userID := uint64(1)
		history, err := api.GetUsernameHistory(context.Background(), users.NewUsernameHistoryBuilder(userID))
		require.NoError(t, err)
		assert.NotNil(t, history)
		assert.Nil(t, history.PreviousPageCursor)
		assert.Nil(t, history.NextPageCursor)
		assert.Empty(t, history.Data)
	})

	// Test case: Attempt to fetch username history for a non-existent user
	t.Run("Fetch Non-existent User Username History", func(t *testing.T) {
		userID := InvalidUserID
		history, err := api.GetUsernameHistory(context.Background(), users.NewUsernameHistoryBuilder(userID))
		require.Error(t, err)
		assert.Nil(t, history)
	})
}
