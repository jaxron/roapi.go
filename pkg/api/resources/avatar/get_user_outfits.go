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
func (r *Resource) GetUserOutfits(ctx context.Context, p UserOutfitsParams) ([]types.OutfitData, error) {
	if err := r.validate.Struct(p); err != nil {
		return nil, err
	}

	var userOutfits struct {
		Data []types.OutfitData `json:"data"`
	}
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

	return userOutfits.Data, nil
}

// UserOutfitsParams holds the parameters for getting user outfits.
type UserOutfitsParams struct {
	UserID          uint64 `json:"userId"          validate:"required"`
	IsEditable      bool   `json:"isEditable"`
	ItemsPerPage    int    `json:"itemsPerPage"    validate:"min=1"`
	OutfitType      string `json:"outfitType"`
	PaginationToken string `json:"paginationToken"`
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
			IsEditable:      false,
			ItemsPerPage:    50,
			OutfitType:      "Avatar",
			PaginationToken: "",
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

// Build returns the UserOutfitsParams.
func (b *UserOutfitsBuilder) Build() UserOutfitsParams {
	return b.params
}
