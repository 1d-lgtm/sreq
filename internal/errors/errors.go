package errors

import (
	"fmt"
	"strings"
)

// Error types for better categorization
type ErrorType int

const (
	ErrConfig ErrorType = iota
	ErrAuth
	ErrProvider
	ErrNetwork
	ErrNotFound
	ErrValidation
)

// SreqError is a user-friendly error with suggestions
type SreqError struct {
	Type       ErrorType
	Message    string
	Cause      error
	Suggestion string
}

func (e *SreqError) Error() string {
	var sb strings.Builder
	sb.WriteString(e.Message)

	if e.Cause != nil {
		sb.WriteString(fmt.Sprintf("\n  Cause: %v", e.Cause))
	}

	if e.Suggestion != "" {
		sb.WriteString(fmt.Sprintf("\n  Suggestion: %s", e.Suggestion))
	}

	return sb.String()
}

func (e *SreqError) Unwrap() error {
	return e.Cause
}

// Config errors
func ConfigNotFound(path string) *SreqError {
	return &SreqError{
		Type:       ErrConfig,
		Message:    fmt.Sprintf("Configuration file not found: %s", path),
		Suggestion: "Run 'sreq init' to create the default configuration.",
	}
}

func ConfigParseError(path string, cause error) *SreqError {
	return &SreqError{
		Type:       ErrConfig,
		Message:    fmt.Sprintf("Failed to parse configuration file: %s", path),
		Cause:      cause,
		Suggestion: "Check the YAML syntax in your config file. Use a YAML validator if needed.",
	}
}

func ServiceNotFound(service string) *SreqError {
	return &SreqError{
		Type:       ErrNotFound,
		Message:    fmt.Sprintf("Service '%s' not found in configuration", service),
		Suggestion: fmt.Sprintf("Add the service using: sreq service add %s --consul-key <key>", service),
	}
}

func ContextNotFound(context string) *SreqError {
	return &SreqError{
		Type:       ErrNotFound,
		Message:    fmt.Sprintf("Context '%s' not found in configuration", context),
		Suggestion: "Check available contexts in ~/.sreq/config.yaml under the 'contexts' section.",
	}
}

// Auth errors
func ConsulAuthFailed(address string, cause error) *SreqError {
	return &SreqError{
		Type:       ErrAuth,
		Message:    fmt.Sprintf("Failed to connect to Consul at %s", address),
		Cause:      cause,
		Suggestion: "Check that:\n  1. Consul is running and accessible\n  2. The address is correct\n  3. CONSUL_HTTP_TOKEN is set (if required)\n  Run 'sreq auth consul' to reconfigure.",
	}
}

func AWSAuthFailed(region string, cause error) *SreqError {
	return &SreqError{
		Type:       ErrAuth,
		Message:    fmt.Sprintf("Failed to authenticate with AWS in region %s", region),
		Cause:      cause,
		Suggestion: "Check that:\n  1. AWS credentials are configured (~/.aws/credentials or env vars)\n  2. The IAM user/role has secretsmanager:GetSecretValue permission\n  3. The region is correct\n  Run 'sreq auth aws' to reconfigure.",
	}
}

// Provider errors
func ProviderNotConfigured(provider string) *SreqError {
	return &SreqError{
		Type:       ErrProvider,
		Message:    fmt.Sprintf("Provider '%s' is not configured", provider),
		Suggestion: fmt.Sprintf("Add the provider configuration to ~/.sreq/config.yaml or run 'sreq auth %s'.", provider),
	}
}

func SecretNotFound(provider, key string) *SreqError {
	return &SreqError{
		Type:       ErrNotFound,
		Message:    fmt.Sprintf("Secret '%s' not found in %s", key, provider),
		Suggestion: "Check that:\n  1. The secret path is correct\n  2. You have permission to access the secret\n  3. The secret exists in the specified environment",
	}
}

func CredentialResolutionFailed(service, env string, cause error) *SreqError {
	return &SreqError{
		Type:       ErrProvider,
		Message:    fmt.Sprintf("Failed to resolve credentials for service '%s' in environment '%s'", service, env),
		Cause:      cause,
		Suggestion: "Run 'sreq config test' to verify provider connectivity.",
	}
}

// Network errors
func RequestFailed(url string, cause error) *SreqError {
	return &SreqError{
		Type:       ErrNetwork,
		Message:    fmt.Sprintf("HTTP request failed: %s", url),
		Cause:      cause,
		Suggestion: "Check that:\n  1. The service URL is correct and accessible\n  2. Your network connection is working\n  3. Any required VPN is connected",
	}
}

func BaseURLMissing(service, env string) *SreqError {
	return &SreqError{
		Type:       ErrValidation,
		Message:    fmt.Sprintf("Could not resolve base_url for service '%s' in environment '%s'", service, env),
		Suggestion: "Ensure the service has a base_url configured in Consul or the service config.",
	}
}

// Validation errors
func InvalidMethod(method string) *SreqError {
	return &SreqError{
		Type:       ErrValidation,
		Message:    fmt.Sprintf("Invalid HTTP method: %s", method),
		Suggestion: "Valid methods are: GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS",
	}
}

func MissingRequiredFlag(flag string) *SreqError {
	return &SreqError{
		Type:       ErrValidation,
		Message:    fmt.Sprintf("Required flag missing: --%s", flag),
		Suggestion: "See 'sreq run --help' for usage information.",
	}
}

// Helper to wrap external errors with context
func Wrap(cause error, message string) *SreqError {
	return &SreqError{
		Message: message,
		Cause:   cause,
	}
}
