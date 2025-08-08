package utils

import (
	"bufio"
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
	ErrInvalidProxyFormat = errors.New("invalid proxy format")
	ErrProxiesFileNotSet  = errors.New("ROAPI_PROXIES_FILE environment variable not set")
	ErrCookiesFileNotSet  = errors.New("ROAPI_COOKIES_FILE environment variable not set")
)

const (
	// ExpectedProxyParts is the number of parts expected in a proxy string (IP:Port:Username:Password).
	ExpectedProxyParts = 4

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
)

// NewTestEnv creates a new client.Client instance and a validator.Validate for testing purposes.
// It sets up the client with proxies and cookies based on environment variables.
func NewTestEnv(opts ...client.Option) (*client.Client, *validator.Validate) {
	// Use a basic logger for testing
	basicLogger := logger.NewBasicLogger()

	// Get the proxies from environment variable
	proxies, err := getProxiesFromEnv(basicLogger)
	if err != nil {
		panic(err)
	}

	// Get the cookies from environment variable
	cookies, err := getCookiesFromEnv(basicLogger)
	if err != nil {
		panic(err)
	}

	// Create and return a new client with the specified options
	authMiddleware := auth.New(cookies)
	proxyMiddleware := proxy.New(proxies)
	httpClient := client.NewClient(
		append([]client.Option{
			client.WithLogger(basicLogger),
			client.WithMiddleware(retry.New(1, 5000, 10000)),
			client.WithMiddleware(proxyMiddleware),
			client.WithMiddleware(authMiddleware),
			client.WithMiddleware(jsonheader.New()),
		}, opts...)...,
	)

	// Shuffle the cookies and proxies
	authMiddleware.Shuffle()
	proxyMiddleware.Shuffle()

	return httpClient, validator.New(validator.WithRequiredStructEnabled())
}

// getProxiesFromEnv loads the proxies from the file specified in the ROAPI_PROXIES_FILE environment variable.
func getProxiesFromEnv(log logger.Logger) ([]*url.URL, error) {
	proxiesFile := os.Getenv("ROAPI_PROXIES_FILE")
	if proxiesFile == "" {
		return nil, ErrProxiesFileNotSet
	}

	proxies, err := readProxiesFromFile(proxiesFile)
	if err != nil {
		return nil, err
	}

	log.WithFields(logger.Int("count", len(proxies))).Debug("Loaded proxies")

	return proxies, nil
}

// readProxiesFromFile reads and parses proxies from the specified file.
// The file should contain one proxy per line in the format: IP:Port:Username:Password.
func readProxiesFromFile(fileName string) ([]*url.URL, error) {
	var proxies []*url.URL

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open proxy file: %w", err)
	}

	defer func() { _ = file.Close() }()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line into parts (IP:Port:Username:Password)
		parts := strings.Split(scanner.Text(), ":")
		if len(parts) != ExpectedProxyParts {
			return nil, fmt.Errorf("%w: %s", ErrInvalidProxyFormat, scanner.Text())
		}

		// Extract proxy components
		ip := parts[0]
		port := parts[1]
		username := parts[2]
		password := parts[3]

		// Construct the proxy URL
		proxyURL := fmt.Sprintf("http://%s:%s@%s", username, password, net.JoinHostPort(ip, port))

		// Parse the proxy URL
		parsedURL, err := url.Parse(proxyURL)
		if err != nil {
			return nil, fmt.Errorf("failed to parse proxy URL: %w", err)
		}

		// Add the proxy to the list
		proxies = append(proxies, parsedURL)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading proxy file: %w", err)
	}

	return proxies, nil
}

// getCookiesFromEnv loads the cookies from the file specified in the ROAPI_COOKIES_FILE environment variable.
func getCookiesFromEnv(log logger.Logger) ([]string, error) {
	cookie := os.Getenv("ROAPI_COOKIES_FILE")
	if cookie == "" {
		return nil, ErrCookiesFileNotSet
	}

	cookies, err := readCookiesFromFile(cookie)
	if err != nil {
		return nil, err
	}

	log.WithFields(logger.Int("count", len(cookies))).Debug("Loaded cookies")

	return cookies, nil
}

// readCookiesFromFile reads and parses cookies from the specified file.
// The file should contain one cookie per line.
func readCookiesFromFile(fileName string) ([]string, error) {
	var cookies []string

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open cookie file: %w", err)
	}

	defer func() { _ = file.Close() }()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cookies = append(cookies, scanner.Text())
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading cookie file: %w", err)
	}

	return cookies, nil
}
