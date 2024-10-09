package groups_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/services/groups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSearchGroups(t *testing.T) {
	// Create a new test service
	api := groups.NewService(utils.NewTestEnv())

	// Test case: Search for groups with a known keyword
	t.Run("Search Known Groups", func(t *testing.T) {
		builder := groups.NewSearchGroupsBuilder("test")
		searchResults, err := api.SearchGroups(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, searchResults)
		assert.NotEmpty(t, searchResults.Data)

		// Check if groups are properly populated
		for _, group := range searchResults.Data {
			assert.NotZero(t, group.ID)
			assert.NotEmpty(t, group.Name)
			assert.NotZero(t, group.MemberCount)
		}
	})

	// Test case: Search with empty keyword
	t.Run("Search with Empty Keyword", func(t *testing.T) {
		builder := groups.NewSearchGroupsBuilder("")
		_, err := api.SearchGroups(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Keyword")
	})

	// Test case: Validate with invalid Limit
	t.Run("Invalid Limit", func(t *testing.T) {
		builder := groups.NewSearchGroupsBuilder(utils.SampleGroupName).WithLimit(101)
		_, err := api.SearchGroups(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Limit")
	})

	// Test case: Valid parameters with all fields set
	t.Run("Valid Parameters", func(t *testing.T) {
		builder := groups.NewSearchGroupsBuilder(utils.SampleGroupName).
			WithPrioritizeExactMatch(true).
			WithLimit(50).
			WithCursor("someCursor")

		params := builder.Build()
		assert.Equal(t, utils.SampleGroupName, params.Keyword)
		assert.True(t, params.PrioritizeExactMatch)
		assert.Equal(t, uint64(50), params.Limit)
		assert.Equal(t, "someCursor", params.Cursor)
	})
}
