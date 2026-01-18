package aws

import (
	"context"
	"fmt"

	"github.com/Priyans-hu/sreq/internal/providers"
)

// Provider implements the providers.Provider interface for AWS Secrets Manager
type Provider struct {
	region  string
	profile string
	paths   map[string]string
}

// Config holds AWS Secrets Manager provider configuration
type Config struct {
	Region  string
	Profile string
	Paths   map[string]string
}

// New creates a new AWS Secrets Manager provider
func New(cfg Config) (*Provider, error) {
	region := cfg.Region
	if region == "" {
		region = "us-east-1" // default
	}

	return &Provider{
		region:  region,
		profile: cfg.Profile,
		paths:   cfg.Paths,
	}, nil
}

// Name returns the provider name
func (p *Provider) Name() string {
	return "aws_secrets"
}

// Get retrieves a secret from AWS Secrets Manager
func (p *Provider) Get(ctx context.Context, key string) (string, error) {
	// TODO: Implement AWS Secrets Manager get
	// Use github.com/aws/aws-sdk-go-v2/service/secretsmanager
	return "", fmt.Errorf("aws secrets manager provider not yet implemented")
}

// GetMultiple retrieves multiple secrets from AWS Secrets Manager
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

// Health checks if AWS Secrets Manager is reachable
func (p *Provider) Health(ctx context.Context) error {
	// TODO: Implement health check
	return nil
}

// Ensure Provider implements the interface
var _ providers.Provider = (*Provider)(nil)
