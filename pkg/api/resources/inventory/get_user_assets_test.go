package inventory_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/inventory"
	"github.com/jaxron/roapi.go/pkg/api/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserAssets(t *testing.T) {
	// Create a new test resource
	api := inventory.New(utils.NewTestEnv())

	t.Run("Fetch User Assets", func(t *testing.T) {
		builder := inventory.NewGetUserAssetsBuilder(
			utils.SampleUserID1,
			types.ItemAssetTypeHat,
			types.ItemAssetTypeShirt,
		)
		result, err := api.GetUserAssets(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.Data)

		// Verify the structure of returned assets
		for _, asset := range result.Data {
			assert.NotZero(t, asset.AssetID)
			assert.NotEmpty(t, asset.Name)
			assert.NotEmpty(t, asset.AssetType)
			assert.NotZero(t, asset.Created)
		}
	})

	t.Run("Test Builder Methods", func(t *testing.T) {
		builder := inventory.NewGetUserAssetsBuilder(utils.SampleUserID1, types.ItemAssetTypeHat).
			WithLimit(25).
			WithSortOrder(types.SortOrderDesc).
			WithFilterDisapprovedAssets(true).
			WithShowApprovedOnly(true).
			WithCursor("someCursor")

		params := builder.Build()
		assert.Equal(t, utils.SampleUserID1, params.UserID)
		assert.Equal(t, []types.ItemAssetType{types.ItemAssetTypeHat}, params.AssetTypes)
		assert.Equal(t, int64(25), params.Limit)
		assert.Equal(t, types.SortOrderDesc, params.SortOrder)
		assert.True(t, params.FilterDisapprovedAssets)
		assert.True(t, params.ShowApprovedOnly)
		assert.Equal(t, "someCursor", params.Cursor)
	})

	t.Run("Invalid User ID", func(t *testing.T) {
		builder := inventory.NewGetUserAssetsBuilder(0, types.ItemAssetTypeHat)
		_, err := api.GetUserAssets(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "UserID")
	})

	t.Run("Invalid Limit", func(t *testing.T) {
		builder := inventory.NewGetUserAssetsBuilder(utils.SampleUserID1, types.ItemAssetTypeHat).
			WithLimit(30)
		_, err := api.GetUserAssets(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Limit")
	})

	t.Run("Invalid Cursor", func(t *testing.T) {
		builder := inventory.NewGetUserAssetsBuilder(utils.SampleUserID1, types.ItemAssetTypeHat).
			WithCursor("invalid-cursor")
		_, err := api.GetUserAssets(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Cursor")
	})
}
