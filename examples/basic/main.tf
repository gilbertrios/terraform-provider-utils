terraform {
  required_providers {
    utils = {
      source = "gilbertrios/utils"
    }
  }
}

provider "utils" {}

# String manipulation examples
locals {
  # Base64 encoding/decoding
  original_text = "Hello, Terraform!"
  encoded       = provider::utils::base64_encode(local.original_text)
  decoded       = provider::utils::base64_decode(local.encoded)

  # Hashing
  password      = "my-secret-password"
  sha256_hash   = provider::utils::sha256(local.password)
  md5_hash      = provider::utils::md5(local.password)

  # UUID generation (deterministic)
  resource_id = "my-unique-resource"
  uuid        = provider::utils::uuidv4(local.resource_id)

  # String transformations
  project_name = "My Awesome Project"
  slug         = provider::utils::slugify(local.project_name)
  uppercase    = provider::utils::to_upper(local.project_name)
  lowercase    = provider::utils::to_lower(local.project_name)

  # String manipulation
  long_description = "This is a very long description that needs to be truncated for display purposes"
  short_desc       = provider::utils::truncate(local.long_description, 30, "...")
  
  reversed_text = provider::utils::reverse("Hello")
  trimmed_text  = provider::utils::trim("  spaces around  ")

  # List operations
  tags      = ["dev", "production", "staging"]
  tags_str  = provider::utils::join(local.tags, ", ")
  
  csv_data     = "apple,banana,orange"
  fruits_list  = provider::utils::split(local.csv_data, ",")
}

# Output examples
output "encoding_example" {
  description = "Base64 encoding demonstration"
  value = {
    original = local.original_text
    encoded  = local.encoded
    decoded  = local.decoded
  }
}

output "hashing_example" {
  description = "Hash functions demonstration"
  value = {
    sha256 = local.sha256_hash
    md5    = local.md5_hash
  }
}

output "uuid_example" {
  description = "Deterministic UUID generation"
  value = {
    resource_id = local.resource_id
    uuid        = local.uuid
  }
}

output "transformation_example" {
  description = "String transformation demonstration"
  value = {
    original  = local.project_name
    slug      = local.slug
    uppercase = local.uppercase
    lowercase = local.lowercase
  }
}

output "manipulation_example" {
  description = "String manipulation demonstration"
  value = {
    original   = local.long_description
    truncated  = local.short_desc
    reversed   = local.reversed_text
    trimmed    = local.trimmed_text
  }
}

output "list_operations" {
  description = "List operations demonstration"
  value = {
    tags_joined  = local.tags_str
    fruits_split = local.fruits_list
  }
}
