package users_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserByID(t *testing.T) {
	// Create a new test resource
	api := users.New(utils.NewTestEnv())

	// Test case: Fetch information for a known user
	t.Run("Fetch Known User", func(t *testing.T) {
		user, err := api.GetUserByID(context.Background(), utils.SampleUserID1)
		require.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, utils.SampleUserID1, user.ID)
	})

	// Test case: Attempt to fetch information for a non-existent user
	t.Run("Fetch Non-existent User", func(t *testing.T) {
		user, err := api.GetUserByID(context.Background(), utils.InvalidUserID)
		require.Error(t, err)
		assert.Nil(t, user)
	})
}
