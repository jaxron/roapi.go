package groups

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jaxron/axonet/pkg/client"
	"github.com/jaxron/roapi.go/pkg/api/types"
)

// ServiceInterface defines the interface for group-related operations.
type ServiceInterface interface {
	GetGroupInfo(ctx context.Context, groupID uint64) (*types.GroupResponse, error)
	GetGroupUsers(ctx context.Context, p GroupUsersParams) (*types.GroupUsersResponse, error)
	GetGroupRoles(ctx context.Context, groupID uint64) (*types.GroupRolesResponse, error)
	GetRoleUsers(ctx context.Context, p RoleUsersParams) (*types.RoleUsersResponse, error)
	SearchGroups(ctx context.Context, p SearchGroupsParams) (*types.SearchGroupsResponse, error)
	LookupGroup(ctx context.Context, groupName string) (*types.GroupLookupResponse, error)
	GetGroupsInfo(ctx context.Context, p GetGroupsInfoParams) (*types.GroupsInfoResponse, error)
}

// Ensure Service implements the ServiceInterface.
var _ ServiceInterface = (*Service)(nil)

// Service provides methods for interacting with group-related endpoints.
type Service struct {
	client   *client.Client
	validate *validator.Validate
}

// NewService creates a new Service with the specified version.
func NewService(client *client.Client, validate *validator.Validate) *Service {
	return &Service{
		client:   client,
		validate: validate,
	}
}
