package avatar_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/avatar"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserOutfits(t *testing.T) {
	// Create a new test resource
	api := avatar.New(utils.NewTestEnv())

	// Test case: Fetch user outfits for a known user
	t.Run("Fetch Known User Outfits", func(t *testing.T) {
		builder := avatar.NewUserOutfitsBuilder(utils.SampleUserID1)
		userOutfits, err := api.GetUserOutfits(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, userOutfits)
		assert.NotEmpty(t, userOutfits)

		// Check if outfits are properly populated
		for _, outfit := range userOutfits {
			assert.NotZero(t, outfit.ID)
			assert.NotEmpty(t, outfit.Name)
			assert.False(t, outfit.IsEditable)
			assert.Equal(t, "Avatar", outfit.OutfitType)
		}
	})

	// Test case: Attempt to fetch user outfits for a non-existent user
	t.Run("Fetch Non-existent User Outfits", func(t *testing.T) {
		builder := avatar.NewUserOutfitsBuilder(utils.InvalidUserID)
		userOutfits, err := api.GetUserOutfits(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Nil(t, userOutfits)
	})

	// Test case: Validate with invalid ItemsPerPage
	t.Run("Invalid ItemsPerPage", func(t *testing.T) {
		builder := avatar.NewUserOutfitsBuilder(utils.SampleUserID1).WithItemsPerPage(0)
		_, err := api.GetUserOutfits(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "ItemsPerPage")
	})

	// Test case: Valid parameters with all fields set
	t.Run("Valid Parameters", func(t *testing.T) {
		builder := avatar.NewUserOutfitsBuilder(utils.SampleUserID1).
			WithIsEditable(false).
			WithItemsPerPage(10).
			WithOutfitType("Avatar")

		params := builder.Build()
		assert.Equal(t, uint64(utils.SampleUserID1), params.UserID)
		assert.False(t, params.IsEditable)
		assert.Equal(t, 10, params.ItemsPerPage)
		assert.Equal(t, "Avatar", params.OutfitType)
	})
}
