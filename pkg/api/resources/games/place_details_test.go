package games_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/games"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetMultiplePlaceDetails(t *testing.T) {
	// Create a new test resource
	api := games.New(utils.NewTestEnv())

	t.Run("Fetch Place Details Successfully", func(t *testing.T) {
		result, err := api.GetMultiplePlaceDetails(context.Background(), []uint64{utils.SampleGameID})
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result)
		assert.Equal(t, utils.SampleGameID, result[0].PlaceID)
	})

	t.Run("Fetch With Invalid Place ID", func(t *testing.T) {
		_, err := api.GetMultiplePlaceDetails(context.Background(), []uint64{utils.InvalidGameID})
		require.Error(t, err)
	})

	t.Run("Fetch With Empty Place IDs", func(t *testing.T) {
		_, err := api.GetMultiplePlaceDetails(context.Background(), []uint64{})
		require.Error(t, err)
	})
}
