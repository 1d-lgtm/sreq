# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Initial project structure with Cobra CLI
- Consul KV provider for credential resolution
- AWS Secrets Manager provider
- YAML-based configuration system
- Service-aware credential resolution with `{service}` and `{env}` placeholders
- Basic HTTP client with authentication support
- Verbose mode for debugging requests
- Project documentation (README, CONTRIBUTING, CODE_OF_CONDUCT, SECURITY)

### Planned

- Credential caching for offline use
- HashiCorp Vault provider
- Environment variable provider
- Interactive mode for building requests
- Request history and replay
- Response formatting (JSON pretty-print, filtering)
- Bruno collection import/export

## [0.1.0] - Unreleased

### Added

- Core CLI structure
- `sreq run` command for executing requests
- Configuration file support (`.sreq.yaml`)
- Provider interface for extensible secret backends
- Consul and AWS Secrets Manager provider implementations

---

## Version History

| Version | Date       | Description              |
| ------- | ---------- | ------------------------ |
| 0.1.0   | TBD        | Initial release          |

[Unreleased]: https://github.com/Priyans-hu/sreq/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/Priyans-hu/sreq/releases/tag/v0.1.0
