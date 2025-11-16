# Function Reference

Complete reference for all available utility functions in the Terraform Provider Utils.

## Function Categories

- [Encoding & Hashing](#encoding--hashing)
- [ID Generation](#id-generation)
- [String Manipulation](#string-manipulation)
- [List Operations](#list-operations)

---

## Encoding & Hashing

### base64_encode

Encodes a string to base64 format.

**Signature:**
```hcl
provider::utils::base64_encode(string) → string
```

**Parameters:**
- `string` (string) - The string to encode

**Returns:** Base64-encoded string

**Example:**
```hcl
locals {
  encoded = provider::utils::base64_encode("hello world")
  # Result: "aGVsbG8gd29ybGQ="
}
```

**Use Cases:**
- Encoding configuration data
- Preparing data for APIs that require base64
- Encoding secrets for storage

---

### base64_decode

Decodes a base64-encoded string back to plain text.

**Signature:**
```hcl
provider::utils::base64_decode(string) → string
```

**Parameters:**
- `string` (string) - The base64 string to decode

**Returns:** Decoded plain text string

**Example:**
```hcl
locals {
  decoded = provider::utils::base64_decode("aGVsbG8gd29ybGQ=")
  # Result: "hello world"
}
```

**Error Handling:**
Returns an error if the input is not valid base64.

---

### sha256

Computes the SHA256 hash of a string and returns it as a hexadecimal string.

**Signature:**
```hcl
provider::utils::sha256(string) → string
```

**Parameters:**
- `string` (string) - The string to hash

**Returns:** SHA256 hash as hexadecimal string (64 characters)

**Example:**
```hcl
locals {
  config_content = jsonencode({ version = "1.0" })
  content_hash   = provider::utils::sha256(local.config_content)
  # Result: "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"
}
```

**Use Cases:**
- Content hashing for cache invalidation
- Generating version identifiers
- Creating content-addressable storage keys
- Detecting configuration changes

---

### md5

Computes the MD5 hash of a string and returns it as a hexadecimal string.

**Signature:**
```hcl
provider::utils::md5(string) → string
```

**Parameters:**
- `string` (string) - The string to hash

**Returns:** MD5 hash as hexadecimal string (32 characters)

**Example:**
```hcl
locals {
  content_hash = provider::utils::md5("content")
  # Result: "9a0364b9e99bb480dd25e1f0284c8555"
}
```

**Note:** MD5 is suitable for non-cryptographic purposes like checksums. For security-sensitive hashing, use SHA256.

---

## ID Generation

### uuidv4

Generates a deterministic UUID v4 from a seed string.

**Signature:**
```hcl
provider::utils::uuidv4(string) → string
```

**Parameters:**
- `string` (string) - The seed value for deterministic generation

**Returns:** UUID v4 string in standard format (e.g., "123e4567-e89b-12d3-a456-426614174000")

**Example:**
```hcl
locals {
  resource_name = "my-app-production"
  resource_id   = provider::utils::uuidv4(local.resource_name)
  # Result: "a1b2c3d4-e5f6-4789-a012-b3c4d5e6f7a8" (deterministic)
}
```

**Characteristics:**
- **Deterministic:** Same input always produces same UUID
- **RFC 4122 compliant:** Valid UUID v4 format
- **Reproducible:** Perfect for Terraform's declarative model

**Use Cases:**
- Generating stable resource identifiers
- Creating consistent IDs across environments
- Database record keys that need to be predictable

---

## String Manipulation

### slugify

Converts a string to a URL-friendly slug format.

**Signature:**
```hcl
provider::utils::slugify(string) → string
```

**Parameters:**
- `string` (string) - The string to convert to a slug

**Returns:** Lowercase slug with hyphens replacing spaces/special characters

**Example:**
```hcl
locals {
  project_name = "My Awesome Project!"
  slug         = provider::utils::slugify(local.project_name)
  # Result: "my-awesome-project"
}
```

**Transformation Rules:**
- Converts to lowercase
- Replaces spaces and special characters with hyphens
- Removes leading/trailing hyphens
- Collapses consecutive hyphens

**Use Cases:**
- AWS resource names (S3 buckets, Lambda functions)
- Azure resource names with length constraints
- URL-friendly identifiers
- Git branch names

---

### truncate

Truncates a string to a specified length with an optional suffix.

**Signature:**
```hcl
provider::utils::truncate(string, length, suffix) → string
```

**Parameters:**
- `string` (string) - The string to truncate
- `length` (number) - Maximum length (including suffix)
- `suffix` (string) - Optional suffix to append (e.g., "...")

**Returns:** Truncated string with suffix if applicable

**Example:**
```hcl
locals {
  long_name = "very-long-resource-name-that-exceeds-limits"
  short     = provider::utils::truncate(local.long_name, 20, "...")
  # Result: "very-long-resour..."
}
```

**Behavior:**
- If string is shorter than `length`, returns unchanged
- Suffix is included in the total length
- If `length` ≤ suffix length, returns suffix only

**Use Cases:**
- Enforcing cloud provider name length limits (Azure 24 chars, etc.)
- Creating display names
- Log message truncation

---

### reverse

Reverses the characters in a string.

**Signature:**
```hcl
provider::utils::reverse(string) → string
```

**Parameters:**
- `string` (string) - The string to reverse

**Returns:** String with characters in reverse order

**Example:**
```hcl
locals {
  reversed = provider::utils::reverse("hello")
  # Result: "olleh"
}
```

**Use Cases:**
- Data obfuscation
- String manipulation puzzles
- Reverse DNS lookups

---

### trim

Removes leading and trailing whitespace from a string.

**Signature:**
```hcl
provider::utils::trim(string) → string
```

**Parameters:**
- `string` (string) - The string to trim

**Returns:** String with whitespace removed from both ends

**Example:**
```hcl
locals {
  trimmed = provider::utils::trim("  hello world  ")
  # Result: "hello world"
}
```

**Whitespace Removed:**
- Spaces
- Tabs
- Newlines
- Carriage returns

**Use Cases:**
- Cleaning user input
- Processing external data sources
- Normalizing configuration values

---

### to_upper

Converts all characters in a string to uppercase.

**Signature:**
```hcl
provider::utils::to_upper(string) → string
```

**Parameters:**
- `string` (string) - The string to convert

**Returns:** Uppercase string

**Example:**
```hcl
locals {
  upper = provider::utils::to_upper("hello world")
  # Result: "HELLO WORLD"
}
```

**Use Cases:**
- Environment variables (e.g., "PRODUCTION")
- Constant identifiers
- Case-insensitive comparisons

---

### to_lower

Converts all characters in a string to lowercase.

**Signature:**
```hcl
provider::utils::to_lower(string) → string
```

**Parameters:**
- `string` (string) - The string to convert

**Returns:** Lowercase string

**Example:**
```hcl
locals {
  lower = provider::utils::to_lower("HELLO WORLD")
  # Result: "hello world"
}
```

**Use Cases:**
- Normalizing resource names
- DNS labels (must be lowercase)
- Case-insensitive matching

---

## List Operations

### join

Joins a list of strings with a separator.

**Signature:**
```hcl
provider::utils::join(list, separator) → string
```

**Parameters:**
- `list` (list of strings) - The list to join
- `separator` (string) - The separator to insert between elements

**Returns:** Single string with elements joined by separator

**Example:**
```hcl
locals {
  tags      = ["production", "web", "critical"]
  tag_string = provider::utils::join(local.tags, "-")
  # Result: "production-web-critical"
}
```

**Use Cases:**
- Creating composite names
- Building connection strings
- Formatting tag strings
- CSV generation

---

### split

Splits a string into a list using a separator.

**Signature:**
```hcl
provider::utils::split(string, separator) → list(string)
```

**Parameters:**
- `string` (string) - The string to split
- `separator` (string) - The separator to split on

**Returns:** List of strings

**Example:**
```hcl
locals {
  ip_ranges = "10.0.1.0/24,10.0.2.0/24,10.0.3.0/24"
  ip_list   = provider::utils::split(local.ip_ranges, ",")
  # Result: ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
}
```

**Use Cases:**
- Parsing CSV data
- Processing comma-separated inputs
- Splitting paths or URLs
- Environment variable parsing

---

## Combining Functions

Functions can be composed for complex transformations:

```hcl
locals {
  # Create a unique S3 bucket name
  app_name     = "My Application"
  environment  = "PROD"
  
  bucket_name = provider::utils::truncate(
    provider::utils::slugify("${local.app_name}-${local.environment}"),
    63,  # S3 bucket name limit
    ""
  )
  # Result: "my-application-prod"
  
  # Generate deterministic bucket ID
  bucket_id = provider::utils::uuidv4(local.bucket_name)
}
```

## Type Safety

All functions include type checking and will return errors if:
- Parameters are missing
- Parameter types don't match expected types
- Values are invalid (e.g., negative length, invalid base64)

Terraform will validate these at plan time, ensuring type safety before apply.
