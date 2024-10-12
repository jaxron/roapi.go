package groups

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jaxron/axonet/pkg/client"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// ResourceInterface defines the interface for group-related operations.
type ResourceInterface interface {
	GetGroupInfo(ctx context.Context, groupID uint64) (*types.GroupResponse, error)
	GetGroupUsers(ctx context.Context, p GroupUsersParams) (*types.GroupUsersResponse, error)
	GetGroupRoles(ctx context.Context, groupID uint64) (*types.GroupRolesResponse, error)
	GetRoleUsers(ctx context.Context, p RoleUsersParams) (*types.RoleUsersResponse, error)
	SearchGroups(ctx context.Context, p SearchGroupsParams) (*types.SearchGroupsResponse, error)
	LookupGroup(ctx context.Context, groupName string) ([]types.GroupLookupResponse, error)
	GetGroupsInfo(ctx context.Context, p GetGroupsInfoParams) ([]types.GroupInfoResponse, error)
	GetUserGroupRoles(ctx context.Context, p UserGroupRolesParams) ([]types.UserGroupRolesResponse, error)
}

// Ensure Resource implements the ResourceInterface.
var _ ResourceInterface = (*Resource)(nil)

// Resource provides methods for interacting with group-related endpoints.
type Resource struct {
	client   *client.Client
	validate *validator.Validate
}

// New creates a new Resource with the specified version.
func New(client *client.Client, validate *validator.Validate) *Resource {
	return &Resource{
		client:   client,
		validate: validate,
	}
}
