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
	"github.com/jaxron/roapi.go/internal/middleware/auth"
	"github.com/jaxron/roapi.go/internal/middleware/jsonheader"
)

var (
	ErrInvalidProxyFormat = errors.New("invalid proxy format")
	ErrProxiesFileNotSet  = errors.New("ROAPI_PROXIES_FILE environment variable not set")
	ErrCookiesFileNotSet  = errors.New("ROAPI_COOKIES_FILE environment variable not set")
)

const (
	// ExpectedProxyParts is the number of parts expected in a proxy string (IP:Port:Username:Password).
	ExpectedProxyParts = 4

	SampleUserID1   = uint64(7380156655)
	SampleUserID2   = uint64(7436054881)
	SampleUserID3   = uint64(7436059676)
	SampleUserID4   = uint64(1)   // Roblox
	SampleUserID5   = uint64(156) // Builderman
	SampleUsername1 = "actuallynotabot1"
	SampleUsername2 = "actuallynotabot2"
	SampleUsername3 = "actuallynotabot3"
	SampleUsername4 = "Roblox"
	SampleUsername5 = "builderman"
	InvalidUserID   = uint64(math.MaxUint64)
	InvalidUsername = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"

	SampleGroupID    = uint64(3336691)
	SampleGroupID2   = uint64(7)
	SampleGroupName  = "test"
	SampleRoleID     = uint64(23018355)
	InvalidGroupID   = uint64(math.MaxUint64)
	InvalidGroupName = "ThisGroupShouldNotExist12345"
	InvalidRoleID    = uint64(math.MaxUint64)
)

// NewTestEnv creates a new client.Client instance and a validator.Validate for testing purposes.
// It sets up the client with proxies and cookies based on environment variables.
func NewTestEnv(opts ...client.Option) (*client.Client, *validator.Validate) {
	// Use a basic logger for testing
	logger := logger.NewBasicLogger()

	// Get the proxies from environment variable
	proxies, err := getProxiesFromEnv(logger)
	if err != nil {
		panic(err)
	}

	// Get the cookies from environment variable
	cookies, err := getCookiesFromEnv(logger)
	if err != nil {
		panic(err)
	}

	// Create and return a new client with the specified options
	auth := auth.New(cookies)
	proxy := proxy.New(proxies)
	client := client.NewClient(
		append([]client.Option{
			client.WithLogger(logger),
			client.WithMiddleware(retry.New(1, 5000, 10000)),
			client.WithMiddleware(auth),
			client.WithMiddleware(proxy),
			client.WithMiddleware(jsonheader.New()),
		}, opts...)...,
	)

	// Shuffle the cookies and proxies
	auth.Shuffle()
	proxy.Shuffle()

	return client, validator.New(validator.WithRequiredStructEnabled())
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
	defer file.Close()

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
	defer file.Close()

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
