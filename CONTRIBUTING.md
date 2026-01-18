# Contributing to sreq

First off, thanks for taking the time to contribute! 🎉

The following is a set of guidelines for contributing to sreq. These are mostly guidelines, not rules. Use your best judgment, and feel free to propose changes to this document in a pull request.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How Can I Contribute?](#how-can-i-contribute)
  - [Reporting Bugs](#reporting-bugs)
  - [Suggesting Enhancements](#suggesting-enhancements)
  - [Your First Code Contribution](#your-first-code-contribution)
  - [Pull Requests](#pull-requests)
- [Development Setup](#development-setup)
- [Style Guidelines](#style-guidelines)
- [Commit Messages](#commit-messages)

## Code of Conduct

This project and everyone participating in it is governed by the [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check existing issues to avoid duplicates. When you create a bug report, include as many details as possible:

- **Use a clear and descriptive title**
- **Describe the exact steps to reproduce the problem**
- **Provide specific examples** (config snippets, commands run)
- **Describe the behavior you observed and what you expected**
- **Include your environment** (OS, Go version, sreq version)

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion:

- **Use a clear and descriptive title**
- **Provide a detailed description of the proposed enhancement**
- **Explain why this enhancement would be useful**
- **List any alternatives you've considered**

### Your First Code Contribution

Unsure where to begin? Look for issues labeled:

- `good first issue` - Simple issues for newcomers
- `help wanted` - Issues that need attention

### Pull Requests

1. Fork the repo and create your branch from `main`
2. If you've added code, add tests
3. Ensure the test suite passes
4. Make sure your code follows the style guidelines
5. Write a clear PR description

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Git

### Setup

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/sreq.git
cd sreq

# Add upstream remote
git remote add upstream https://github.com/Priyans-hu/sreq.git

# Install dependencies
go mod download

# Build
go build -o sreq ./cmd/sreq

# Run tests
go test ./...
```

### Project Structure

```
sreq/
├── cmd/sreq/           # CLI entrypoint
│   └── cmd/            # Cobra commands
├── internal/
│   ├── config/         # Configuration management
│   ├── client/         # HTTP client
│   └── providers/      # Secret providers
│       ├── consul/     # Consul KV provider
│       └── aws/        # AWS Secrets Manager provider
├── pkg/
│   └── types/          # Public types
└── docs/               # Documentation
```

### Running Locally

```bash
# Build and run
go build -o sreq ./cmd/sreq
./sreq --help

# Or use go run
go run ./cmd/sreq --help
```

## Style Guidelines

### Go Code

- Follow [Effective Go](https://golang.org/doc/effective_go)
- Use `gofmt` for formatting
- Use `golint` and `go vet` for linting
- Keep functions small and focused
- Write meaningful comments for exported functions

```bash
# Format code
gofmt -w .

# Run linters
go vet ./...
golangci-lint run
```

### Documentation

- Use clear, concise language
- Include code examples where helpful
- Keep README and docs up to date with changes

## Commit Messages

Follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
<type>(<scope>): <description>

[optional body]

[optional footer]
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

### Examples

```
feat(providers): add HashiCorp Vault provider

fix(aws): handle pagination in secret listing

docs(readme): add installation instructions for Homebrew

chore(deps): update cobra to v1.8.0
```

## Questions?

Feel free to open an issue with your question or reach out to the maintainers.

Thank you for contributing! 🙏
