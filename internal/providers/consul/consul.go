package consul

import (
	"context"
	"fmt"

	"github.com/Priyans-hu/sreq/internal/providers"
)

// Provider implements the providers.Provider interface for Consul KV
type Provider struct {
	address    string
	token      string
	datacenter string
	paths      map[string]string
}

// Config holds Consul provider configuration
type Config struct {
	Address    string
	Token      string
	Datacenter string
	Paths      map[string]string
}

// New creates a new Consul provider
func New(cfg Config) (*Provider, error) {
	if cfg.Address == "" {
		return nil, fmt.Errorf("consul address is required")
	}

	return &Provider{
		address:    cfg.Address,
		token:      cfg.Token,
		datacenter: cfg.Datacenter,
		paths:      cfg.Paths,
	}, nil
}

// Name returns the provider name
func (p *Provider) Name() string {
	return "consul"
}

// Get retrieves a value from Consul KV
func (p *Provider) Get(ctx context.Context, key string) (string, error) {
	// TODO: Implement Consul KV get
	// Use github.com/hashicorp/consul/api
	return "", fmt.Errorf("consul provider not yet implemented")
}

// GetMultiple retrieves multiple values from Consul KV
func (p *Provider) GetMultiple(ctx context.Context, keys []string) (map[string]string, error) {
	results := make(map[string]string)
	for _, key := range keys {
		value, err := p.Get(ctx, key)
		if err != nil {
			return nil, err
		}
		results[key] = value
	}
	return results, nil
}

// Health checks if Consul is reachable
func (p *Provider) Health(ctx context.Context) error {
	// TODO: Implement health check
	return nil
}

// Ensure Provider implements the interface
var _ providers.Provider = (*Provider)(nil)
