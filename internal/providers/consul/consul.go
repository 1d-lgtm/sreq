package consul

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	sreerrors "github.com/Priyans-hu/sreq/internal/errors"
	"github.com/Priyans-hu/sreq/internal/providers"
	"github.com/hashicorp/consul/api"
)

// ContextKey is the type for context keys used by this package
type ContextKey string

// EnvContextKey is used to pass environment to the provider via context
const EnvContextKey ContextKey = "sreq_env"

// Provider implements the providers.Provider interface for Consul KV
type Provider struct {
	defaultAddress string
	envAddresses   map[string]string // env -> address mapping
	token          string
	datacenter     string
	paths          map[string]string

	// Client pool - lazily created clients per address
	clients   map[string]*api.Client
	clientsMu sync.RWMutex
}

// Config holds Consul provider configuration
type Config struct {
	Address      string
	EnvAddresses map[string]string // env -> address overrides
	Token        string
	Datacenter   string
	Paths        map[string]string
}

// New creates a new Consul provider
func New(cfg Config) (*Provider, error) {
	// Must have at least one address (default or env-specific)
	if cfg.Address == "" && len(cfg.EnvAddresses) == 0 {
		return nil, sreerrors.ConsulAddressRequired()
	}

	// Resolve token from environment variable if needed
	token := cfg.Token
	if strings.HasPrefix(token, "${") && strings.HasSuffix(token, "}") {
		envVar := token[2 : len(token)-1]
		token = os.Getenv(envVar)
	}

	return &Provider{
		defaultAddress: cfg.Address,
		envAddresses:   cfg.EnvAddresses,
		token:          token,
		datacenter:     cfg.Datacenter,
		paths:          cfg.Paths,
		clients:        make(map[string]*api.Client),
	}, nil
}

// getAddressForEnv returns the appropriate address for the given environment
func (p *Provider) getAddressForEnv(env string) string {
	if p.envAddresses != nil {
		if addr, ok := p.envAddresses[env]; ok {
			return addr
		}
	}
	return p.defaultAddress
}

// getClientForEnv returns a Consul client for the given environment
// Clients are lazily created and cached
func (p *Provider) getClientForEnv(env string) (*api.Client, error) {
	address := p.getAddressForEnv(env)
	if address == "" {
		return nil, sreerrors.ConsulAddressRequired()
	}

	// Check if client already exists
	p.clientsMu.RLock()
	client, exists := p.clients[address]
	p.clientsMu.RUnlock()

	if exists {
		return client, nil
	}

	// Create new client
	p.clientsMu.Lock()
	defer p.clientsMu.Unlock()

	// Double-check after acquiring write lock
	if client, exists := p.clients[address]; exists {
		return client, nil
	}

	consulConfig := api.DefaultConfig()
	consulConfig.Address = address

	if p.token != "" {
		consulConfig.Token = p.token
	}

	if p.datacenter != "" {
		consulConfig.Datacenter = p.datacenter
	}

	client, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create consul client for %s: %w", address, err)
	}

	p.clients[address] = client
	return client, nil
}

// Name returns the provider name
func (p *Provider) Name() string {
	return "consul"
}

// Get retrieves a value from Consul KV
// Uses environment from context (set via EnvContextKey) to determine which Consul server to use
func (p *Provider) Get(ctx context.Context, key string) (string, error) {
	// Extract environment from context
	env, _ := ctx.Value(EnvContextKey).(string)

	client, err := p.getClientForEnv(env)
	if err != nil {
		return "", err
	}

	kv := client.KV()

	// Get the value
	pair, _, err := kv.Get(key, p.queryOptions(ctx))
	if err != nil {
		return "", sreerrors.ConsulGetFailed(key, err)
	}

	if pair == nil {
		return "", sreerrors.ConsulKeyNotFound(key)
	}

	return string(pair.Value), nil
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

// GetWithTemplate retrieves a value using a path template
// Supports placeholders: {service}, {env}, {region}, {project}
func (p *Provider) GetWithTemplate(ctx context.Context, template string, vars map[string]string) (string, error) {
	key := ResolvePath(template, vars)
	return p.Get(ctx, key)
}

// Health checks if Consul is reachable
// Checks all configured Consul servers (default + env-specific)
func (p *Provider) Health(ctx context.Context) error {
	// Collect all unique addresses to check
	addresses := make(map[string]bool)
	if p.defaultAddress != "" {
		addresses[p.defaultAddress] = true
	}
	for _, addr := range p.envAddresses {
		addresses[addr] = true
	}

	var lastErr error
	for addr := range addresses {
		client, err := p.getClientForEnv("") // Get client for this address
		if err != nil {
			lastErr = err
			continue
		}

		// Override: get specific client for this address
		p.clientsMu.RLock()
		if c, ok := p.clients[addr]; ok {
			client = c
		}
		p.clientsMu.RUnlock()

		_, err = client.Status().Leader()
		if err != nil {
			lastErr = fmt.Errorf("consul health check failed for %s: %w", addr, err)
		}
	}

	return lastErr
}

// HealthForEnv checks if Consul is reachable for a specific environment
func (p *Provider) HealthForEnv(ctx context.Context, env string) error {
	client, err := p.getClientForEnv(env)
	if err != nil {
		return err
	}

	_, err = client.Status().Leader()
	if err != nil {
		addr := p.getAddressForEnv(env)
		return fmt.Errorf("consul health check failed for %s (env: %s): %w", addr, env, err)
	}
	return nil
}

// ListKeys lists all keys under a prefix
func (p *Provider) ListKeys(ctx context.Context, prefix string) ([]string, error) {
	// Extract environment from context
	env, _ := ctx.Value(EnvContextKey).(string)

	client, err := p.getClientForEnv(env)
	if err != nil {
		return nil, err
	}

	kv := client.KV()

	keys, _, err := kv.Keys(prefix, "", p.queryOptions(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to list keys with prefix '%s': %w", prefix, err)
	}

	return keys, nil
}

// GetAddresses returns all configured addresses for debugging/display
func (p *Provider) GetAddresses() map[string]string {
	result := make(map[string]string)
	if p.defaultAddress != "" {
		result["default"] = p.defaultAddress
	}
	for env, addr := range p.envAddresses {
		result[env] = addr
	}
	return result
}

// queryOptions creates query options with context
func (p *Provider) queryOptions(ctx context.Context) *api.QueryOptions {
	opts := &api.QueryOptions{}
	if p.datacenter != "" {
		opts.Datacenter = p.datacenter
	}
	return opts.WithContext(ctx)
}

// ResolvePath replaces placeholders in a path template
// Placeholders: {service}, {env}, {region}, {project}
func ResolvePath(template string, vars map[string]string) string {
	result := template
	for key, value := range vars {
		result = strings.ReplaceAll(result, "{"+key+"}", value)
	}
	return result
}

// ResolvePathSimple is a convenience function for basic service/env resolution
func ResolvePathSimple(template, service, env string) string {
	return ResolvePath(template, map[string]string{
		"service": service,
		"env":     env,
	})
}

// Ensure Provider implements the interface
var _ providers.Provider = (*Provider)(nil)
