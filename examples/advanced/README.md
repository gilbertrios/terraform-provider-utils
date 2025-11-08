# Advanced Example

This example demonstrates practical, real-world use cases for the utils provider functions in Terraform configurations.

## Use Cases Demonstrated

### 1. Resource Naming and Identification
- Generate deterministic UUIDs for resource identification
- Create URL-friendly slugs for resource names
- Generate content hashes for cache busting
- Truncate names to meet length restrictions

### 2. Credential Management
- Base64 encode sensitive data
- Generate deterministic user IDs from emails
- Hash passwords and secrets

### 3. Data Processing
- Parse CSV strings into lists
- Join list values into formatted strings
- Process configuration data

## Usage

```bash
# Install provider locally
cd ../..
make install

# Run the example
cd examples/advanced
terraform init
terraform plan
terraform apply
```

## Real-World Applications

These patterns are useful for:
- Consistent naming across multi-region deployments
- Cache invalidation strategies
- Processing external data sources
- Managing secrets in Terraform state
- Working with length-restricted resource names (AWS, Azure)
