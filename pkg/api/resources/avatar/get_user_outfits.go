package avatar

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetUserOutfits fetches the outfits for a specific user.
// GET https://avatar.roblox.com/v2/avatar/users/{userId}/outfits
func (r *Resource) GetUserOutfits(ctx context.Context, p UserOutfitsParams) (*types.OutfitResponse, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidRequest, err)
	}

	var userOutfits types.OutfitResponse
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v2/avatar/users/%d/outfits", types.AvatarEndpoint, p.UserID)).
		Query("isEditable", strconv.FormatBool(p.IsEditable)).
		Query("itemsPerPage", strconv.Itoa(p.ItemsPerPage)).
		Query("outfitType", p.OutfitType).
		Query("paginationToken", p.PaginationToken).
		Result(&userOutfits).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	if err := r.validate.Struct(&userOutfits); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidResponse, err)
	}

	return &userOutfits, nil
}

// UserOutfitsParams holds the parameters for getting user outfits.
type UserOutfitsParams struct {
	UserID          uint64 `json:"userId"          validate:"required,gt=0"`
	IsEditable      bool   `json:"isEditable"`
	ItemsPerPage    int    `json:"itemsPerPage"    validate:"min=1"`
	OutfitType      string `json:"outfitType"`
	PaginationToken string `json:"paginationToken"`
	Page            uint32 `json:"page"            validate:"min=1"`
}

// UserOutfitsBuilder is a builder for UserOutfitsParams.
type UserOutfitsBuilder struct {
	params UserOutfitsParams
}

// NewUserOutfitsBuilder creates a new UserOutfitsBuilder with default values.
func NewUserOutfitsBuilder(userID uint64) *UserOutfitsBuilder {
	return &UserOutfitsBuilder{
		params: UserOutfitsParams{
			UserID:          userID,
			IsEditable:      true,
			ItemsPerPage:    50,
			OutfitType:      "Avatar",
			PaginationToken: "",
			Page:            1,
		},
	}
}

// WithIsEditable sets whether the outfits are editable.
func (b *UserOutfitsBuilder) WithIsEditable(isEditable bool) *UserOutfitsBuilder {
	b.params.IsEditable = isEditable
	return b
}

// WithItemsPerPage sets the number of items per page.
func (b *UserOutfitsBuilder) WithItemsPerPage(itemsPerPage int) *UserOutfitsBuilder {
	b.params.ItemsPerPage = itemsPerPage
	return b
}

// WithOutfitType sets the type of outfit.
func (b *UserOutfitsBuilder) WithOutfitType(outfitType string) *UserOutfitsBuilder {
	b.params.OutfitType = outfitType
	return b
}

// WithPaginationToken sets the pagination token.
func (b *UserOutfitsBuilder) WithPaginationToken(paginationToken string) *UserOutfitsBuilder {
	b.params.PaginationToken = paginationToken
	return b
}

// WithPage sets the page number.
func (b *UserOutfitsBuilder) WithPage(page uint32) *UserOutfitsBuilder {
	b.params.Page = page
	return b
}

// Build returns the UserOutfitsParams.
func (b *UserOutfitsBuilder) Build() UserOutfitsParams {
	return b.params
}
