# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2025-11-08

### Added
- Initial release of terraform-provider-utils
- Base64 encoding and decoding functions
- Cryptographic hash functions (SHA256, MD5)
- Deterministic UUID v4 generation
- String manipulation functions:
  - `slugify` - URL-friendly slug generation
  - `truncate` - String truncation with suffix
  - `reverse` - String reversal
  - `trim` - Whitespace trimming
  - `to_upper` - Uppercase conversion
  - `to_lower` - Lowercase conversion
- List operation functions:
  - `join` - Join list elements with separator
  - `split` - Split string into list
- Comprehensive documentation and examples
- GitHub Actions CI/CD pipeline
- Makefile for build automation
- Unit tests for all functions

[Unreleased]: https://github.com/gilbertrios/terraform-provider-utils/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/gilbertrios/terraform-provider-utils/releases/tag/v0.1.0
