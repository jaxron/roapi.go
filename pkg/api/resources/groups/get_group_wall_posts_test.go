package groups_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/groups"
	"github.com/jaxron/roapi.go/pkg/api/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetGroupWallPosts(t *testing.T) {
	// Create a new test resource
	api := groups.New(utils.NewTestEnv())

	// Test case: Fetch wall posts for a known group
	t.Run("Fetch Known Group Wall Posts", func(t *testing.T) {
		builder := groups.NewGroupWallPostsBuilder(utils.SampleGroupID3).WithLimit(100)
		wallPosts, err := api.GetGroupWallPosts(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, wallPosts)
		assert.NotEmpty(t, wallPosts.Data)

		// Check if wall posts are properly populated
		for _, post := range wallPosts.Data {
			assert.NotZero(t, post.ID)
			assert.NotEmpty(t, post.Body)
			if post.Poster != nil {
				assert.NotZero(t, post.Poster.User.UserID)
				assert.NotEmpty(t, post.Poster.User.Username)
				assert.NotZero(t, post.Poster.Role.ID)
				assert.NotEmpty(t, post.Poster.Role.Name)
			}
			assert.NotZero(t, post.Created)
			assert.NotZero(t, post.Updated)
		}
	})

	// Test case: Attempt to fetch wall posts for a non-existent group
	t.Run("Fetch Non-existent Group Wall Posts", func(t *testing.T) {
		builder := groups.NewGroupWallPostsBuilder(utils.InvalidGroupID)
		wallPosts, err := api.GetGroupWallPosts(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Nil(t, wallPosts)
	})

	// Test case: Validate with invalid Limit
	t.Run("Invalid Limit", func(t *testing.T) {
		builder := groups.NewGroupWallPostsBuilder(utils.SampleGroupID3).WithLimit(101)
		_, err := api.GetGroupWallPosts(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Limit")
	})

	// Test case: Valid parameters with all fields set
	t.Run("Valid Parameters", func(t *testing.T) {
		builder := groups.NewGroupWallPostsBuilder(utils.SampleGroupID3).
			WithLimit(50).
			WithCursor("someCursor").
			WithSortOrderDesc()

		params := builder.Build()
		assert.Equal(t, utils.SampleGroupID3, params.GroupID)
		assert.Equal(t, uint64(50), params.Limit)
		assert.Equal(t, "someCursor", params.Cursor)
		assert.Equal(t, types.SortOrderDesc, params.SortOrder)
	})

	// Test case: Test pagination
	t.Run("Test Pagination", func(t *testing.T) {
		builder := groups.NewGroupWallPostsBuilder(utils.SampleGroupID3).WithLimit(10)
		wallPosts, err := api.GetGroupWallPosts(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, wallPosts)

		if wallPosts.NextPageCursor != nil {
			// Fetch next page
			builder.WithCursor(*wallPosts.NextPageCursor)
			nextPage, err := api.GetGroupWallPosts(context.Background(), builder.Build())
			require.NoError(t, err)
			assert.NotNil(t, nextPage)
			assert.NotEmpty(t, nextPage.Data)
		}
	})
}
