package games_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/games"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUniverseIDFromPlace(t *testing.T) {
	// Create a new test resource
	api := games.New(utils.NewTestEnv())

	t.Run("Fetch Universe ID Successfully", func(t *testing.T) {
		result, err := api.GetUniverseIDFromPlace(context.Background(), utils.SampleGameID)
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, utils.SampleUniverseID, result.UniverseID)
	})

	t.Run("Fetch With Invalid Place ID", func(t *testing.T) {
		_, err := api.GetUniverseIDFromPlace(context.Background(), utils.InvalidGameID)
		require.Error(t, err)
	})
}
