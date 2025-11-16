# Usage Guide

This guide provides comprehensive examples and patterns for using Terraform Provider Utils in your Terraform configurations.

## Table of Contents

- [Basic Setup](#basic-setup)
- [Common Patterns](#common-patterns)
- [Real-World Examples](#real-world-examples)
- [Best Practices](#best-practices)

## Basic Setup

### Provider Configuration

Add to your Terraform configuration:

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

No configuration is required—the provider is ready to use immediately.

## Common Patterns

### 1. Resource Naming

Generate consistent, URL-friendly resource names:

```hcl
locals {
  environment = "production"
  application = "web-app"
  region      = "us-east-1"
  
  # Generate base name
  base_name = provider::utils::slugify("${local.application} ${local.environment}")
  # Result: "web-app-production"
  
  # Create region-specific name
  regional_name = provider::utils::slugify("${local.base_name} ${local.region}")
  # Result: "web-app-production-us-east-1"
}

resource "aws_s3_bucket" "app" {
  bucket = local.base_name
}
```

### 2. Deterministic ID Generation

Create stable, reproducible UUIDs:

```hcl
locals {
  resource_name = "my-database-production"
  
  # Generate deterministic UUID
  resource_id = provider::utils::uuidv4(local.resource_name)
  # Same input always produces same UUID
}

resource "aws_db_instance" "main" {
  identifier = local.resource_id
  # ... other configuration
}
```

### 3. Content Hashing

Generate version identifiers and cache keys:

```hcl
locals {
  lambda_code = file("${path.module}/lambda/handler.py")
  
  # Create content hash for versioning
  code_hash = provider::utils::sha256(local.lambda_code)
}

resource "aws_lambda_function" "api" {
  filename         = "lambda.zip"
  source_code_hash = local.code_hash
  # Lambda will update when code changes
}
```

### 4. Length-Constrained Naming

Handle cloud provider name length restrictions:

```hcl
locals {
  long_name = "very-long-application-name-that-exceeds-azure-limits"
  
  # Azure Storage Account (24 char limit)
  storage_name = provider::utils::truncate(
    provider::utils::slugify(local.long_name),
    24,
    ""
  )
  # Result: "verylong-applicationnam"
}
```

### 5. Data Transformation

Parse and transform external data:

```hcl
locals {
  # Parse CSV input
  allowed_ips = "10.0.1.0/24,10.0.2.0/24,10.0.3.0/24"
  ip_list     = provider::utils::split(local.allowed_ips, ",")
  
  # Create security group rules
  ingress_rules = [
    for ip in local.ip_list : {
      cidr_block = ip
      from_port  = 443
      to_port    = 443
    }
  ]
}
```

### 6. Environment-Specific Configuration

```hcl
variable "environment" {
  type = string
}

locals {
  # Normalize environment name
  env = provider::utils::to_lower(var.environment)
  
  # Generate environment-specific resources
  db_name = provider::utils::slugify("app-${local.env}-db")
  
  # Create deterministic passwords (not recommended for production!)
  admin_hash = provider::utils::sha256("admin-${local.env}")
}
```

## Real-World Examples

### Example 1: Multi-Region S3 Buckets

```hcl
variable "regions" {
  type    = list(string)
  default = ["us-east-1", "us-west-2", "eu-west-1"]
}

locals {
  app_name = "media-storage"
  
  # Generate bucket names for each region
  regional_buckets = {
    for region in var.regions :
    region => provider::utils::slugify("${local.app_name}-${region}")
  }
}

resource "aws_s3_bucket" "regional" {
  for_each = local.regional_buckets
  
  bucket = each.value
  # Result: "media-storage-us-east-1", "media-storage-us-west-2", etc.
}
```

### Example 2: Dynamic Tagging

```hcl
locals {
  tags = {
    Environment = "production"
    Application = "web-api"
    ManagedBy   = "terraform"
  }
  
  # Create tag string for resources that need it
  tag_string = provider::utils::join([
    for k, v in local.tags : "${k}:${v}"
  ], ",")
  # Result: "Environment:production,Application:web-api,ManagedBy:terraform"
}
```

### Example 3: Configuration Versioning

```hcl
locals {
  app_config = {
    version  = "2.0"
    features = ["auth", "api", "cache"]
    limits   = { requests = 1000, storage = "10GB" }
  }
  
  # Generate config hash for change detection
  config_json = jsonencode(local.app_config)
  config_hash = provider::utils::sha256(local.config_json)
  
  # Use hash in resource names for immutable deployments
  deployment_id = provider::utils::truncate(local.config_hash, 8, "")
}

resource "kubernetes_config_map" "app" {
  metadata {
    name = "app-config-${local.deployment_id}"
  }
  
  data = {
    config = local.config_json
  }
}
```

### Example 4: Secret Management

```hcl
locals {
  # Encode sensitive data for secure outputs
  database_url = "postgresql://user:pass@host:5432/db"
  encoded_url  = provider::utils::base64_encode(local.database_url)
}

output "database_connection" {
  value     = local.encoded_url
  sensitive = true
}

# Later decode in another module:
# decoded = provider::utils::base64_decode(var.encoded_connection)
```

### Example 5: Resource ID Correlation

```hcl
locals {
  project = "ecommerce"
  
  # Generate correlated IDs for related resources
  vpc_id        = provider::utils::uuidv4("${local.project}-vpc")
  subnet_id     = provider::utils::uuidv4("${local.project}-subnet-public")
  db_cluster_id = provider::utils::uuidv4("${local.project}-db-cluster")
  
  # Same inputs always produce same IDs across runs
}
```

### Example 6: Processing External Data

```hcl
data "http" "ip_whitelist" {
  url = "https://api.example.com/allowed-ips"
}

locals {
  # Parse comma-separated response
  allowed_ips = provider::utils::split(
    provider::utils::trim(data.http.ip_whitelist.response_body),
    ","
  )
  
  # Create security group rules
  security_rules = [
    for ip in local.allowed_ips : {
      cidr_block  = ip
      description = "Whitelisted IP ${ip}"
    }
  ]
}
```

### Example 7: Complex Name Generation

```hcl
variable "project_name" {
  type = string
}

variable "environment" {
  type = string
}

locals {
  # Multi-step transformation
  base_slug = provider::utils::slugify(var.project_name)
  env_lower = provider::utils::to_lower(var.environment)
  
  # Combine and truncate
  full_name = "${local.base_slug}-${local.env_lower}"
  
  # Different limits for different resources
  s3_bucket_name = provider::utils::truncate(local.full_name, 63, "")  # S3 limit
  azure_storage  = provider::utils::truncate(local.full_name, 24, "")  # Azure limit
  rds_identifier = provider::utils::truncate(local.full_name, 60, "")  # RDS limit
}
```

## Best Practices

### 1. Use Locals for Reusability

```hcl
locals {
  # Define once, use everywhere
  resource_prefix = provider::utils::slugify(var.project_name)
}

resource "aws_s3_bucket" "data" {
  bucket = "${local.resource_prefix}-data"
}

resource "aws_s3_bucket" "logs" {
  bucket = "${local.resource_prefix}-logs"
}
```

### 2. Validate Inputs

```hcl
variable "environment" {
  type = string
  
  validation {
    condition = contains(["dev", "staging", "prod"], 
                        provider::utils::to_lower(var.environment))
    error_message = "Environment must be dev, staging, or prod."
  }
}
```

### 3. Document Function Usage

```hcl
locals {
  # Generate URL-safe bucket name (S3 requires lowercase, no special chars)
  bucket_name = provider::utils::slugify("${var.app_name}-${var.environment}")
  
  # Create deterministic UUID for stable resource identification across applies
  db_identifier = provider::utils::uuidv4(local.bucket_name)
}
```

### 4. Combine with Terraform Built-ins

```hcl
locals {
  # Use provider functions with Terraform's built-in functions
  all_regions = ["us-east-1", "us-west-2", "eu-west-1"]
  
  bucket_names = [
    for region in local.all_regions :
    provider::utils::slugify("${var.app}-${region}")
  ]
  
  # Result: ["myapp-us-east-1", "myapp-us-west-2", "myapp-eu-west-1"]
}
```

### 5. Handle Edge Cases

```hcl
locals {
  # Safely handle potentially empty or null values
  user_input = var.custom_name != "" ? var.custom_name : "default-app"
  safe_name  = provider::utils::slugify(local.user_input)
  
  # Ensure minimum length
  final_name = length(local.safe_name) >= 3 ? local.safe_name : "app-${local.safe_name}"
}
```

### 6. Version Your Configurations

```hcl
locals {
  config_version = "1.0.0"
  
  # Include version in resource names for immutable infrastructure
  versioned_name = provider::utils::slugify(
    "${var.app_name}-v${replace(local.config_version, ".", "-")}"
  )
  # Result: "my-app-v1-0-0"
}
```

## Function Composition

Combine multiple functions for complex operations:

```hcl
locals {
  # Input: "  My Complex App Name 2024  "
  
  cleaned = provider::utils::trim(var.app_name)
  # Result: "My Complex App Name 2024"
  
  slugified = provider::utils::slugify(local.cleaned)
  # Result: "my-complex-app-name-2024"
  
  truncated = provider::utils::truncate(local.slugified, 20, "")
  # Result: "my-complex-app-name-"
  
  hashed = provider::utils::sha256(local.truncated)
  # Result: "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"
  
  short_hash = provider::utils::truncate(local.hashed, 8, "")
  # Result: "5e884898"
}
```

## Next Steps

- Explore [Function Reference](FUNCTIONS.md) for detailed API documentation
- Check [examples/](../examples/) for complete working configurations
- See [Development Guide](DEVELOPMENT.md) to contribute new functions

