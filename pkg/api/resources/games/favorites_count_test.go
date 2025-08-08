package games_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/games"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetGameFavoritesCount(t *testing.T) {
	// Create a new test resource
	api := games.New(utils.NewTestEnv())

	t.Run("Fetch Favorites Count Successfully", func(t *testing.T) {
		result, err := api.GetGameFavoritesCount(context.Background(), utils.SampleUniverseID)
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.GreaterOrEqual(t, result.FavoritesCount, int64(0))
	})

	t.Run("Fetch With Invalid Universe ID", func(t *testing.T) {
		_, err := api.GetGameFavoritesCount(context.Background(), utils.InvalidUniverseID)
		require.Error(t, err)
	})
}
