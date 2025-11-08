# Basic Example

This example demonstrates all available functions in the utils provider with simple, straightforward use cases.

## Functions Demonstrated

- **base64_encode / base64_decode**: Encode and decode strings in base64
- **sha256 / md5**: Generate cryptographic hashes
- **uuidv4**: Generate deterministic UUIDs
- **slugify**: Convert strings to URL-friendly slugs
- **to_upper / to_lower**: Case conversion
- **truncate**: Shorten strings with optional suffix
- **reverse**: Reverse string characters
- **trim**: Remove whitespace
- **join / split**: List and string conversions

## Usage

```bash
# Install provider locally
cd ../..
make install

# Run the example
cd examples/basic
terraform init
terraform plan
terraform apply
```

## Expected Output

The plan will show various string transformations and manipulations applied to sample data.
