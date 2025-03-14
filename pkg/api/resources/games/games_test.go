package games_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/games"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetGamesByUniverseIDs(t *testing.T) {
	// Create a new test resource
	api := games.New(utils.NewTestEnv())

	t.Run("Fetch Games Successfully", func(t *testing.T) {
		result, err := api.GetGamesByUniverseIDs(context.Background(), []uint64{utils.SampleUniverseID})
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.Data)
		assert.Equal(t, utils.SampleUniverseID, result.Data[0].ID)
	})

	t.Run("Fetch With Invalid Universe ID", func(t *testing.T) {
		_, err := api.GetGamesByUniverseIDs(context.Background(), []uint64{utils.InvalidUniverseID})
		require.Error(t, err)
	})

	t.Run("Fetch With Empty Universe IDs", func(t *testing.T) {
		_, err := api.GetGamesByUniverseIDs(context.Background(), []uint64{})
		require.Error(t, err)
	})
}
