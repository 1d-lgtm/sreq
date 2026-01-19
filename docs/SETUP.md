# sreq Setup Guide

This guide helps you set up sreq (service-aware API client CLI) on your machine. It can be followed by humans or AI agents performing automated setup.

## Prerequisites

### Required

| Tool | Version | Check Command | Install |
|------|---------|---------------|---------|
| Go | 1.21+ | `go version` | https://go.dev/dl/ |
| Git | Any | `git --version` | https://git-scm.com/downloads |

### Optional (for specific providers)

| Tool | Purpose | Check Command | Install |
|------|---------|---------------|---------|
| AWS CLI | AWS Secrets Manager | `aws --version` | https://aws.amazon.com/cli/ |
| Consul CLI | HashiCorp Consul | `consul --version` | https://developer.hashicorp.com/consul/install |

## Installation

### Option 1: Build from Source (Recommended)

```bash
# Clone the repository
git clone https://github.com/Priyans-hu/sreq.git
cd sreq

# Build the binary
go build -o sreq ./cmd/sreq

# Move to PATH (optional)
sudo mv sreq /usr/local/bin/
# Or for user-local install:
mkdir -p ~/bin && mv sreq ~/bin/
# Ensure ~/bin is in your PATH
```

### Option 2: Go Install

```bash
go install github.com/Priyans-hu/sreq/cmd/sreq@latest
```

### Verify Installation

```bash
sreq --version
sreq --help
```

## Initial Setup

### Step 1: Initialize Configuration

```bash
sreq init
```

This creates:
- `~/.sreq/config.yaml` - Main configuration file
- `~/.sreq/.key` - Encryption key for credential caching (auto-generated)

### Step 2: Configure Your First Service

Edit `~/.sreq/config.yaml`:

```yaml
default_env: dev

services:
  my-api:
    environments:
      dev:
        provider: consul
        path: "services/my-api/dev"
      prod:
        provider: aws
        path: "prod/my-api/credentials"
```

### Step 3: Configure Providers

#### Consul Provider

```yaml
providers:
  consul:
    address: "consul.internal:8500"
    # Optional: environment-specific addresses
    env_addresses:
      prod: "consul-prod.internal:8500"
```

**Environment Variables:**
```bash
export CONSUL_HTTP_ADDR="consul.internal:8500"
export CONSUL_HTTP_TOKEN="your-acl-token"  # If ACL enabled
```

#### AWS Secrets Manager Provider

```yaml
providers:
  aws:
    region: "us-east-1"
```

**Authentication (choose one):**
```bash
# Option 1: Environment variables
export AWS_ACCESS_KEY_ID="your-key"
export AWS_SECRET_ACCESS_KEY="your-secret"
export AWS_REGION="us-east-1"

# Option 2: AWS CLI profile
aws configure --profile myprofile
export AWS_PROFILE="myprofile"

# Option 3: IAM role (EC2/ECS/Lambda)
# Automatically uses instance metadata
```

## Verify Setup

### Test Provider Connection

```bash
# Test Consul connection
sreq auth test consul

# Test AWS connection
sreq auth test aws
```

### Test a Service Request

```bash
# Make a test request
sreq run -s my-api -e dev GET /health
```

## Common Configuration Examples

### Example 1: Simple Consul Setup

```yaml
default_env: dev

providers:
  consul:
    address: "localhost:8500"

services:
  auth-service:
    environments:
      dev:
        provider: consul
        path: "services/auth/dev/credentials"
      prod:
        provider: consul
        path: "services/auth/prod/credentials"
```

### Example 2: Multi-Provider Setup

```yaml
default_env: dev

providers:
  consul:
    address: "consul.internal:8500"
    env_addresses:
      prod: "consul-prod.internal:8500"
  aws:
    region: "us-west-2"

services:
  # Development uses Consul, Production uses AWS
  payment-api:
    environments:
      dev:
        provider: consul
        path: "dev/payment/creds"
      staging:
        provider: consul
        path: "staging/payment/creds"
      prod:
        provider: aws
        path: "prod/payment-api"
```

### Example 3: Expected Credential Format

Credentials stored in providers should be JSON with these fields:

```json
{
  "base_url": "https://api.example.com",
  "username": "api-user",
  "password": "secret-password",
  "api_key": "optional-api-key",
  "headers": {
    "X-Custom-Header": "value"
  }
}
```

**Required field:** `base_url`
**Optional fields:** `username`, `password`, `api_key`, `headers`

## Directory Structure

After setup, your sreq directory looks like:

```
~/.sreq/
├── config.yaml      # Main configuration
├── .key             # Encryption key (auto-generated)
├── cache/           # Encrypted credential cache
│   ├── dev/
│   └── prod/
└── history.json     # Request history
```

## Environment Variables

| Variable | Purpose | Default |
|----------|---------|---------|
| `SREQ_CONFIG` | Custom config path | `~/.sreq/config.yaml` |
| `SREQ_NO_CACHE` | Disable caching (`1` to disable) | - |
| `CI` | Auto-disable cache in CI (`true` or `1`) | - |

## Troubleshooting

### "provider not configured"

```bash
# Check your config file
cat ~/.sreq/config.yaml

# Ensure provider is defined under 'providers:' section
```

### "failed to connect to Consul"

```bash
# Test Consul connectivity
curl http://consul.internal:8500/v1/status/leader

# Check environment variables
echo $CONSUL_HTTP_ADDR
echo $CONSUL_HTTP_TOKEN
```

### "AWS credentials not found"

```bash
# Verify AWS credentials
aws sts get-caller-identity

# Check environment
echo $AWS_PROFILE
echo $AWS_REGION
```

### "cache decryption failed"

```bash
# Clear corrupted cache
sreq cache clear

# Re-sync credentials
sreq sync dev
```

### Permission Issues

```bash
# Fix config directory permissions
chmod 700 ~/.sreq
chmod 600 ~/.sreq/config.yaml
chmod 600 ~/.sreq/.key
```

## Quick Reference

```bash
# Initialize
sreq init

# Make requests
sreq run -s SERVICE -e ENV METHOD /path
sreq run -s my-api GET /users
sreq run -s my-api POST /users -d '{"name":"test"}'

# View/manage history
sreq history
sreq history 5 --curl
sreq history 5 --replay

# Cache management
sreq sync dev
sreq sync --all
sreq cache status
sreq cache clear

# Interactive TUI
sreq tui

# Test providers
sreq auth test consul
sreq auth test aws
```

## For AI Agents

When setting up sreq programmatically:

1. **Check prerequisites**: Verify Go 1.21+ is installed
2. **Clone and build**: Use the build from source method
3. **Initialize**: Run `sreq init` to create config directory
4. **Configure**: Write config.yaml based on user's infrastructure
5. **Set credentials**: Configure environment variables for providers
6. **Verify**: Run `sreq auth test <provider>` to confirm connectivity

**Key files to create/modify:**
- `~/.sreq/config.yaml` - Service and provider configuration

**Key commands to verify setup:**
- `sreq --version` - Confirm installation
- `sreq auth test consul` - Test Consul provider
- `sreq auth test aws` - Test AWS provider
- `sreq run -s SERVICE GET /health` - Test end-to-end
