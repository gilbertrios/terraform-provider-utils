# Development Guide

This guide covers how to develop, test, and contribute to the Terraform Provider Utils.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Building the Provider](#building-the-provider)
- [Testing](#testing)
- [Code Quality](#code-quality)
- [Making Changes](#making-changes)
- [CI/CD Pipeline](#cicd-pipeline)

## Prerequisites

### Required

- **Go**: 1.21 or later ([Installation Guide](https://golang.org/doc/install))
- **Terraform**: 1.8 or later ([Installation Guide](https://www.terraform.io/downloads.html))
- **Make**: For build automation (pre-installed on macOS/Linux)

### Recommended

- **golangci-lint**: For code linting ([Installation Guide](https://golangci-lint.run/usage/install/))
- **VS Code** with Go extension: For development
- **Git**: For version control

## Getting Started

1. **Clone the repository:**
   ```bash
   git clone https://github.com/gilbertrios/terraform-provider-utils.git
   cd terraform-provider-utils
   ```

2. **Download dependencies:**
   ```bash
   go mod download
   ```

3. **Build the provider:**
   ```bash
   make build
   ```

4. **Install locally:**
   ```bash
   make install
   ```

5. **Verify installation:**
   ```bash
   cd examples/basic
   terraform init
   terraform plan
   ```

## Project Structure

```
terraform-provider-utils/
├── main.go                      # Provider entry point
├── go.mod                       # Go module definition
├── go.sum                       # Dependency checksums
├── Makefile                     # Build automation
├── README.md                    # Root documentation
├── LICENSE                      # MIT License
├── CHANGELOG.md                 # Version history
│
├── internal/
│   └── provider/
│       ├── provider.go          # Provider schema & configuration
│       ├── provider_test.go     # Provider-level tests
│       ├── functions.go         # Function implementations
│       └── functions_test.go    # Function unit tests
│
├── examples/                    # Example configurations
│   ├── basic/                   # Basic usage examples
│   └── advanced/                # Advanced patterns
│
└── docs/                        # Documentation
    ├── INSTALLATION.md          # Installation guide
    ├── QUICKSTART.md            # Quick start guide
    ├── FUNCTIONS.md             # Function reference
    ├── DEVELOPMENT.md           # This file
    └── CONTRIBUTING.md          # Contributing guidelines
```

### Key Files

- **`main.go`**: Entry point that registers the provider with Terraform's plugin framework
- **`internal/provider/provider.go`**: Provider definition and function registration
- **`internal/provider/functions.go`**: All function implementations
- **`internal/provider/*_test.go`**: Test files using Go's testing framework

## Building the Provider

### Standard Build

```bash
# Build for your current platform
make build

# Output: ./terraform-provider-utils
```

### Cross-Platform Build

```bash
# Build for specific OS/architecture
GOOS=linux GOARCH=amd64 go build -o terraform-provider-utils

# Common platforms:
# macOS ARM64:   GOOS=darwin GOARCH=arm64
# macOS AMD64:   GOOS=darwin GOARCH=amd64
# Linux AMD64:   GOOS=linux GOARCH=amd64
# Windows AMD64: GOOS=windows GOARCH=amd64
```

### Install Locally

```bash
# Build and install to Terraform plugin directory
make install
```

This installs to: `~/.terraform.d/plugins/registry.terraform.io/gilbertrios/utils/0.1.0/{OS_ARCH}/`

## Testing

### Run All Tests

```bash
make test
```

### Run Tests with Coverage

```bash
make test-coverage

# View coverage in browser
go tool cover -html=coverage.out
```

### Run Specific Tests

```bash
# Test a specific function
go test ./internal/provider -run TestBase64Encode -v

# Test with verbose output
go test ./internal/provider -v

# Test a specific package
go test ./internal/provider/...
```

### Test Structure

Each function has corresponding tests in `functions_test.go`:

```go
func TestBase64Encode(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"simple", "hello", "aGVsbG8="},
        {"empty", "", ""},
        {"unicode", "hello 世界", "aGVsbG8g5LiW55WM"},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := base64Encode(tt.input)
            if result != tt.expected {
                t.Errorf("expected %s, got %s", tt.expected, result)
            }
        })
    }
}
```

### Integration Testing

Test with real Terraform configurations:

```bash
cd examples/basic
terraform init
terraform plan
terraform apply -auto-approve
terraform destroy -auto-approve
```

## Code Quality

### Formatting

```bash
# Format all Go code
make fmt

# Check formatting
go fmt ./...
```

### Linting

```bash
# Run linter (requires golangci-lint)
golangci-lint run

# Auto-fix issues
golangci-lint run --fix
```

### Pre-commit Checks

Before committing, run:

```bash
make fmt test
```

## Making Changes

### Adding a New Function

1. **Define the function in `internal/provider/functions.go`:**

```go
func NewMyFunction() function.Function {
    return &myFunction{}
}

type myFunction struct{}

func (f *myFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
    resp.Name = "my_function"
}

func (f *myFunction) Definition(_ context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
    resp.Definition = function.Definition{
        Summary: "Brief description",
        Description: "Detailed description",
        Parameters: []function.Parameter{
            function.StringParameter{
                Name: "input",
                Description: "Input parameter description",
            },
        },
        Return: function.StringReturn{},
    }
}

func (f *myFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
    var input string
    
    resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))
    if resp.Error != nil {
        return
    }
    
    // Implementation here
    result := processInput(input)
    
    resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}
```

2. **Register in `internal/provider/provider.go`:**

```go
func (p *utilsProvider) Functions(_ context.Context) []func() function.Function {
    return []func() function.Function{
        // ... existing functions
        NewMyFunction,
    }
}
```

3. **Add tests in `internal/provider/functions_test.go`:**

```go
func TestMyFunction(t *testing.T) {
    // Test cases here
}
```

4. **Document in `docs/FUNCTIONS.md`**

5. **Add example in `examples/basic/main.tf`**

### Modifying an Existing Function

1. Update implementation in `functions.go`
2. Update or add tests in `functions_test.go`
3. Update documentation in `docs/FUNCTIONS.md`
4. Update `CHANGELOG.md`

### Running Examples

After making changes:

```bash
# Rebuild and reinstall
make install

# Test with examples
cd examples/basic
terraform init -upgrade
terraform plan
```

## CI/CD Pipeline

The project includes GitHub Actions workflows for automated testing and building.

### Continuous Integration

On every push and pull request:
- Runs tests on multiple Go versions
- Performs linting
- Checks code formatting
- Builds for multiple platforms

### Manual Testing

```bash
# Simulate CI locally
make fmt test build

# Check for common issues
go vet ./...
staticcheck ./...
```

## Makefile Commands

```bash
make build          # Build the provider binary
make install        # Build and install to local plugin directory
make test           # Run all tests
make test-coverage  # Run tests with coverage report
make fmt            # Format all Go code
make clean          # Remove build artifacts
```

## Debugging

### Enable Terraform Debug Logging

```bash
export TF_LOG=DEBUG
export TF_LOG_PATH=terraform.log
terraform plan
```

### Debug Provider Code

Add debug prints (removed before commit):

```go
fmt.Printf("DEBUG: input=%v\n", input)
```

Or use the provider's logging:

```go
resp.Diagnostics.AddWarning("Debug", fmt.Sprintf("Value: %v", value))
```

## Version Management

This project follows [Semantic Versioning](https://semver.org/):

- **MAJOR**: Breaking changes
- **MINOR**: New features (backward compatible)
- **PATCH**: Bug fixes (backward compatible)

Update version in:
- `CHANGELOG.md`
- Git tags: `git tag v0.2.0`

## Resources

- [Terraform Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework)
- [Provider Functions Guide](https://developer.hashicorp.com/terraform/plugin/framework/functions)
- [Go Testing Documentation](https://golang.org/pkg/testing/)
- [Effective Go](https://golang.org/doc/effective_go)

## Getting Help

- Create an issue on GitHub
- Check existing issues and PRs
- Review the [Contributing Guidelines](CONTRIBUTING.md)

---

Ready to contribute? Check out the [Contributing Guidelines](CONTRIBUTING.md)!
