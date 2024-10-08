package friends_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/services/friends"
	"github.com/jaxron/roapi.go/pkg/api/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFollowings(t *testing.T) {
	// Create a new test service and validator
	api := friends.NewService(utils.NewTestEnv())

	// Test case: Fetch followings for a known user
	t.Run("Fetch Known User Followings", func(t *testing.T) {
		builder := friends.NewGetFollowingsBuilder(utils.SampleUserID1)
		followings, err := api.GetFollowings(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, followings)
		assert.NotEmpty(t, followings.Data)
	})

	// Test case: Attempt to fetch followings for a non-existent user
	t.Run("Fetch Non-existent User Followings", func(t *testing.T) {
		builder := friends.NewGetFollowingsBuilder(utils.InvalidUserID)
		followings, err := api.GetFollowings(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Nil(t, followings)
	})

	// Test case: Validate with invalid Limit
	t.Run("Invalid Limit", func(t *testing.T) {
		builder := friends.NewGetFollowingsBuilder(utils.SampleUserID1).WithLimit(23)
		_, err := api.GetFollowings(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Limit")
	})

	// Test case: Validate with invalid Cursor
	t.Run("Invalid Cursor", func(t *testing.T) {
		builder := friends.NewGetFollowingsBuilder(utils.SampleUserID1).WithCursor("invalidCursor")
		_, err := api.GetFollowings(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Cursor")
	})

	// Test case: Valid parameters with all fields set
	t.Run("Valid Parameters", func(t *testing.T) {
		builder := friends.NewGetFollowingsBuilder(utils.SampleUserID1).
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
