# Roadmap

Detailed planning and feature specifications for sreq.

## Current Status

- [x] Project setup
- [x] Core CLI structure (Cobra)
- [ ] Consul provider
- [ ] AWS Secrets Manager provider
- [ ] HTTP client with auth
- [ ] Environment switching
- [ ] Config file management

## Planned Features

### Phase 1: Core Functionality

- Consul KV provider
- AWS Secrets Manager provider
- HTTP client with Basic Auth and API key support
- Environment switching (dev/staging/prod)
- YAML configuration management

### Phase 2: Enhanced Providers

#### HashiCorp Vault Provider

```yaml
providers:
  vault:
    address: https://vault.example.com:8200
    token: ${VAULT_TOKEN}
    paths:
      password: "secret/data/{service}/{env}#password"
```

#### Environment Variable Provider

For local development or CI/CD where secrets are in env vars:

```yaml
providers:
  env:
    paths:
      api_key: "{SERVICE}_API_KEY"
      password: "{SERVICE}_{ENV}_PASSWORD"
```

### Phase 3: Credential Caching

Cache credentials locally for faster requests and offline use.

**Commands:**
```bash
# Sync credentials for an environment
sreq sync dev

# Sync all environments
sreq sync --all

# Use cached credentials (offline mode)
sreq GET /api/v1/users -s auth-service --offline
```

**Cache location:** `~/.sreq/cache/{env}/{service}.json`

**Security measures:**
- File permissions `600` (owner read/write only)
- Optional encryption with machine key
- Configurable TTL (default: 1 hour)
- Environment variable to disable caching in CI/CD

**Cache structure:**
```
~/.sreq/
├── config.yaml
└── cache/
    ├── dev/
    │   ├── auth-service.json
    │   └── billing-service.json
    └── staging/
        └── auth-service.json
```

### Phase 4: Developer Experience

#### Request History

Save and replay previous requests:

```bash
sreq history           # List recent requests
sreq history 5         # Replay request #5
sreq history --clear   # Clear history
```

#### TUI Mode (Bubble Tea)

Interactive terminal UI for building and executing requests.

```bash
sreq tui
```

### Phase 5: Integrations

#### Homebrew Formula

```bash
brew install sreq
```

#### Bruno Extension

Export sreq configs to Bruno collections, or extend Bruno with sreq's credential resolution.

## Ideas / Backlog

- Response formatting (JSON pretty-print, jq-style filtering)
- Request templates
- Collection support (group related requests)
- Team config sharing via git
- Metrics/timing output
- Retry with backoff
- Proxy support
