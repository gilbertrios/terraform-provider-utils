# Terraform Provider Utils

[![Go Version](https://img.shields.io/github/go-mod/go-version/gilbertrios/terraform-provider-utils)](https://golang.org)
[![License](https://img.shields.io/github/license/gilbertrios/terraform-provider-utils)](LICENSE)

A function-only Terraform provider that provides utility functions for data manipulation and transformation in your Terraform configurations.

## 🎯 Key Features

- **Encoding & Hashing** - Base64 encoding/decoding, SHA256, MD5 hashing
- **Deterministic ID Generation** - UUID v4 generation from seed values
- **String Manipulation** - Slugify, truncate, reverse, trim, case conversion
- **List Operations** - Join and split operations for list handling
- **Zero Configuration** - No provider configuration required
- **Lightweight** - Pure function provider with no external dependencies
- **Type-Safe** - Strong typing with proper error handling

## 🌟 What This Repo Demonstrates

### Terraform Best Practices
- ✅ Function-only provider implementation
- ✅ Terraform Plugin Framework usage
- ✅ Type-safe function definitions
- ✅ Comprehensive testing strategy

### Development Best Practices
- ✅ Clean, modular Go code
- ✅ Extensive unit test coverage
- ✅ CI/CD automation with GitHub Actions
- ✅ Cross-platform build support

### Documentation
- ✅ Comprehensive function reference
- ✅ Real-world usage examples
- ✅ Developer-friendly guides

## 🛠️ Tech Stack

**Application**
- Go 1.21+ - Modern Go with generics support
- Terraform Plugin Framework - Official provider framework
- Terraform 1.8+ - Provider-defined functions support

**DevOps**
- GitHub Actions - CI/CD automation
- Makefile - Build automation
- golangci-lint - Code quality checks

## 📋 Available Functions

| Category | Functions |
|----------|-----------|
| **Encoding & Hashing** | `base64_encode`, `base64_decode`, `sha256`, `md5` |
| **ID Generation** | `uuidv4` |
| **String Manipulation** | `slugify`, `truncate`, `reverse`, `trim`, `to_upper`, `to_lower` |
| **List Operations** | `join`, `split` |

See [Function Reference](docs/functions.md) for complete documentation.

## 💻 Quick Start

### Installation

```bash
git clone https://github.com/gilbertrios/terraform-provider-utils.git
cd terraform-provider-utils
make install
```

See [Installation Guide](docs/installation.md) for manual installation and platform-specific instructions.

### Basic Usage

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
  tags       = ["production", "web", "critical"]
  tag_string = provider::utils::join(local.tags, "-")
  # Result: "production-web-critical"
}
```

### Example: Content Hashing

```hcl
locals {
  config_content = jsonencode({
    version  = "1.0"
    features = ["auth", "api"]
  })
  
  # Generate content hash for cache busting
  config_hash = provider::utils::sha256(local.config_content)
}
```

## 🏗️ Repository Structure

```
terraform-provider-utils/
├── main.go                      # Provider entry point
├── go.mod                       # Go module definition
├── Makefile                     # Build automation
├── README.md                    # This file
├── LICENSE                      # MIT License
├── CHANGELOG.md                 # Version history
│
├── internal/
│   └── provider/
│       ├── provider.go          # Provider definition
│       ├── provider_test.go     # Provider tests
│       ├── functions.go         # Function implementations
│       └── functions_test.go    # Function tests
│
├── examples/                    # Example configurations
│   ├── basic/                   # Basic usage examples
│   └── advanced/                # Real-world use cases
│
└── docs/                        # Documentation
    ├── installation.md          # Installation guide
    ├── quickstart.md            # Quick start guide
    ├── functions.md             # Function reference
    ├── usage.md                 # Usage patterns
    ├── development.md           # Development guide
    └── contributing.md          # Contributing guidelines
```

## 📚 Documentation

### Getting Started
- [Installation Guide](docs/installation.md) - Install the provider
- [Quick Start Guide](docs/quickstart.md) - Get up and running quickly
- [Usage Guide](docs/usage.md) - Common patterns and best practices

### Reference
- [Function Reference](docs/functions.md) - Complete API documentation
- [Examples](examples/) - Working example configurations

### Development
- [Development Guide](docs/development.md) - Build and test the provider
- [Contributing Guidelines](docs/contributing.md) - How to contribute

## 🧪 Testing

Run the test suite:

```bash
# Run all tests
make test

# Run with coverage
make test-coverage

# Run specific test
go test ./internal/provider -run TestBase64Encode
```

See [Development Guide](docs/development.md) for detailed testing documentation.

## 🚀 CI/CD Pipeline

Automated workflow for:
- ✅ Running tests on multiple Go versions
- ✅ Linting and formatting checks
- ✅ Building binaries for multiple platforms
- ✅ Release automation ready

## 💡 Use Cases

### Consistent Resource Naming
Generate deterministic, URL-friendly names across environments:
```hcl
resource_name = provider::utils::slugify("${var.app_name} ${var.environment}")
```

### Content Hashing
Create cache keys and version identifiers:
```hcl
version_id = provider::utils::sha256(local.config_content)
```

### Data Processing
Transform external data sources (CSV, JSON) for Terraform:
```hcl
ip_list = provider::utils::split(data.http.allowed_ips.body, ",")
```

### Length Constraints
Handle cloud provider name length restrictions:
```hcl
bucket_name = provider::utils::truncate(local.full_name, 63, "")
```

See [Usage Guide](docs/usage.md) for more examples and patterns.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-function`)
3. Commit your changes (`git commit -m 'Add some amazing function'`)
4. Push to the branch (`git push origin feature/amazing-function`)
5. Open a Pull Request

See [Contributing Guidelines](docs/contributing.md) for detailed instructions.

### 🌐 Connect With Me
Interested in Infrastructure as Code, Azure, or DevOps? Let's connect!

- 💼 LinkedIn: [Connect with me](https://linkedin.com/in/gilbert-rios-22586918)
- 📧 Email: [gilbertrios@hotmail.com](mailto:gilbertrios@hotmail.com)
- 💡 GitHub: [@gilbertrios](https://github.com/gilbertrios)

## 🎓 Quick Links

- [Quick Start Guide](docs/quickstart.md) - Get started in 5 minutes
- [Function Reference](docs/functions.md) - Complete API documentation
- [Usage Guide](docs/usage.md) - Real-world patterns and examples
- [Examples Directory](examples/) - Working configurations

---

⭐ If you find this project useful, please consider giving it a star!
