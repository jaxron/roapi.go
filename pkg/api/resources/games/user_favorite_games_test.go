package games_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/games"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserFavoriteGames(t *testing.T) {
	// Create a new test resource
	api := games.New(utils.NewTestEnv())

	t.Run("Fetch User Favorite Games Successfully", func(t *testing.T) {
		builder := games.NewUserFavoriteGamesBuilder(utils.SampleUserID1)
		result, err := api.GetUserFavoriteGames(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotNil(t, result.Data)
	})

	t.Run("Fetch With Invalid User ID", func(t *testing.T) {
		builder := games.NewUserFavoriteGamesBuilder(0)
		_, err := api.GetUserFavoriteGames(context.Background(), builder.Build())
		require.Error(t, err)
	})

	t.Run("Test Builder Methods", func(t *testing.T) {
		builder := games.NewUserFavoriteGamesBuilder(utils.SampleUserID1).
			WithAccessFilter(games.AccessFilterPublic).
			WithLimit(25).
			WithCursor("nextPageCursor")

		params := builder.Build()
		assert.Equal(t, utils.SampleUserID1, params.UserID)
		assert.Equal(t, games.AccessFilterPublic, params.AccessFilter)
		assert.Equal(t, int64(25), params.Limit)
		assert.Equal(t, "nextPageCursor", params.Cursor)
	})

	t.Run("Invalid Limit", func(t *testing.T) {
		builder := games.NewUserFavoriteGamesBuilder(utils.SampleUserID1).
			WithLimit(30) // Invalid limit (not 10, 25, 50, or 100)
		_, err := api.GetUserFavoriteGames(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Limit")
	})
}
