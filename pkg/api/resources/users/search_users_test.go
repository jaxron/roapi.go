package users_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSearchUsers(t *testing.T) {
	// Create a new test resource and validator
	api := users.New(utils.NewTestEnv())

	// Test case: Search for a known user
	t.Run("Search Known User", func(t *testing.T) {
		builder := users.NewSearchUsersBuilder(utils.SampleUsername4)
		res, err := api.SearchUsers(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, res)
		assert.Nil(t, res.PreviousPageCursor)
		assert.NotNil(t, res.NextPageCursor)
		assert.Len(t, res.Data, 10)

		user := res.Data[0]
		assert.Equal(t, int64(1), user.ID)
		assert.Equal(t, utils.SampleUsername4, user.Name)
	})

	// Test case: Search for a non-existent user
	t.Run("Search Non-existent User", func(t *testing.T) {
		username := utils.InvalidUsername
		builder := users.NewSearchUsersBuilder(username)

		res, err := api.SearchUsers(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, res)
		assert.Nil(t, res.PreviousPageCursor)
		assert.Nil(t, res.NextPageCursor)
		assert.Empty(t, res.Data)
	})

	// Test case: Validate with invalid Limit
	t.Run("Invalid Limit", func(t *testing.T) {
		builder := users.NewSearchUsersBuilder("TestUser").WithLimit(12)
		_, err := api.SearchUsers(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Limit")
	})

	// Test case: Validate with invalid Cursor
	t.Run("Invalid Cursor", func(t *testing.T) {
		builder := users.NewSearchUsersBuilder("TestUser").WithCursor("invalidCursor")
		_, err := api.SearchUsers(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Cursor")
	})

	// Test case: Valid parameters with all fields set
	t.Run("Valid Parameters", func(t *testing.T) {
		builder := users.NewSearchUsersBuilder("TestUser").
			WithLimit(50).
			WithCursor("someCursor")

		params := builder.Build()
		assert.Equal(t, "TestUser", params.Username)
		assert.Equal(t, int64(50), params.Limit)
		assert.Equal(t, "someCursor", params.Cursor)
	})
}
