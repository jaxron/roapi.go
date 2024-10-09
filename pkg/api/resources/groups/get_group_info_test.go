package groups_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/groups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetGroupInfo(t *testing.T) {
	// Create a new test resource
	api := groups.New(utils.NewTestEnv())

	// Test case: Fetch group info for a known group
	t.Run("Fetch Known Group Info", func(t *testing.T) {
		groupInfo, err := api.GetGroupInfo(context.Background(), utils.SampleGroupID)
		require.NoError(t, err)
		assert.NotNil(t, groupInfo)
		assert.Equal(t, utils.SampleGroupID, groupInfo.ID)
		assert.NotEmpty(t, groupInfo.Name)
		assert.NotEmpty(t, groupInfo.Description)
		assert.NotZero(t, groupInfo.Owner.UserID)
		assert.NotEmpty(t, groupInfo.Owner.Username)
	})

	// Test case: Attempt to fetch group info for a non-existent group
	t.Run("Fetch Non-existent Group Info", func(t *testing.T) {
		groupInfo, err := api.GetGroupInfo(context.Background(), utils.InvalidGroupID)
		require.Error(t, err)
		assert.Nil(t, groupInfo)
	})
}
