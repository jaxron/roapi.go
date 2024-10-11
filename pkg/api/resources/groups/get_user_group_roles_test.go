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
		userGroupRoles, err := api.GetUserGroupRoles(context.Background(), utils.SampleUserID1)
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
		userGroupRoles, err := api.GetUserGroupRoles(context.Background(), utils.InvalidUserID)
		require.Error(t, err)
		assert.Nil(t, userGroupRoles)
	})
}
