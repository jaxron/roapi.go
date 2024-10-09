package groups_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/groups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetGroupRoles(t *testing.T) {
	// Create a new test resource
	api := groups.New(utils.NewTestEnv())

	// Test case: Fetch group roles for a known group
	t.Run("Fetch Known Group Roles", func(t *testing.T) {
		groupRoles, err := api.GetGroupRoles(context.Background(), utils.SampleGroupID)
		require.NoError(t, err)
		assert.NotNil(t, groupRoles)
		assert.Equal(t, utils.SampleGroupID, groupRoles.GroupID)
		assert.NotEmpty(t, groupRoles.Roles)

		// Check if roles are properly populated
		for _, role := range groupRoles.Roles {
			assert.NotZero(t, role.ID)
			assert.NotEmpty(t, role.Name)
		}
	})

	// Test case: Attempt to fetch group roles for a non-existent group
	t.Run("Fetch Non-existent Group Roles", func(t *testing.T) {
		groupRoles, err := api.GetGroupRoles(context.Background(), utils.InvalidGroupID)
		require.Error(t, err)
		assert.Nil(t, groupRoles)
	})
}
