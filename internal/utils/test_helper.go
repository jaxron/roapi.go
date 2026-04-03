package utils

import (
	"errors"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jaxron/axonet/middleware/proxy"
	"github.com/jaxron/axonet/middleware/retry"
	"github.com/jaxron/axonet/pkg/client"
	"github.com/jaxron/axonet/pkg/client/logger"
	"github.com/jaxron/roapi.go/pkg/api/middleware/auth"
	"github.com/jaxron/roapi.go/pkg/api/middleware/jsonheader"
)

var (
	ErrInvalidProxyFormat = errors.New("invalid proxy format, expected IP:Port:Username:Password")
	ErrProxyNotSet        = errors.New("ROAPI_PROXY environment variable not set")
	ErrCookieNotSet       = errors.New("ROAPI_COOKIE environment variable not set")
)

// ExpectedProxyParts is the number of parts expected in a proxy string (IP:Port:Username:Password).
const ExpectedProxyParts = 4

const (
	SampleUserID1   = int64(7380156655)
	SampleUserID2   = int64(7436054881)
	SampleUserID3   = int64(7436059676)
	SampleUserID4   = int64(1)   // Roblox
	SampleUserID5   = int64(156) // Builderman
	SampleUsername1 = "actuallynotabot1"
	SampleUsername2 = "actuallynotabot2"
	SampleUsername3 = "actuallynotabot3"
	SampleUsername4 = "Roblox"
	SampleUsername5 = "builderman"
	InvalidUserID   = int64(0)
	InvalidUsername = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"

	SampleGroupID    = int64(3336691)
	SampleGroupID2   = int64(7)
	SampleGroupID3   = int64(6303746)
	SampleGroupName  = "test"
	SampleRoleID     = int64(23018355)
	InvalidGroupID   = int64(math.MaxInt64)
	InvalidGroupName = "ThisGroupShouldNotExist12345"
	InvalidRoleID    = int64(math.MaxInt64)

	SampleUniverseID  = int64(6591173970)
	SampleGameID      = int64(104971911222178)
	SampleGameID2     = int64(116495829188952)
	InvalidUniverseID = int64(0)
	InvalidGameID     = int64(0)

	SampleOutfitID  = int64(13993719293)
	InvalidOutfitID = int64(math.MaxInt64)

	SampleAssetID  = int64(3360686498)
	SampleAssetID2 = int64(48474356)
	InvalidAssetID = int64(0)
)

// NewTestEnv creates a new client.Client instance and a validator.Validate for testing purposes.
// It reads proxy and cookie values directly from environment variables.
func NewTestEnv(opts ...client.Option) (*client.Client, *validator.Validate) {
	basicLogger := logger.NewBasicLogger()

	proxyURL, err := parseProxy(os.Getenv("ROAPI_PROXY"))
	if err != nil {
		panic(err)
	}

	cookie := os.Getenv("ROAPI_COOKIE")
	if cookie == "" {
		panic(ErrCookieNotSet)
	}

	authMiddleware := auth.New([]string{cookie})
	proxyMiddleware := proxy.New([]*url.URL{proxyURL})
	httpClient := client.NewClient(
		append([]client.Option{
			client.WithLogger(basicLogger),
			client.WithMiddleware(retry.New(1, 5000, 10000)),
			client.WithMiddleware(proxyMiddleware),
			client.WithMiddleware(authMiddleware),
			client.WithMiddleware(jsonheader.New()),
		}, opts...)...,
	)

	return httpClient, validator.New(validator.WithRequiredStructEnabled())
}

// parseProxy parses a proxy string in the format IP:Port:Username:Password into a URL.
func parseProxy(raw string) (*url.URL, error) {
	if raw == "" {
		return nil, ErrProxyNotSet
	}

	parts := strings.Split(raw, ":")
	if len(parts) != ExpectedProxyParts {
		return nil, fmt.Errorf("%w: %s", ErrInvalidProxyFormat, raw)
	}

	proxyURL := fmt.Sprintf("http://%s:%s@%s", parts[2], parts[3], net.JoinHostPort(parts[0], parts[1]))

	parsedURL, err := url.Parse(proxyURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse proxy URL: %w", err)
	}

	return parsedURL, nil
}
