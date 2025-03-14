package games_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/games"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetGameServers(t *testing.T) {
	// Create a new test resource
	api := games.New(utils.NewTestEnv())

	t.Run("Fetch Game Servers Successfully", func(t *testing.T) {
		builder := games.NewGameServersBuilder(utils.SampleGameID2)
		result, err := api.GetGameServers(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotNil(t, result.Data)
		assert.Equal(t, 10, len(result.Data))
	})

	t.Run("Fetch With Invalid Place ID", func(t *testing.T) {
		builder := games.NewGameServersBuilder(utils.InvalidGameID)
		_, err := api.GetGameServers(context.Background(), builder.Build())
		require.Error(t, err)
	})

	t.Run("Test Builder Methods", func(t *testing.T) {
		builder := games.NewGameServersBuilder(utils.SampleGameID).
			WithServerType(games.ServerTypePrivate).
			WithSortOrder(games.SortOrderDesc).
			WithExcludeFullGames(true).
			WithLimit(50).
			WithCursor("nextPageCursor")

		params := builder.Build()
		assert.Equal(t, utils.SampleGameID, params.PlaceID)
		assert.Equal(t, games.ServerTypePrivate, params.ServerType)
		assert.Equal(t, games.SortOrderDesc, params.SortOrder)
		assert.True(t, params.ExcludeFullGames)
		assert.Equal(t, int32(50), params.Limit)
		assert.Equal(t, "nextPageCursor", params.Cursor)
	})

	t.Run("Invalid Limit", func(t *testing.T) {
		builder := games.NewGameServersBuilder(utils.SampleGameID).
			WithLimit(101) // Invalid limit (> 100)
		_, err := api.GetGameServers(context.Background(), builder.Build())
		require.Error(t, err)
	})
}
