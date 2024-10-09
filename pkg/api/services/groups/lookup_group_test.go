package groups_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/services/groups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLookupGroup(t *testing.T) {
	// Create a new test service
	api := groups.NewService(utils.NewTestEnv())

	// Test case: Lookup groups with a known name
	t.Run("Lookup Known Groups", func(t *testing.T) {
		lookupResults, err := api.LookupGroup(context.Background(), utils.SampleGroupName)
		require.NoError(t, err)
		assert.NotNil(t, lookupResults)
		assert.NotEmpty(t, lookupResults.Data)

		// Check if groups are properly populated
		for _, group := range lookupResults.Data {
			assert.NotZero(t, group.ID)
			assert.NotEmpty(t, group.Name)
			assert.NotZero(t, group.MemberCount)
		}
	})

	// Test case: Lookup with empty group name
	t.Run("Lookup with Empty Group Name", func(t *testing.T) {
		_, err := api.LookupGroup(context.Background(), "")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Name")
	})

	// Test case: Lookup with non-existent group name
	t.Run("Lookup Non-existent Group", func(t *testing.T) {
		lookupResults, err := api.LookupGroup(context.Background(), utils.InvalidGroupName)
		require.NoError(t, err)
		assert.NotNil(t, lookupResults)
		assert.Empty(t, lookupResults.Data)
	})
}
