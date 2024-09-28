package utils

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"

	"github.com/jaxron/roapi.go/internal/handler"
	"github.com/jaxron/roapi.go/pkg/client"
	"github.com/jaxron/roapi.go/pkg/logger"
	"go.uber.org/zap"
)

var (
	ErrInvalidProxyFormat = errors.New("invalid proxy format")
	ErrProxiesFileNotSet  = errors.New("ROAPI_PROXIES_FILE environment variable not set")
	ErrCookiesFileNotSet  = errors.New("ROAPI_COOKIES_FILE environment variable not set")
)

const (
	// ExpectedProxyParts is the number of parts expected in a proxy string (IP:Port:Username:Password).
	ExpectedProxyParts = 4
)

// NewTestClient creates a new client.Client instance for testing purposes.
// It sets up the client with proxies and cookies based on environment variables.
func NewTestClient(useProxies bool, useCookie bool, opts ...client.Option) (*client.Client, error) {
	// Create a development logger
	logger := logger.NewDevelopmentLogger()

	// Get the proxies from environment variable
	proxies, err := getProxiesFromEnv(logger, useProxies)
	if err != nil {
		return nil, err
	}

	// Get the cookies from environment variable
	cookies, err := getCookiesFromEnv(logger, useCookie)
	if err != nil {
		return nil, err
	}

	// Create and return a new client with the specified options
	return client.NewClient(
		append([]client.Option{
			client.WithProxies(proxies),
			client.WithCookies(cookies),
			client.WithRetry(1, handler.DefaultRetryInitialInterval, handler.DefaultRetryMaxInterval),
			client.WithLogger(logger),
		}, opts...)...,
	), nil
}

// getProxiesFromEnv loads the proxies from the file specified in the ROAPI_PROXIES_FILE environment variable.
func getProxiesFromEnv(logger logger.Logger, useProxies bool) ([]*url.URL, error) {
	if !useProxies {
		return []*url.URL{}, nil
	}

	proxiesFile := os.Getenv("ROAPI_PROXIES_FILE")
	if proxiesFile == "" {
		return nil, ErrProxiesFileNotSet
	}

	proxies, err := readProxiesFromFile(proxiesFile)
	if err != nil {
		return nil, err
	}

	logger.Debug("Loaded proxies", zap.Int("count", len(proxies)))
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
func getCookiesFromEnv(logger logger.Logger, useCookie bool) ([]string, error) {
	if !useCookie {
		return []string{}, nil
	}

	cookie := os.Getenv("ROAPI_COOKIES_FILE")
	if cookie == "" {
		return nil, ErrCookiesFileNotSet
	}

	cookies, err := readCookiesFromFile(cookie)
	if err != nil {
		return nil, err
	}

	logger.Debug("Loaded cookies", zap.Int("count", len(cookies)))
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
