package loadbalancer

import (
	"cdn-simulator/internal/server"
	"errors"
)

// LoadBalancer distributes traffic between servers
type LoadBalancer struct {
	Servers []*server.Server
	current int
}

// GetServer returns the server that has the requested file cached,
// or it falls back to the round-robin strategy if no server has the file cached.
func (lb *LoadBalancer) GetServer(filePath string) (*server.Server, error) {
	if len(lb.Servers) == 0 {
		return nil, errors.New("no servers available")
	}

	// First, check if any server has the file in cache
	for _, srv := range lb.Servers {
		if _, exists := srv.Cache[filePath]; exists {
			return srv, nil // Return the server that has the file cached
		}
	}

	// If no server has the file cached, fall back to round-robin strategy
	lb.current = (lb.current + 1) % len(lb.Servers)
	return lb.Servers[lb.current], nil
}
