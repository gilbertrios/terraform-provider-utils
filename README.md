# Terraform Provider Utils

A function-only Terraform provider that provides utility functions for data manipulation and transformation in your Terraform configurations.

[![Go Version](https://img.shields.io/github/go-mod/go-version/gilbertrios/terraform-provider-utils)](https://golang.org)
[![License](https://img.shields.io/github/license/gilbertrios/terraform-provider-utils)](LICENSE)

## Overview

This provider offers a collection of utility functions for common string manipulation, encoding, hashing, and data transformation tasks in Terraform. Unlike traditional providers that manage infrastructure resources, this is a **function-only provider** that enhances Terraform's built-in functions with additional capabilities.

## Features

- 🔐 **Encoding & Hashing**: Base64 encoding/decoding, SHA256, MD5
- 🆔 **ID Generation**: Deterministic UUID v4 generation
- 📝 **String Manipulation**: Slugify, truncate, reverse, trim, case conversion
- 📊 **List Operations**: Join and split operations
- 🚀 **Zero Configuration**: No provider configuration required
- 📦 **Lightweight**: Pure function provider with no external dependencies

## Available Functions

### Encoding & Hashing

| Function | Description | Example |
|----------|-------------|---------|
| `base64_encode(string)` | Encodes a string to base64 | `provider::utils::base64_encode("hello")` |
| `base64_decode(string)` | Decodes a base64 string | `provider::utils::base64_decode("aGVsbG8=")` |
| `sha256(string)` | Computes SHA256 hash (hex) | `provider::utils::sha256("password")` |
| `md5(string)` | Computes MD5 hash (hex) | `provider::utils::md5("content")` |

### ID Generation

| Function | Description | Example |
|----------|-------------|---------|
| `uuidv4(string)` | Generates deterministic UUID v4 | `provider::utils::uuidv4("seed-value")` |

### String Manipulation

| Function | Description | Example |
|----------|-------------|---------|
| `slugify(string)` | Converts to URL-friendly slug | `provider::utils::slugify("My Project")` → `"my-project"` |
| `truncate(string, length, suffix)` | Truncates with optional suffix | `provider::utils::truncate("long text", 5, "...")` → `"lo..."` |
| `reverse(string)` | Reverses a string | `provider::utils::reverse("hello")` → `"olleh"` |
| `trim(string)` | Removes leading/trailing whitespace | `provider::utils::trim("  text  ")` → `"text"` |
| `to_upper(string)` | Converts to uppercase | `provider::utils::to_upper("hello")` → `"HELLO"` |
| `to_lower(string)` | Converts to lowercase | `provider::utils::to_lower("HELLO")` → `"hello"` |

### List Operations

| Function | Description | Example |
|----------|-------------|---------|
| `join(list, separator)` | Joins list with separator | `provider::utils::join(["a", "b"], "-")` → `"a-b"` |
| `split(string, separator)` | Splits string into list | `provider::utils::split("a,b,c", ",")` → `["a", "b", "c"]` |

## Installation

### Local Development

1. Clone this repository:
   ```bash
   git clone https://github.com/gilbertrios/terraform-provider-utils.git
   cd terraform-provider-utils
   ```

2. Build and install the provider:
   ```bash
   make install
   ```

3. The provider will be installed to your local Terraform plugins directory.

### Manual Installation

1. Build the provider:
   ```bash
   go build -o terraform-provider-utils
   ```

2. Create the plugin directory:
   ```bash
   mkdir -p ~/.terraform.d/plugins/registry.terraform.io/gilbertrios/utils/0.1.0/darwin_arm64
   ```
   *Adjust the OS/architecture path as needed (linux_amd64, darwin_amd64, etc.)*

3. Move the binary:
   ```bash
   mv terraform-provider-utils ~/.terraform.d/plugins/registry.terraform.io/gilbertrios/utils/0.1.0/darwin_arm64/
   ```

## Usage

### Basic Setup

Add the provider to your Terraform configuration:

```hcl
terraform {
  required_providers {
    utils = {
      source = "gilbertrios/utils"
    }
  }
}

provider "utils" {}
```

### Example: Resource Naming

```hcl
locals {
  environment = "production"
  application = "web-app"
  
  # Generate URL-friendly resource name
  resource_name = provider::utils::slugify("${local.application} ${local.environment}")
  # Result: "web-app-production"
  
  # Create deterministic UUID
  resource_id = provider::utils::uuidv4(local.resource_name)
  # Result: "a1b2c3d4-e5f6-4789-a012-b3c4d5e6f7a8"
}
```

### Example: Data Transformation

```hcl
locals {
  # Parse CSV data
  ip_ranges = "10.0.1.0/24,10.0.2.0/24,10.0.3.0/24"
  ip_list   = provider::utils::split(local.ip_ranges, ",")
  
  # Join tags
  tags = ["production", "web", "critical"]
  tag_string = provider::utils::join(local.tags, "-")
  # Result: "production-web-critical"
}
```

### Example: Content Hashing

```hcl
locals {
  config_content = jsonencode({
    version = "1.0"
    features = ["auth", "api"]
  })
  
  # Generate content hash for cache busting
  config_hash = provider::utils::sha256(local.config_content)
}
```

## Examples

Check out the [examples](./examples) directory for complete working examples:

- **[basic](./examples/basic)**: Demonstrates all available functions
- **[advanced](./examples/advanced)**: Real-world use cases and patterns

To run an example:

```bash
cd examples/basic
terraform init
terraform plan
terraform apply
```

## Development

### Prerequisites

- [Go](https://golang.org/doc/install) 1.21+
- [Terraform](https://www.terraform.io/downloads.html) 1.8+

### Building

```bash
# Download dependencies
go mod download

# Build the provider
make build

# Install locally
make install

# Run tests
make test

# Format code
make fmt
```

### Project Structure

```
.
├── main.go                      # Provider entry point
├── go.mod                       # Go module definition
├── go.sum                       # Go dependencies checksums
├── Makefile                     # Build automation
├── README.md                    # This file
├── LICENSE                      # MIT License
├── CHANGELOG.md                 # Version history
├── .gitignore                   # Git ignore rules
├── .golangci.yml                # Linting configuration
├── .goreleaser.yml              # Release automation config
│
├── .github/
│   └── workflows/
│       └── ci.yml               # GitHub Actions CI/CD pipeline
│
├── internal/
│   └── provider/
│       ├── provider.go          # Provider definition
│       ├── provider_test.go     # Provider tests
│       ├── functions.go         # Function implementations
│       └── functions_test.go    # Function tests
│
├── examples/                    # Example configurations
│   ├── basic/
│   │   ├── main.tf             # Basic usage examples
│   │   └── README.md           # Basic example docs
│   └── advanced/
│       ├── main.tf             # Real-world use cases
│       └── README.md           # Advanced example docs
│
└── docs/                        # Documentation
    ├── QUICKSTART.md            # Quick start guide
    └── CONTRIBUTING.md          # Contributing guidelines
```

## Testing

```bash
# Run all tests
make test

# Run with coverage
make test-coverage

# Run specific test
go test ./internal/provider -run TestBase64Encode
```

## CI/CD

This project includes a GitHub Actions workflow that:
- Runs tests on multiple Go versions
- Performs linting and formatting checks
- Builds binaries for multiple platforms
- Could publish to Terraform Registry (configured but not active)

See [.github/workflows/ci.yml](.github/workflows/ci.yml) for details.

## Why Use This Provider?

### Use Cases

1. **Consistent Resource Naming**: Generate deterministic, URL-friendly names across environments
2. **Content Hashing**: Create cache keys and version identifiers
3. **Data Processing**: Transform external data sources (CSV, JSON) for use in Terraform
4. **Secret Management**: Encode/decode sensitive data in outputs
5. **Length Constraints**: Handle cloud provider name length restrictions

### Advantages

- **Type-Safe**: Functions are strongly typed with proper error handling
- **Deterministic**: Same inputs always produce same outputs (perfect for Terraform)
- **Self-Contained**: No external API calls or dependencies
- **Performance**: Pure computation, no I/O operations
- **Portable**: Works with any Terraform backend

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-function`)
3. Commit your changes (`git commit -m 'Add some amazing function'`)
4. Push to the branch (`git push origin feature/amazing-function`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework)
- Inspired by the need for additional utility functions in Terraform configurations

## Author

**Gilbert Rios**
- GitHub: [@gilbertrios](https://github.com/gilbertrios)

---

⭐ If you find this project useful, please consider giving it a star!
