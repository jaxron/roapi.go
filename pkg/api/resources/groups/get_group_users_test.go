package groups_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/groups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetGroupUsers(t *testing.T) {
	// Create a new test resource
	api := groups.New(utils.NewTestEnv())

	// Test case: Fetch group users for a known group
	t.Run("Fetch Known Group Users", func(t *testing.T) {
		builder := groups.NewGroupUsersBuilder(utils.SampleGroupID)
		groupUsers, err := api.GetGroupUsers(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, groupUsers)
		assert.NotEmpty(t, groupUsers.Data)
	})

	// Test case: Attempt to fetch group users for a non-existent group
	t.Run("Fetch Non-existent Group Users", func(t *testing.T) {
		builder := groups.NewGroupUsersBuilder(utils.InvalidGroupID)
		groupUsers, err := api.GetGroupUsers(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Nil(t, groupUsers)
	})

	// Test case: Validate with invalid Limit
	t.Run("Invalid Limit", func(t *testing.T) {
		builder := groups.NewGroupUsersBuilder(utils.SampleGroupID).WithLimit(101)
		_, err := api.GetGroupUsers(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Limit")
	})

	// Test case: Valid parameters with all fields set
	t.Run("Valid Parameters", func(t *testing.T) {
		builder := groups.NewGroupUsersBuilder(utils.SampleGroupID).
			WithLimit(50).
			WithCursor("someCursor").
			WithSortOrderDesc()

		params := builder.Build()
		assert.Equal(t, utils.SampleGroupID, params.GroupID)
		assert.Equal(t, uint64(50), params.Limit)
		assert.Equal(t, "someCursor", params.Cursor)
		assert.Equal(t, "Desc", params.SortOrder)
	})
}
