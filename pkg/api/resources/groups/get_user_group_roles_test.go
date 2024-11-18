package groups_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/groups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserGroupRoles(t *testing.T) {
	// Create a new test resource
	api := groups.New(utils.NewTestEnv())

	// Test case: Fetch user group roles for a known user
	t.Run("Fetch Known User Group Roles", func(t *testing.T) {
		builder := groups.NewUserGroupRolesBuilder(utils.SampleUserID1)
		userGroupRoles, err := api.GetUserGroupRoles(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, userGroupRoles)
		assert.NotEmpty(t, userGroupRoles.Data)

		// Check if user group roles are properly populated
		for _, userGroup := range userGroupRoles.Data {
			assert.NotZero(t, userGroup.Group.ID)
			assert.NotEmpty(t, userGroup.Group.Name)
			assert.NotZero(t, userGroup.Role.ID)
			assert.NotEmpty(t, userGroup.Role.Name)
		}
	})

	// Test case: Attempt to fetch user group roles for a non-existent user
	t.Run("Fetch Non-existent User Group Roles", func(t *testing.T) {
		builder := groups.NewUserGroupRolesBuilder(utils.InvalidUserID)
		userGroupRoles, err := api.GetUserGroupRoles(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Nil(t, userGroupRoles)
	})

	// Test case: Valid parameters with all fields set
	t.Run("Valid Parameters", func(t *testing.T) {
		builder := groups.NewUserGroupRolesBuilder(utils.SampleUserID1).
			IncludeLocked(true).
			IncludeNotificationPreferences(true)

		params := builder.Build()
		assert.Equal(t, utils.SampleUserID1, params.UserID)
		assert.True(t, params.IncludeLocked)
		assert.True(t, params.IncludeNotificationPreferences)
	})
}
