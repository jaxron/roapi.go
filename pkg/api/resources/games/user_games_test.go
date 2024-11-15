package games_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/games"
	"github.com/jaxron/roapi.go/pkg/api/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserGames(t *testing.T) {
	// Create a new test resource
	api := games.New(utils.NewTestEnv())

	t.Run("Fetch User Games Successfully", func(t *testing.T) {
		builder := games.NewUserGamesBuilder(utils.SampleUserID1)
		result, err := api.GetUserGames(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotNil(t, result.Data)
	})

	t.Run("Fetch With Invalid User ID", func(t *testing.T) {
		builder := games.NewUserGamesBuilder(0)
		_, err := api.GetUserGames(context.Background(), builder.Build())
		require.Error(t, err)
	})

	t.Run("Test Builder Methods", func(t *testing.T) {
		builder := games.NewUserGamesBuilder(utils.SampleUserID1).
			WithAccessFilter(games.AccessFilterPublic).
			WithLimit(25).
			WithCursor("nextPageCursor").
			WithSortOrder(types.SortOrderDesc)

		params := builder.Build()
		assert.Equal(t, utils.SampleUserID1, params.UserID)
		assert.Equal(t, games.AccessFilterPublic, params.AccessFilter)
		assert.Equal(t, uint64(25), params.Limit)
		assert.Equal(t, "nextPageCursor", params.Cursor)
		assert.Equal(t, types.SortOrderDesc, params.SortOrder)
	})

	t.Run("Invalid Limit", func(t *testing.T) {
		builder := games.NewUserGamesBuilder(utils.SampleUserID1).
			WithLimit(30) // Invalid limit (not 10, 25, or 50)
		_, err := api.GetUserGames(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Limit")
	})

	t.Run("Invalid Cursor", func(t *testing.T) {
		builder := games.NewUserGamesBuilder(utils.SampleUserID1).
			WithCursor("invalid!cursor") // Invalid base64
		_, err := api.GetUserGames(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Cursor")
	})
}
