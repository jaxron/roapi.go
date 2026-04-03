package catalog_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/catalog"
	"github.com/jaxron/roapi.go/pkg/api/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetItemDetails(t *testing.T) {
	// Create a new test resource
	api := catalog.New(utils.NewTestEnv())

	t.Run("Fetch Known Asset Details", func(t *testing.T) {
		builder := catalog.NewGetItemDetailsBuilder(
			catalog.CatalogItemRequest{ItemType: types.CatalogItemTypeAsset, ID: utils.SampleAssetID},
			catalog.CatalogItemRequest{ItemType: types.CatalogItemTypeAsset, ID: utils.SampleAssetID2},
		)
		result, err := api.GetItemDetails(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 2)

		// Validate common fields present on all items
		for _, item := range result.Data {
			assert.NotZero(t, item.ID)
			assert.NotEmpty(t, item.ItemType)
			assert.NotEmpty(t, item.Name)
			assert.NotEmpty(t, item.CreatorType)
			assert.NotZero(t, item.CreatorTargetID)
			assert.NotEmpty(t, item.CreatorName)
		}
	})

	t.Run("Fetch Single Asset", func(t *testing.T) {
		builder := catalog.NewGetItemDetailsBuilder(
			catalog.CatalogItemRequest{ItemType: types.CatalogItemTypeAsset, ID: utils.SampleAssetID},
		)
		result, err := api.GetItemDetails(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Data, 1)
		assert.Equal(t, utils.SampleAssetID, result.Data[0].ID)
	})

	t.Run("Invalid Asset ID", func(t *testing.T) {
		builder := catalog.NewGetItemDetailsBuilder(
			catalog.CatalogItemRequest{ItemType: types.CatalogItemTypeAsset, ID: utils.InvalidAssetID},
		)
		_, err := api.GetItemDetails(context.Background(), builder.Build())
		require.Error(t, err)
	})

	t.Run("Empty Items", func(t *testing.T) {
		builder := catalog.NewGetItemDetailsBuilder()
		_, err := api.GetItemDetails(context.Background(), builder.Build())
		require.Error(t, err)
	})

	t.Run("Test Builder Methods", func(t *testing.T) {
		item1 := catalog.CatalogItemRequest{ItemType: types.CatalogItemTypeAsset, ID: utils.SampleAssetID}
		item2 := catalog.CatalogItemRequest{ItemType: types.CatalogItemTypeAsset, ID: utils.SampleAssetID2}

		builder := catalog.NewGetItemDetailsBuilder(item1).
			WithItems(item2)

		params := builder.Build()
		assert.Len(t, params.Items, 2)

		builder.RemoveItems(utils.SampleAssetID)
		params = builder.Build()
		assert.Len(t, params.Items, 1)
		assert.Equal(t, utils.SampleAssetID2, params.Items[0].ID)
	})
}
