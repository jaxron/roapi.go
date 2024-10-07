package users_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/services/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetAuthUserInfo tests the GetAuthUserInfo method of the user.Service.
func TestGetAuthUserInfo(t *testing.T) {
	// Create a new test service and validator
	api := users.NewService(utils.NewTestEnv())

	// Test case: Fetch authenticated user info
	t.Run("Fetch Authenticated User Info", func(t *testing.T) {
		authUser, err := api.GetAuthUserInfo(context.Background())
		require.NoError(t, err)
		assert.NotNil(t, authUser)
		assert.NotZero(t, authUser.ID)
		assert.NotEmpty(t, authUser.Name)
	})
}
