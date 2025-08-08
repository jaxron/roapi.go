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

func TestGetRoleUsers(t *testing.T) {
	// Create a new test resource
	api := groups.New(utils.NewTestEnv())

	// Test case: Fetch role users for a known group and role
	t.Run("Fetch Known Role Users", func(t *testing.T) {
		builder := groups.NewRoleUsersBuilder(utils.SampleGroupID, utils.SampleRoleID)
		roleUsers, err := api.GetRoleUsers(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, roleUsers)
		assert.NotEmpty(t, roleUsers.Data)

		// Check if users are properly populated
		for _, user := range roleUsers.Data {
			assert.NotZero(t, user.UserID)
			assert.NotEmpty(t, user.Username)
			assert.NotEmpty(t, user.DisplayName)
		}
	})

	// Test case: Attempt to fetch role users for a non-existent group
	t.Run("Fetch Non-existent Group Role Users", func(t *testing.T) {
		builder := groups.NewRoleUsersBuilder(utils.InvalidGroupID, utils.SampleRoleID)
		roleUsers, err := api.GetRoleUsers(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Nil(t, roleUsers)
	})

	// Test case: Validate with invalid Limit
	t.Run("Invalid Limit", func(t *testing.T) {
		builder := groups.NewRoleUsersBuilder(utils.SampleGroupID, utils.SampleRoleID).WithLimit(101)
		_, err := api.GetRoleUsers(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Limit")
	})

	// Test case: Valid parameters with all fields set
	t.Run("Valid Parameters", func(t *testing.T) {
		builder := groups.NewRoleUsersBuilder(utils.SampleGroupID, utils.SampleRoleID).
			WithLimit(50).
			WithCursor("someCursor").
			WithSortOrderDesc()

		params := builder.Build()
		assert.Equal(t, utils.SampleGroupID, params.GroupID)
		assert.Equal(t, utils.SampleRoleID, params.RoleID)
		assert.Equal(t, int64(50), params.Limit)
		assert.Equal(t, "someCursor", params.Cursor)
		assert.Equal(t, types.SortOrderDesc, params.SortOrder)
	})
}
