# Quick Start Guide

This guide will help you get up and running with the terraform-provider-utils in under 5 minutes.

## Installation

### Step 1: Clone the repository

```bash
git clone https://github.com/gilbertrios/terraform-provider-utils.git
cd terraform-provider-utils
```

### Step 2: Build and install

```bash
make install
```

This will:
- Build the provider binary
- Install it to your local Terraform plugins directory (`~/.terraform.d/plugins/`)

## Usage

### Step 1: Create a Terraform configuration

Create a new directory and a `main.tf` file:

```bash
mkdir my-terraform-project
cd my-terraform-project
```

Create `main.tf`:

```hcl
terraform {
  required_providers {
    utils = {
      source = "gilbertrios/utils"
    }
  }
}

provider "utils" {}

locals {
  # Example: Create a slug from a project name
  project_name = "My Awesome Project"
  slug         = provider::utils::slugify(local.project_name)
  
  # Example: Generate a deterministic UUID
  uuid = provider::utils::uuidv4(local.slug)
  
  # Example: Hash a value
  config_hash = provider::utils::sha256("my-config-v1")
}

output "results" {
  value = {
    original = local.project_name
    slug     = local.slug
    uuid     = local.uuid
    hash     = local.config_hash
  }
}
```

### Step 2: Initialize and apply

```bash
terraform init
terraform plan
terraform apply
```

### Step 3: View the output

```bash
terraform output
```

You should see:

```
results = {
  "hash" = "a1b2c3d4e5f6..."
  "original" = "My Awesome Project"
  "slug" = "my-awesome-project"
  "uuid" = "12345678-1234-4567-8901-234567890abc"
}
```

## Try the Examples

We have prepared two complete examples:

### Basic Example

Shows all available functions:

```bash
cd examples/basic
terraform init
terraform apply
```

### Advanced Example

Real-world use cases:

```bash
cd examples/advanced
terraform init
terraform apply
```

## Common Use Cases

### 1. Resource Naming

```hcl
locals {
  env  = "production"
  app  = "web"
  name = provider::utils::slugify("${local.app} ${local.env}")
  # Result: "web-production"
}
```

### 2. Content Hashing

```hcl
locals {
  config = jsonencode({ version = "1.0", features = ["auth"] })
  hash   = provider::utils::sha256(local.config)
  # Use for cache keys, etags, etc.
}
```

### 3. String Truncation

```hcl
locals {
  description = "Very long description..."
  short       = provider::utils::truncate(local.description, 20, "...")
  # Result: "Very long descri..."
}
```

### 4. List Operations

```hcl
locals {
  tags      = ["prod", "web", "critical"]
  tag_str   = provider::utils::join(local.tags, "-")
  # Result: "prod-web-critical"
  
  csv       = "a,b,c"
  items     = provider::utils::split(local.csv, ",")
  # Result: ["a", "b", "c"]
}
```

## Next Steps

- Read the [README](../README.md) for complete function reference
- Check out [examples/basic](../examples/basic) for all functions
- Check out [examples/advanced](../examples/advanced) for real-world patterns
- Explore the [source code](../internal/provider) to understand implementation

## Troubleshooting

### Provider not found

If you get "provider not found", make sure you:

1. Ran `make install` from the provider directory
2. Used the correct source in your `required_providers` block
3. Are in the correct directory when running `terraform init`

### Build errors

If the build fails:

```bash
# Clean and rebuild
make clean
go mod tidy
make build
```

### Need help?

Open an issue on GitHub: https://github.com/gilbertrios/terraform-provider-utils/issues

## Development

To modify the provider:

1. Make your changes to the code
2. Run tests: `make test`
3. Rebuild and reinstall: `make install`
4. Test in your Terraform config

Happy Terraforming! 🚀
