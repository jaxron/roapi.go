package friends_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/friends"
	"github.com/jaxron/roapi.go/pkg/api/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFollowers(t *testing.T) {
	// Create a new test resource and validator
	api := friends.New(utils.NewTestEnv())

	// Test case: Fetch followers for a known user
	t.Run("Fetch Known User Followers", func(t *testing.T) {
		builder := friends.NewGetFollowersBuilder(utils.SampleUserID1)
		followers, err := api.GetFollowers(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, followers)
		assert.NotEmpty(t, followers.Data)
	})

	// Test case: Attempt to fetch followers for a non-existent user
	t.Run("Fetch Non-existent User Followers", func(t *testing.T) {
		builder := friends.NewGetFollowersBuilder(utils.InvalidUserID)
		followers, err := api.GetFollowers(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Nil(t, followers)
	})

	// Test case: Validate with invalid Limit
	t.Run("Invalid Limit", func(t *testing.T) {
		builder := friends.NewGetFollowersBuilder(utils.SampleUserID1).WithLimit(23)
		_, err := api.GetFollowers(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Limit")
	})

	// Test case: Validate with invalid Cursor
	t.Run("Invalid Cursor", func(t *testing.T) {
		builder := friends.NewGetFollowersBuilder(utils.SampleUserID1).WithCursor("invalidCursor")
		_, err := api.GetFollowers(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Cursor")
	})

	// Test case: Valid parameters with all fields set
	t.Run("Valid Parameters", func(t *testing.T) {
		builder := friends.NewGetFollowersBuilder(utils.SampleUserID1).
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
