package friends

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaxron/roapi.go/pkg/api/models"
	"github.com/jaxron/roapi.go/pkg/client"
)

// GetFriends fetches the friends of a user.
// GET https://friends.roblox.com/v1/users/{userID}/friends
func (s *Service) GetFriends(ctx context.Context, userID uint64) ([]models.UserResponse, error) {
	var friends struct {
		Data []models.UserResponse `json:"data"` // List of friend information
	}
	req := client.NewRequest().
		Method(http.MethodGet).
		URL(fmt.Sprintf("%s/v1/users/%d/friends", FriendsEndpoint, userID)).
		Result(&friends)

	resp, err := s.client.Do(ctx, req.Build())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return friends.Data, nil
}
