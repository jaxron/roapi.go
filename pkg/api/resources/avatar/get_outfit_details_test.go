package avatar_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/avatar"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetOutfitDetails(t *testing.T) {
	// Create a new test resource
	api := avatar.New(utils.NewTestEnv())

	// Test case: Fetch outfit details for a known outfit
	t.Run("Fetch Known Outfit Details", func(t *testing.T) {
		outfitDetails, err := api.GetOutfitDetails(context.Background(), utils.SampleOutfitID)
		require.NoError(t, err)
		assert.NotNil(t, outfitDetails)

		// Validate outfit details
		assert.NotZero(t, outfitDetails.ID)
		assert.NotEmpty(t, outfitDetails.Name)
		assert.NotEmpty(t, outfitDetails.Assets)
		assert.NotEmpty(t, outfitDetails.PlayerAvatarType)
		assert.NotEmpty(t, outfitDetails.OutfitType)

		// Validate body colors
		assert.NotEmpty(t, outfitDetails.BodyColors.HeadColor3)
		assert.NotEmpty(t, outfitDetails.BodyColors.TorsoColor3)
		assert.NotEmpty(t, outfitDetails.BodyColors.RightArmColor3)
		assert.NotEmpty(t, outfitDetails.BodyColors.LeftArmColor3)
		assert.NotEmpty(t, outfitDetails.BodyColors.RightLegColor3)
		assert.NotEmpty(t, outfitDetails.BodyColors.LeftLegColor3)

		// Validate scale
		assert.NotZero(t, outfitDetails.Scale.Height)
		assert.NotZero(t, outfitDetails.Scale.Width)
		assert.NotZero(t, outfitDetails.Scale.Head)
		assert.NotZero(t, outfitDetails.Scale.Depth)

		// Validate assets
		for _, asset := range outfitDetails.Assets {
			assert.NotZero(t, asset.ID)
			assert.NotEmpty(t, asset.Name)
			assert.NotZero(t, asset.AssetType.ID)
			assert.NotEmpty(t, asset.AssetType.Name)
			assert.NotZero(t, asset.CurrentVersionID)
		}
	})

	// Test case: Attempt to fetch outfit details with invalid outfit ID
	t.Run("Invalid Outfit ID", func(t *testing.T) {
		outfitDetails, err := api.GetOutfitDetails(context.Background(), utils.InvalidOutfitID)
		require.Error(t, err)
		assert.Nil(t, outfitDetails)
	})
}
