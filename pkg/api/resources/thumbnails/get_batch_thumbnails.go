package thumbnails

import (
	"context"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetBatchThumbnails fetches batch thumbnails.
// POST https://thumbnails.roblox.com/v1/batch
func (r *Resource) GetBatchThumbnails(ctx context.Context, p BatchThumbnailsParams) (*types.BatchThumbnailsResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, err
	}

	var batchThumbnails types.BatchThumbnailsResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodPost).
		URL(types.ThumbnailsEndpoint + "/v1/batch").
		MarshalBody(p.Requests).
		Result(&batchThumbnails).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return &batchThumbnails, nil
}

// BatchThumbnailsParams holds the parameters for getting batch thumbnails.
type BatchThumbnailsParams struct {
	Requests []types.ThumbnailRequest `json:"requests" validate:"required,min=1,max=100"`
}

// BatchThumbnailsBuilder is a builder for BatchThumbnailsParams.
type BatchThumbnailsBuilder struct {
	params BatchThumbnailsParams
}

// NewBatchThumbnailsBuilder creates a new BatchThumbnailsBuilder.
func NewBatchThumbnailsBuilder() *BatchThumbnailsBuilder {
	return &BatchThumbnailsBuilder{
		params: BatchThumbnailsParams{
			Requests: make([]types.ThumbnailRequest, 0),
		},
	}
}

// AddRequest adds a new ThumbnailRequest to the builder.
func (b *BatchThumbnailsBuilder) AddRequest(request types.ThumbnailRequest) *BatchThumbnailsBuilder {
	b.params.Requests = append(b.params.Requests, request)
	return b
}

// RemoveRequest removes a ThumbnailRequest from the builder based on RequestID.
func (b *BatchThumbnailsBuilder) RemoveRequest(requestID string) *BatchThumbnailsBuilder {
	for i, req := range b.params.Requests {
		if req.RequestID == requestID {
			b.params.Requests = append(b.params.Requests[:i], b.params.Requests[i+1:]...)
			break
		}
	}
	return b
}

// Build returns the BatchThumbnailsParams.
func (b *BatchThumbnailsBuilder) Build() BatchThumbnailsParams {
	return b.params
}
