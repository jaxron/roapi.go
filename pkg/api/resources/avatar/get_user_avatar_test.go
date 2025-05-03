package avatar_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/avatar"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserAvatar(t *testing.T) {
	// Create a new test resource
	api := avatar.New(utils.NewTestEnv())

	// Test case: Fetch avatar details for a known user
	t.Run("Fetch Known User Avatar", func(t *testing.T) {
		userAvatar, err := api.GetUserAvatar(context.Background(), utils.SampleUserID1)
		require.NoError(t, err)
		assert.NotNil(t, userAvatar)

		// Validate avatar details
		assert.NotNil(t, userAvatar.Scales)
		assert.NotEmpty(t, userAvatar.PlayerAvatarType)
		assert.NotNil(t, userAvatar.BodyColors)
		assert.NotNil(t, userAvatar.Assets)
		assert.NotNil(t, userAvatar.Emotes)

		// Validate scales
		assert.NotZero(t, userAvatar.Scales.Height)
		assert.NotZero(t, userAvatar.Scales.Width)
		assert.NotZero(t, userAvatar.Scales.Head)
		assert.NotZero(t, userAvatar.Scales.Depth)

		// Validate body colors
		assert.NotEmpty(t, userAvatar.BodyColors.HeadColor3)
		assert.NotEmpty(t, userAvatar.BodyColors.TorsoColor3)
		assert.NotEmpty(t, userAvatar.BodyColors.RightArmColor3)
		assert.NotEmpty(t, userAvatar.BodyColors.LeftArmColor3)
		assert.NotEmpty(t, userAvatar.BodyColors.RightLegColor3)
		assert.NotEmpty(t, userAvatar.BodyColors.LeftLegColor3)

		// Validate assets
		for _, asset := range userAvatar.Assets {
			assert.NotZero(t, asset.ID)
			assert.NotEmpty(t, asset.Name)
			assert.NotZero(t, asset.AssetType.ID)
			assert.NotEmpty(t, asset.AssetType.Name)
			assert.NotZero(t, asset.CurrentVersionID)
		}

		// Validate emotes
		for _, emote := range userAvatar.Emotes {
			assert.NotZero(t, emote.AssetID)
			assert.NotEmpty(t, emote.AssetName)
			assert.NotZero(t, emote.Position)
		}
	})

	// Test case: Attempt to fetch avatar details with invalid user ID
	t.Run("Invalid User ID", func(t *testing.T) {
		userAvatar, err := api.GetUserAvatar(context.Background(), utils.InvalidUserID)
		require.Error(t, err)
		assert.Nil(t, userAvatar)
	})
}
