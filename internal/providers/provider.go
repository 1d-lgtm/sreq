package providers

import "context"

// Provider is the interface that all secret providers must implement
type Provider interface {
	// Name returns the provider name
	Name() string

	// Get retrieves a value by key
	Get(ctx context.Context, key string) (string, error)

	// GetMultiple retrieves multiple values by keys
	GetMultiple(ctx context.Context, keys []string) (map[string]string, error)

	// Health checks if the provider is reachable
	Health(ctx context.Context) error
}

// ProviderConfig contains common provider configuration
type ProviderConfig struct {
	Address string
	Token   string
	Paths   map[string]string
}
