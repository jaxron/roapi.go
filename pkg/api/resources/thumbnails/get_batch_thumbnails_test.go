package thumbnails_test

import (
	"context"
	"testing"

	"github.com/jaxron/roapi.go/internal/utils"
	"github.com/jaxron/roapi.go/pkg/api/resources/thumbnails"
	"github.com/jaxron/roapi.go/pkg/api/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBatchThumbnails(t *testing.T) {
	// Create a new test resource
	api := thumbnails.New(utils.NewTestEnv())

	// Test case: Fetch batch thumbnails for known targets
	t.Run("Fetch Known Batch Thumbnails", func(t *testing.T) {
		builder := thumbnails.NewBatchThumbnailsBuilder()
		builder.AddRequest(types.ThumbnailRequest{
			Type:       types.AvatarHeadShot,
			TargetID:   utils.SampleUserID1,
			Size:       types.Size420x420,
			Format:     types.PNG,
			IsCircular: false,
			RequestID:  "AvatarHeadShot:420x420:png:regular",
		})
		builder.AddRequest(types.ThumbnailRequest{
			Type:       types.GroupIcon,
			TargetID:   utils.SampleGroupID,
			Size:       types.Size150x150,
			Format:     types.PNG,
			IsCircular: true,
			RequestID:  "GroupIcon:150x150:png:circular",
		})

		batchThumbnails, err := api.GetBatchThumbnails(context.Background(), builder.Build())
		require.NoError(t, err)
		assert.NotNil(t, batchThumbnails)
		assert.Len(t, batchThumbnails, 2)

		// Check if thumbnails are properly populated
		for _, thumbnail := range batchThumbnails {
			assert.NotEmpty(t, thumbnail.RequestID)
			assert.NotZero(t, thumbnail.TargetID)
			assert.NotEmpty(t, thumbnail.State)
			assert.NotEmpty(t, thumbnail.ImageURL)
		}
	})

	// Test case: Fetch batch thumbnails with empty requests
	t.Run("Fetch Batch Thumbnails with Empty Requests", func(t *testing.T) {
		builder := thumbnails.NewBatchThumbnailsBuilder()
		_, err := api.GetBatchThumbnails(context.Background(), builder.Build())
		require.Error(t, err)
		assert.Contains(t, err.Error(), "Requests")
	})

	// Test case: Fetch batch thumbnails with invalid request
	t.Run("Fetch Batch Thumbnails with Invalid Request", func(t *testing.T) {
		builder := thumbnails.NewBatchThumbnailsBuilder()
		builder.AddRequest(types.ThumbnailRequest{
			Type:       "InvalidType",
			TargetID:   utils.SampleUserID1,
			Size:       "InvalidSize",
			Format:     "InvalidFormat",
			IsCircular: false,
			RequestID:  "InvalidRequest",
		})

		batchThumbnails, err := api.GetBatchThumbnails(context.Background(), builder.Build())
		require.NoError(t, err) // The API should still return a response
		assert.NotNil(t, batchThumbnails)
		assert.Len(t, batchThumbnails, 1)
		assert.NotZero(t, batchThumbnails[0].ErrorCode)
		assert.NotEmpty(t, batchThumbnails[0].ErrorMessage)
	})

	// Test case: Test builder methods
	t.Run("Test Builder Methods", func(t *testing.T) {
		builder := thumbnails.NewBatchThumbnailsBuilder()
		builder.AddRequest(types.ThumbnailRequest{
			Type:       types.AvatarHeadShot,
			TargetID:   1,
			Size:       types.Size420x420,
			Format:     types.PNG,
			IsCircular: false,
			RequestID:  "Request1",
		})
		builder.AddRequest(types.ThumbnailRequest{
			Type:       types.GroupIcon,
			TargetID:   2,
			Size:       types.Size150x150,
			Format:     types.PNG,
			IsCircular: true,
			RequestID:  "Request2",
		})
		builder.RemoveRequest("Request1")

		params := builder.Build()
		assert.Len(t, params.Requests, 1)
		assert.Equal(t, uint64(2), params.Requests[0].TargetID)
		assert.Equal(t, "Request2", params.Requests[0].RequestID)
	})
}
