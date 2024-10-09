package groups_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/services/groups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetGroupsInfo(t *testing.T) {
	// Create a new test service
	api := groups.NewService(utils.NewTestEnv())

	// Test case: Get info for known groups
	t.Run("Get Known Groups Info", func(t *testing.T) {
		builder := groups.NewGetGroupsInfoBuilder(utils.SampleGroupID, utils.SampleGroupID2)
		groupsInfo, err := api.GetGroupsInfo(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, groupsInfo)
		assert.Len(t, groupsInfo.Data, 2)

		// Check if groups are properly populated
		for _, group := range groupsInfo.Data {
			assert.NotZero(t, group.ID)
			assert.NotEmpty(t, group.Name)
			assert.NotZero(t, group.Owner.ID)
			assert.NotEmpty(t, group.Owner.Type)
		}
	})

	// Test case: Get info with empty group IDs
	t.Run("Get Info with Empty Group IDs", func(t *testing.T) {
		builder := groups.NewGetGroupsInfoBuilder()
		_, err := api.GetGroupsInfo(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "GroupIDs")
	})

	// Test case: Test builder methods
	t.Run("Test Builder Methods", func(t *testing.T) {
		builder := groups.NewGetGroupsInfoBuilder(1, 2, 3, 4).
			RemoveGroupIDs(3)

		params := builder.Build()
		assert.Len(t, params.GroupIDs, 3)
		assert.Contains(t, params.GroupIDs, "1")
		assert.Contains(t, params.GroupIDs, "2")
		assert.Contains(t, params.GroupIDs, "4")
		assert.NotContains(t, params.GroupIDs, "3")
	})
}
