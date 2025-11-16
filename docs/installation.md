# Installation Guide

This guide covers different ways to install the Terraform Provider Utils.

## Prerequisites

- [Terraform](https://www.terraform.io/downloads.html) 1.8+
- [Go](https://golang.org/doc/install) 1.21+ (for building from source)

## Local Development Installation

For local development and testing:

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

## Manual Installation

### Step 1: Build the Provider

```bash
go build -o terraform-provider-utils
```

### Step 2: Create Plugin Directory

Create the appropriate plugin directory for your OS and architecture:

**macOS (ARM64)**
```bash
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/gilbertrios/utils/0.1.0/darwin_arm64
```

**macOS (AMD64)**
```bash
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/gilbertrios/utils/0.1.0/darwin_amd64
```

**Linux (AMD64)**
```bash
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/gilbertrios/utils/0.1.0/linux_amd64
```

**Windows (AMD64)**
```powershell
mkdir %APPDATA%\terraform.d\plugins\registry.terraform.io\gilbertrios\utils\0.1.0\windows_amd64
```

### Step 3: Move the Binary

Move the compiled binary to the plugin directory:

**macOS/Linux**
```bash
mv terraform-provider-utils ~/.terraform.d/plugins/registry.terraform.io/gilbertrios/utils/0.1.0/{OS_ARCH}/
```

**Windows**
```powershell
move terraform-provider-utils.exe %APPDATA%\terraform.d\plugins\registry.terraform.io\gilbertrios\utils\0.1.0\windows_amd64\
```

Replace `{OS_ARCH}` with your platform (e.g., `darwin_arm64`, `linux_amd64`).

## Verify Installation

Create a test Terraform configuration:

```hcl
terraform {
  required_providers {
    utils = {
      source = "gilbertrios/utils"
    }
  }
}

provider "utils" {}

output "test" {
  value = provider::utils::slugify("Hello World")
}
```

Run:
```bash
terraform init
terraform plan
```

You should see the output: `"hello-world"`

## Troubleshooting

### Provider Not Found

If Terraform can't find the provider:

1. Verify the plugin directory path matches your OS and architecture
2. Check that the binary is executable: `chmod +x terraform-provider-utils`
3. Ensure the version in the directory path matches your configuration

### Architecture Mismatch

Run this command to find your system architecture:

```bash
# macOS/Linux
uname -m

# Windows (PowerShell)
$env:PROCESSOR_ARCHITECTURE
```

Common outputs:
- `arm64` → use `darwin_arm64` or `linux_arm64`
- `x86_64` → use `darwin_amd64`, `linux_amd64`, or `windows_amd64`

### Permission Issues

If you encounter permission errors:

```bash
# macOS/Linux
chmod +x terraform-provider-utils
```

## Next Steps

- See [Quick Start Guide](QUICKSTART.md) for usage examples
- Explore [Function Reference](FUNCTIONS.md) for all available functions
- Check out [examples](../examples) for real-world use cases
