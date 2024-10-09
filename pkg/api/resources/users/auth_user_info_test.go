package users_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAuthUserInfo(t *testing.T) {
	// Create a new test resource and validator
	api := users.New(utils.NewTestEnv())

	// Test case: Fetch authenticated user info
	t.Run("Fetch Authenticated User Info", func(t *testing.T) {
		authUser, err := api.GetAuthUserInfo(context.Background())
		require.NoError(t, err)
		assert.NotNil(t, authUser)
		assert.NotZero(t, authUser.ID)
		assert.NotEmpty(t, authUser.Name)
	})
}
