package groups

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// SearchGroups searches for groups based on the provided parameters.
// GET https://groups.roblox.com/v1/groups/search
func (r *Resource) SearchGroups(ctx context.Context, p SearchGroupsParams) (*types.SearchGroupsResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var searchResults types.SearchGroupsResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(types.GroupsEndpoint+"/v1/groups/search").
		Query("keyword", p.Keyword).
		Query("prioritizeExactMatch", strconv.FormatBool(p.PrioritizeExactMatch)).
		Query("limit", strconv.FormatUint(p.Limit, 10)).
		Query("cursor", p.Cursor).
		Result(&searchResults).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	if err := r.validate.Struct(&searchResults); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &searchResults, nil
}

// SearchGroupsParams holds the parameters for searching groups.
type SearchGroupsParams struct {
	Keyword              string `json:"keyword"              validate:"required"`
	PrioritizeExactMatch bool   `json:"prioritizeExactMatch"`
	Limit                uint64 `json:"limit"                validate:"omitempty,oneof=10 25 50 100"`
	Cursor               string `json:"cursor"               validate:"omitempty,base64"`
}

// SearchGroupsBuilder is a builder for SearchGroupsParams.
type SearchGroupsBuilder struct {
	params SearchGroupsParams
}

// NewSearchGroupsBuilder creates a new SearchGroupsBuilder with default values.
func NewSearchGroupsBuilder(keyword string) *SearchGroupsBuilder {
	return &SearchGroupsBuilder{
		params: SearchGroupsParams{
			Keyword:              keyword,
			PrioritizeExactMatch: false,
			Limit:                10,
			Cursor:               "",
		},
	}
}

// WithPrioritizeExactMatch sets the prioritizeExactMatch flag.
func (b *SearchGroupsBuilder) WithPrioritizeExactMatch(prioritize bool) *SearchGroupsBuilder {
	b.params.PrioritizeExactMatch = prioritize
	return b
}

// WithLimit sets the limit.
func (b *SearchGroupsBuilder) WithLimit(limit uint64) *SearchGroupsBuilder {
	b.params.Limit = limit
	return b
}

// WithCursor sets the cursor.
func (b *SearchGroupsBuilder) WithCursor(cursor string) *SearchGroupsBuilder {
	b.params.Cursor = cursor
	return b
}

// Build returns the SearchGroupsParams.
func (b *SearchGroupsBuilder) Build() SearchGroupsParams {
	return b.params
}
