.PHONY: help build install test test-coverage fmt lint clean

# Default target
help:
	@echo "Available targets:"
	@echo "  build          - Build the provider binary"
	@echo "  install        - Install the provider locally"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  fmt            - Format Go code"
	@echo "  lint           - Run golangci-lint"
	@echo "  clean          - Remove build artifacts"

# Variables
BINARY_NAME=terraform-provider-utils
VERSION?=0.1.0
PROVIDER_NAMESPACE=gilbertrios
PROVIDER_TYPE=utils

# Detect OS and architecture
UNAME_S := $(shell uname -s)
UNAME_M := $(shell uname -m)

ifeq ($(UNAME_S),Darwin)
    OS=darwin
endif
ifeq ($(UNAME_S),Linux)
    OS=linux
endif

ifeq ($(UNAME_M),x86_64)
    ARCH=amd64
endif
ifeq ($(UNAME_M),arm64)
    ARCH=arm64
endif
ifeq ($(UNAME_M),aarch64)
    ARCH=arm64
endif

PLUGIN_DIR=~/.terraform.d/plugins/registry.terraform.io/$(PROVIDER_NAMESPACE)/$(PROVIDER_TYPE)/$(VERSION)/$(OS)_$(ARCH)

# Build the provider
build:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) -ldflags="-X 'main.version=$(VERSION)'" .

# Install the provider to local Terraform plugins directory
install: build
	@echo "Installing provider to $(PLUGIN_DIR)..."
	@mkdir -p $(PLUGIN_DIR)
	@cp $(BINARY_NAME) $(PLUGIN_DIR)/
	@echo "Provider installed successfully!"
	@echo ""
	@echo "You can now use it in your Terraform configurations:"
	@echo "  terraform {"
	@echo "    required_providers {"
	@echo "      utils = {"
	@echo "        source = \"$(PROVIDER_NAMESPACE)/$(PROVIDER_TYPE)\""
	@echo "      }"
	@echo "    }"
	@echo "  }"

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Run linter (requires golangci-lint to be installed)
lint:
	@echo "Running linter..."
	@which golangci-lint > /dev/null || (echo "golangci-lint not found. Install from https://golangci-lint.run/usage/install/" && exit 1)
	golangci-lint run ./...

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
	@rm -f coverage.out coverage.html
	@rm -rf dist/
	@echo "Clean complete!"

# Development mode - build and install
dev: install
	@echo "Provider ready for development!"

# Build for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	@mkdir -p dist
	GOOS=darwin GOARCH=amd64 go build -o dist/$(BINARY_NAME)_darwin_amd64 -ldflags="-X 'main.version=$(VERSION)'" .
	GOOS=darwin GOARCH=arm64 go build -o dist/$(BINARY_NAME)_darwin_arm64 -ldflags="-X 'main.version=$(VERSION)'" .
	GOOS=linux GOARCH=amd64 go build -o dist/$(BINARY_NAME)_linux_amd64 -ldflags="-X 'main.version=$(VERSION)'" .
	GOOS=linux GOARCH=arm64 go build -o dist/$(BINARY_NAME)_linux_arm64 -ldflags="-X 'main.version=$(VERSION)'" .
	GOOS=windows GOARCH=amd64 go build -o dist/$(BINARY_NAME)_windows_amd64.exe -ldflags="-X 'main.version=$(VERSION)'" .
	@echo "Multi-platform build complete! Binaries are in dist/"
