package types

// ServiceConfig represents a service's configuration
type ServiceConfig struct {
	Name      string `yaml:"name"`
	ConsulKey string `yaml:"consul_key"`
	AWSPrefix string `yaml:"aws_prefix"`
}

// ProviderConfig represents a secret provider's configuration
type ProviderConfig struct {
	Type    string            `yaml:"type"`
	Address string            `yaml:"address,omitempty"`
	Token   string            `yaml:"token,omitempty"`
	Region  string            `yaml:"region,omitempty"`
	Profile string            `yaml:"profile,omitempty"`
	Paths   map[string]string `yaml:"paths"`
}

// Config represents the main configuration
type Config struct {
	Providers    map[string]ProviderConfig `yaml:"providers"`
	Environments []string                  `yaml:"environments"`
	DefaultEnv   string                    `yaml:"default_env"`
	Services     map[string]ServiceConfig  `yaml:"services"`
}

// ResolvedCredentials contains the resolved credentials for a service
type ResolvedCredentials struct {
	BaseURL  string
	Username string
	Password string
	APIKey   string
	Headers  map[string]string
}

// Request represents an HTTP request to be made
type Request struct {
	Method      string
	Path        string
	Service     string
	Environment string
	Body        string
	Headers     map[string]string
}

// Response represents an HTTP response
type Response struct {
	StatusCode int
	Status     string
	Headers    map[string][]string
	Body       []byte
}
