package handler

import (
	"math/rand/v2"
	"net/http"
	"net/url"
	"sync"

	"go.uber.org/zap"
)

// ProxyManager handles proxy rotation and provides a proxy function for HTTP requests.
type ProxyManager struct {
	handler *Handler
	proxies []*url.URL   // List of available proxies
	current int          // Index of the current proxy in use
	mu      sync.RWMutex // Mutex to protect concurrent access
}

// NewProxyManager creates a new ProxyManager instance with the provided handler.
func NewProxyManager(handler *Handler) *ProxyManager {
	return &ProxyManager{
		handler: handler,
		proxies: []*url.URL{},
		current: 0,
		mu:      sync.RWMutex{},
	}
}

// NextProxy returns a function that selects the next proxy for each request.
// This method is designed to be used with http.Transport's Proxy field.
func (pm *ProxyManager) NextProxy(_ *http.Request) (*url.URL, error) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	// Select the current proxy
	proxy := pm.proxies[pm.current]

	// Move to the next proxy for the next request (round-robin)
	pm.current = (pm.current + 1) % len(pm.proxies)

	pm.handler.Logger.Debug("Next Proxy", zap.String("proxy", proxy.Host))

	return proxy, nil
}

// UpdateProxies updates the list of proxies at runtime.
// It replaces the existing proxy list with the new one and randomizes the starting proxy.
func (pm *ProxyManager) UpdateProxies(newProxies []*url.URL) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.proxies = newProxies
	if len(newProxies) > 0 {
		// Randomize the starting proxy to distribute load
		pm.current = rand.IntN(len(newProxies))
	}

	pm.handler.Logger.Debug("Proxy list updated", zap.Int("proxy_count", len(newProxies)))
}

// GetProxyCount returns the current number of proxies in the list.
func (pm *ProxyManager) GetProxyCount() int {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	return len(pm.proxies)
}
