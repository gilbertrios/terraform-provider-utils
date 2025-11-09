terraform {
  required_providers {
    utils = {
      source = "gilbertrios/utils"
    }
  }
}

provider "utils" {}

# Practical use case: Generating resource names and tags
locals {
  environment = "production"
  application = "web-app"
  region      = "us-east-1"

  # Generate a unique but deterministic ID for resources
  resource_seed = "${local.application}-${local.environment}-${local.region}"
  resource_id   = provider::utils::uuidv4(local.resource_seed)

  # Create URL-friendly resource names
  resource_name = provider::utils::slugify("${local.application} ${local.environment}")

  # Generate content hash for cache busting
  config_content = jsonencode({
    app = local.application
    env = local.environment
    region = local.region
  })
  config_hash = provider::utils::sha256(local.config_content)
  
  # Truncate for length-restricted fields (e.g., AWS resource name limits)
  long_description = "This is a detailed description of the ${local.resource_name} application running in ${local.environment} environment"
  short_name      = provider::utils::truncate(local.resource_name, 32, "")

  # Create formatted tags
  tag_values = [local.environment, local.resource_name, local.region]
  tags_string = provider::utils::join(local.tag_values, "-")
}

# Practical use case: Managing secrets and credentials
locals {
  # Base64 encode sensitive data for storage
  api_key_plain = "sk-1234567890abcdef"
  api_key_b64   = provider::utils::base64_encode(local.api_key_plain)
  
  # Generate deterministic IDs for external systems
  user_email = "user@example.com"
  user_uuid  = provider::utils::uuidv4(local.user_email)
}

# Practical use case: Processing CSV data
locals {
  # Parse comma-separated values
  allowed_ips = "10.0.1.0/24,10.0.2.0/24,10.0.3.0/24"
  ip_list     = provider::utils::split(local.allowed_ips, ",")
  
  # Join list values
  dns_servers = ["8.8.8.8", "8.8.4.4", "1.1.1.1"]
  dns_string  = provider::utils::join(local.dns_servers, ",")
}

# Outputs
output "resource_naming" {
  description = "Generated resource names and identifiers"
  value = {
    resource_id   = local.resource_id
    resource_name = local.resource_name
    short_name    = local.short_name
    config_hash   = local.config_hash
    tags_string   = local.tags_string
  }
}

output "credential_management" {
  description = "Credential encoding and ID generation"
  value = {
    api_key_encoded = local.api_key_b64
    user_uuid       = local.user_uuid
  }
  sensitive = true
}

output "data_processing" {
  description = "CSV and list processing"
  value = {
    ip_list    = local.ip_list
    dns_string = local.dns_string
  }
}
