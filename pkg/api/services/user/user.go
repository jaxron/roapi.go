package user

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jaxron/roapi.go/pkg/api/models"
	"github.com/jaxron/roapi.go/pkg/client"
	apierrors "github.com/jaxron/roapi.go/pkg/errors"
)

const UsersEndpoint = "https://users.roblox.com"

// ServiceInterface defines the interface for user-related operations.
type ServiceInterface interface {
	GetUserInfo(ctx context.Context, userID uint64) (*models.UserInfo, error)
	GetAuthUserInfo(ctx context.Context) (*models.AuthUserInfo, error)
	GetUsersByUsernames(ctx context.Context, b *UsersByUsernamesBuilder) (*models.UsersByUsernames, error)
	GetUsersByIDs(ctx context.Context, b *UsersByIDsBuilder) (*models.UsersByIDs, error)
	GetUsernameHistory(ctx context.Context, b *UsernameHistoryBuilder) (*models.UsernameHistory, error)
	SearchUser(ctx context.Context, b *SearchUserBuilder) (*models.SearchResult, error)
}

// Ensure Service implements the ServiceInterface.
var _ ServiceInterface = (*Service)(nil)

// Service provides methods for interacting with user-related endpoints.
type Service struct {
	Client *client.Client
}

// NewService creates a new Service with the specified version.
func NewService(client *client.Client) *Service {
	return &Service{
		Client: client,
	}
}

// GetUserInfo fetches information for a user with the given ID.
// GET https://users.roblox.com/v1/users/{userID}
func (s *Service) GetUserInfo(ctx context.Context, userID uint64) (*models.UserInfo, error) {
	var user models.UserInfo
	req := client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d", UsersEndpoint, userID)).
		Result(&user)

	resp, err := s.Client.Do(ctx, req.Build())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user info: %w", err)
	}
	defer resp.Body.Close()

	return &user, nil
}

// GetAuthUserInfo fetches information for the authenticated user.
// GET https://users.roblox.com/v1/users/authenticated
func (s *Service) GetAuthUserInfo(ctx context.Context) (*models.AuthUserInfo, error) {
	if s.Client.Handler.Auth.GetCookieCount() == 0 {
		return nil, apierrors.ErrNoCookie
	}

	var user models.AuthUserInfo
	req := client.NewRequest().
		Method(http.MethodGet).
		URL(UsersEndpoint + "/v1/users/authenticated").
		Result(&user).
		UseCookie(true)

	resp, err := s.Client.Do(ctx, req.Build())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch authenticated user info: %w", err)
	}
	defer resp.Body.Close()

	return &user, nil
}

// GetUsersByUsernames fetches information for users with the given usernames.
// POST https://users.roblox.com/v1/usernames/users
func (s *Service) GetUsersByUsernames(ctx context.Context, b *UsersByUsernamesBuilder) (*models.UsersByUsernames, error) {
	var users models.UsersByUsernames
	req, err := client.NewRequest().
		Method(http.MethodPost).
		URL(UsersEndpoint + "/v1/usernames/users").
		Result(&users).
		JSONBody(b.MarshalJSON)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req.Build())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &users, nil
}

// GetUsersByIDs fetches information for users with the given IDs.
// POST https://users.roblox.com/v1/users
func (s *Service) GetUsersByIDs(ctx context.Context, b *UsersByIDsBuilder) (*models.UsersByIDs, error) {
	var users models.UsersByIDs
	req, err := client.NewRequest().
		Method(http.MethodPost).
		URL(UsersEndpoint + "/v1/users").
		Result(&users).
		JSONBody(b.MarshalJSON)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(ctx, req.Build())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &users, nil
}

// GetUsernameHistory fetches the username history for a user.
// GET https://users.roblox.com/v1/users/{userID}/username-history
func (s *Service) GetUsernameHistory(ctx context.Context, b *UsernameHistoryBuilder) (*models.UsernameHistory, error) {
	var history models.UsernameHistory
	req := client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/username-history", UsersEndpoint, b.userID)).
		Query("limit", strconv.FormatUint(b.limit, 10)).
		Query("sortOrder", b.sortOrder).
		Query("cursor", b.cursor).
		Result(&history)

	resp, err := s.Client.Do(ctx, req.Build())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch username history: %w", err)
	}
	defer resp.Body.Close()

	return &history, nil
}

// SearchUser searches for a user with the given username.
// GET https://users.roblox.com/v1/users/search
func (s *Service) SearchUser(ctx context.Context, b *SearchUserBuilder) (*models.SearchResult, error) {
	var result models.SearchResult
	req := client.NewRequest().
		Method(http.MethodGet).
		URL(UsersEndpoint+"/v1/users/search").
		Query("keyword", b.username).
		Query("limit", strconv.FormatUint(b.limit, 10)).
		Query("cursor", b.cursor).
		Result(&result)

	resp, err := s.Client.Do(ctx, req.Build())
	if err != nil {
		return nil, fmt.Errorf("failed to search for user: %w", err)
	}
	defer resp.Body.Close()

	return &result, nil
}
