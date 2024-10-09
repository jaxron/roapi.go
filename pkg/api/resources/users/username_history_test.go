package users_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/users"
	"github.com/jaxron/roapi.go/pkg/api/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUsernameHistory(t *testing.T) {
	// Create a new test resource and validator
	api := users.New(utils.NewTestEnv())

	// Test case: Fetch username history for a known user
	t.Run("Fetch Known User Username History", func(t *testing.T) {
		builder := users.NewUsernameHistoryBuilder(utils.SampleUserID1)
		history, err := api.GetUsernameHistory(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, history)
		assert.Nil(t, history.PreviousPageCursor)
		assert.Nil(t, history.NextPageCursor)
		assert.Empty(t, history.Data)
	})

	// Test case: Attempt to fetch username history for a non-existent user
	t.Run("Fetch Non-existent User Username History", func(t *testing.T) {
		builder := users.NewUsernameHistoryBuilder(utils.InvalidUserID)
		history, err := api.GetUsernameHistory(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Nil(t, history)
	})

	// Test case: Validate with invalid Limit
	t.Run("Invalid Limit", func(t *testing.T) {
		builder := users.NewUsernameHistoryBuilder(utils.SampleUserID1).WithLimit(101)
		_, err := api.GetUsernameHistory(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Limit")
	})

	// Test case: Validate with invalid Cursor
	t.Run("Invalid Cursor", func(t *testing.T) {
		builder := users.NewUsernameHistoryBuilder(utils.SampleUserID1).WithCursor("invalidCursor")
		_, err := api.GetUsernameHistory(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Cursor")
	})

	// Test case: Valid parameters with all fields set
	t.Run("Valid Parameters", func(t *testing.T) {
		builder := users.NewUsernameHistoryBuilder(utils.SampleUserID1).
			WithLimit(50).
			WithCursor("someCursor").
			WithSortOrderDesc()

		params := builder.Build()
		assert.Equal(t, uint64(utils.SampleUserID1), params.UserID)
		assert.Equal(t, uint64(50), params.Limit)
		assert.Equal(t, "someCursor", params.Cursor)
		assert.Equal(t, types.SortOrderDesc, params.SortOrder)
	})
}
