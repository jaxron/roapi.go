package user_test

import (
	"context"
	"math"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/services/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	InvalidUserID   = uint64(math.MaxUint64)
	InvalidUsername = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
)

// NewTestService creates a new user.Service instance for testing purposes.
// It initializes the service with or without a cookie based on the useCookie parameter.
func NewTestService(useCookie bool) (*user.Service, error) {
	// Create a new test client
	client, err := utils.NewTestClient(true, useCookie)
	if err != nil {
		return nil, err
	}

	return user.NewService(client), nil
}

// TestGetUserInfo tests the GetUserInfo method of the user.Service.
func TestGetUserInfo(t *testing.T) {
	// Create a new test service
	api, err := NewTestService(false)
	require.NoError(t, err)

	// Test case: Fetch information for a known user
	t.Run("Fetch Known User", func(t *testing.T) {
		userID := uint64(1)
		user, err := api.GetUserInfo(context.Background(), 1)
		require.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, userID, user.ID)
		assert.Equal(t, "Roblox", user.Name)
	})

	// Test case: Attempt to fetch information for a non-existent user
	t.Run("Fetch Non-existent User", func(t *testing.T) {
		userID := InvalidUserID
		user, err := api.GetUserInfo(context.Background(), userID)
		require.Error(t, err)
		assert.Nil(t, user)
	})
}

// TestGetAuthUserInfo tests the GetAuthUserInfo method of the user.Service.
func TestGetAuthUserInfo(t *testing.T) {
	// Test case: Fetch authenticated user info
	t.Run("Fetch Authenticated User Info", func(t *testing.T) {
		// Create a new test service with authentication
		api, err := NewTestService(true)
		require.NoError(t, err)

		authUser, err := api.GetAuthUserInfo(context.Background())
		require.NoError(t, err)
		assert.NotNil(t, authUser)
		assert.NotZero(t, authUser.ID)
		assert.NotEmpty(t, authUser.Name)
	})

	// Test case: Attempt to fetch authenticated user info without a cookie
	t.Run("Fetch Authenticated User Info Without Cookie", func(t *testing.T) {
		// Create a new test service without authentication
		apiWithoutCookie, err := NewTestService(false)
		require.NoError(t, err)

		authUser, err := apiWithoutCookie.GetAuthUserInfo(context.Background())
		require.Error(t, err)
		assert.Nil(t, authUser)
		assert.Contains(t, err.Error(), "no .ROBLOSECURITY cookie found")
	})
}

// TestGetUsersByUsernames tests the GetUsersByUsernames method of the user.Service.
func TestGetUsersByUsernames(t *testing.T) {
	// Create a new test service
	api, err := NewTestService(false)
	require.NoError(t, err)

	t.Run("Fetch Known Users", func(t *testing.T) {
		usernames := []string{"Roblox", "builderman"}
		builder := user.NewUsersByUsernamesBuilder(usernames)
		result, err := api.GetUsersByUsernames(context.Background(), builder)

		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 2)

		// Check if the returned users match the requested usernames
		for _, user := range result.Data {
			assert.Contains(t, usernames, user.Name)
		}
	})

	t.Run("Fetch With Non-existent Username", func(t *testing.T) {
		usernames := []string{"Roblox", InvalidUsername}
		builder := user.NewUsersByUsernamesBuilder(usernames)
		result, err := api.GetUsersByUsernames(context.Background(), builder)

		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 1) // Only one user should be returned

		assert.Equal(t, "Roblox", result.Data[0].Name)
	})
}

// TestGetUsersByIDs tests the GetUsersByIDs method of the user.Service.
func TestGetUsersByIDs(t *testing.T) {
	// Create a new test service
	api, err := NewTestService(false)
	require.NoError(t, err)

	t.Run("Fetch Known Users", func(t *testing.T) {
		userIDs := []uint64{1, 156} // IDs for Roblox and Builderman
		builder := user.NewUsersByIDsBuilder(userIDs)
		result, err := api.GetUsersByIDs(context.Background(), builder)

		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 2)

		// Check if the returned users match the requested IDs
		for _, user := range result.Data {
			assert.Contains(t, userIDs, user.ID)
		}
	})

	t.Run("Fetch With Non-existent User ID", func(t *testing.T) {
		userIDs := []uint64{1, math.MaxUint64}
		builder := user.NewUsersByIDsBuilder(userIDs)
		result, err := api.GetUsersByIDs(context.Background(), builder)

		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 1) // Only one user should be returned

		assert.Equal(t, uint64(1), result.Data[0].ID)
	})
}

// TestGetUsernameHistory tests the GetUsernameHistory method of the user.Service.
func TestGetUsernameHistory(t *testing.T) {
	// Create a new test service
	api, err := NewTestService(false)
	require.NoError(t, err)

	// Test case: Fetch username history for a known user
	t.Run("Fetch Known User Username History", func(t *testing.T) {
		userID := uint64(1)
		history, err := api.GetUsernameHistory(context.Background(), user.NewUsernameHistoryBuilder(userID))
		require.NoError(t, err)
		assert.NotNil(t, history)
		assert.Nil(t, history.PreviousPageCursor)
		assert.Nil(t, history.NextPageCursor)
		assert.Empty(t, history.Data)
	})

	// Test case: Attempt to fetch username history for a non-existent user
	t.Run("Fetch Non-existent User Username History", func(t *testing.T) {
		userID := InvalidUserID
		history, err := api.GetUsernameHistory(context.Background(), user.NewUsernameHistoryBuilder(userID))
		require.Error(t, err)
		assert.Nil(t, history)
	})
}

// TestSearchByUsername tests the SearchUser method of the user.Service.
func TestSearchByUsername(t *testing.T) {
	// Create a new test service
	api, err := NewTestService(false)
	require.NoError(t, err)

	// Test case: Search for a known user
	t.Run("Search Known User", func(t *testing.T) {
		username := "Roblox"
		res, err := api.SearchUser(context.Background(), user.NewSearchUserBuilder(username))
		require.NoError(t, err)
		assert.NotNil(t, res)
		assert.Nil(t, res.PreviousPageCursor)
		assert.NotNil(t, res.NextPageCursor)
		assert.Len(t, res.Data, 10)

		user := res.Data[0]
		assert.Equal(t, uint64(1), user.ID)
		assert.Equal(t, username, user.Name)
	})

	// Test case: Search for a non-existent user
	t.Run("Search Non-existent User", func(t *testing.T) {
		username := InvalidUsername
		res, err := api.SearchUser(context.Background(), user.NewSearchUserBuilder(username))
		require.NoError(t, err)
		assert.NotNil(t, res)
		assert.Nil(t, res.PreviousPageCursor)
		assert.Nil(t, res.NextPageCursor)
		assert.Empty(t, res.Data)
	})
}
