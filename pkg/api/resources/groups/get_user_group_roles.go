package groups

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/errors"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// GetUserGroupRoles fetches the group roles for a specific user.
// GET https://groups.roblox.com/v1/users/{userId}/groups/roles
func (r *Resource) GetUserGroupRoles(ctx context.Context, params UserGroupRolesParams) ([]types.UserGroupRolesResponse, error) {
	var userGroupRoles struct {
		Data []types.UserGroupRolesResponse `json:"data"`
	}
	resp, err := r.client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/groups/roles", types.GroupsEndpoint, params.UserID)).
		Query("includeLocked", strconv.FormatBool(params.IncludeLocked)).
		Query("includeNotificationPreferences", strconv.FormatBool(params.IncludeNotificationPreferences)).
		Result(&userGroupRoles).
		Do(ctx)
	if err != nil {
		return nil, errors.HandleAPIError(resp, err)
	}
	defer resp.Body.Close()

	return userGroupRoles.Data, nil
}

// UserGroupRolesParams holds the parameters for fetching user group roles.
type UserGroupRolesParams struct {
	UserID                         uint64 `json:"userId"`
	IncludeLocked                  bool   `json:"includeLocked"`
	IncludeNotificationPreferences bool   `json:"includeNotificationPreferences"`
}

// UserGroupRolesBuilder is a builder for UserGroupRolesParams.
type UserGroupRolesBuilder struct {
	params UserGroupRolesParams
}

// NewUserGroupRolesBuilder creates a new UserGroupRolesBuilder.
func NewUserGroupRolesBuilder(userID uint64) *UserGroupRolesBuilder {
	return &UserGroupRolesBuilder{
		params: UserGroupRolesParams{
			UserID:                         userID,
			IncludeLocked:                  false,
			IncludeNotificationPreferences: false,
		},
	}
}

// IncludeLocked sets the IncludeLocked parameter.
func (b *UserGroupRolesBuilder) IncludeLocked(includeLocked bool) *UserGroupRolesBuilder {
	b.params.IncludeLocked = includeLocked
	return b
}

// IncludeNotificationPreferences sets the IncludeNotificationPreferences parameter.
func (b *UserGroupRolesBuilder) IncludeNotificationPreferences(includeNotificationPreferences bool) *UserGroupRolesBuilder {
	b.params.IncludeNotificationPreferences = includeNotificationPreferences
	return b
}

// Build builds the UserGroupRolesParams.
func (b *UserGroupRolesBuilder) Build() UserGroupRolesParams {
	return b.params
}
