# Contributing to terraform-provider-utils

Thank you for your interest in contributing! This document provides guidelines and instructions for contributing to this project.

## Code of Conduct

Be respectful, inclusive, and constructive. We're all here to learn and build better tools.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check existing issues to avoid duplicates. When creating a bug report, include:

- **Clear title and description**
- **Steps to reproduce**
- **Expected vs actual behavior**
- **Terraform version** (`terraform version`)
- **Provider version**
- **Operating system**

### Suggesting Enhancements

Enhancement suggestions are welcome! Please provide:

- **Clear use case**: Why is this needed?
- **Example usage**: How would it be used?
- **Alternatives considered**: What other approaches did you think about?

### Adding New Functions

Want to add a new function? Great! Here's how:

1. **Check if it fits**: Functions should be:
   - Deterministic (same input = same output)
   - Pure (no side effects)
   - Generally useful
   - Not duplicating Terraform's built-in functions

2. **Implementation steps**:
   ```
   a. Add function implementation to internal/provider/functions.go
   b. Register function in internal/provider/provider.go
   c. Add tests to internal/provider/functions_test.go
   d. Update README.md function table
   e. Add example usage
   ```

3. **Function template**:
   ```go
   // NewMyFunction creates a new my_function
   func NewMyFunction() function.Function {
       return &MyFunction{}
   }
   
   type MyFunction struct{}
   
   func (f *MyFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
       resp.Name = "my_function"
   }
   
   func (f *MyFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
       resp.Definition = function.Definition{
           Summary:     "Short description",
           Description: "Detailed description",
           Parameters: []function.Parameter{
               function.StringParameter{
                   Name:        "input",
                   Description: "Input description",
               },
           },
           Return: function.StringReturn{},
       }
   }
   
   func (f *MyFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
       var input string
       resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))
       if resp.Error != nil {
           return
       }
       
       // Your logic here
       result := processInput(input)
       
       resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
   }
   ```

## Development Workflow

### Setup

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/terraform-provider-utils.git
cd terraform-provider-utils

# Add upstream remote
git remote add upstream https://github.com/gilbertrios/terraform-provider-utils.git

# Install dependencies
go mod download
```

### Making Changes

```bash
# Create a branch
git checkout -b feature/my-new-function

# Make your changes
# ... edit files ...

# Format code
make fmt

# Run tests
make test

# Build
make build
```

### Testing

#### Unit Tests

```bash
# Run all tests
make test

# Run specific test
go test ./internal/provider -run TestMyFunction

# Run with coverage
make test-coverage
```

#### Integration Tests

```bash
# Install locally
make install

# Test with actual Terraform
cd examples/basic
terraform init
terraform plan
```

### Commit Guidelines

Use conventional commits:

```
feat: add new string_replace function
fix: correct slugify behavior with unicode
docs: update README with new examples
test: add tests for truncate edge cases
chore: update dependencies
```

## Pull Request Process

1. **Update documentation**
   - Add/update function documentation in README.md
   - Add usage examples
   - Update QUICKSTART.md if needed

2. **Add tests**
   - Unit tests for new functions
   - Edge case testing

3. **Ensure CI passes**
   - All tests pass
   - Code is formatted
   - No linting errors

4. **Create PR**
   - Clear title and description
   - Reference any related issues
   - Include examples of new functionality

5. **Review process**
   - Address review comments
   - Keep commits clean and focused

## Project Structure

```
terraform-provider-utils/
├── main.go                    # Provider entry point
├── internal/
│   └── provider/
│       ├── provider.go        # Provider definition
│       ├── provider_test.go   # Provider tests
│       ├── functions.go       # Function implementations
│       └── functions_test.go  # Function tests
├── examples/                  # Usage examples
├── docs/                      # Documentation
└── .github/                   # CI/CD workflows
```

## Code Style

- Follow standard Go conventions
- Use `gofmt` and `goimports`
- Write clear, self-documenting code
- Add comments for complex logic
- Keep functions focused and single-purpose

## Questions?

- Open a [GitHub Discussion](https://github.com/gilbertrios/terraform-provider-utils/discussions)
- Create an [Issue](https://github.com/gilbertrios/terraform-provider-utils/issues)

Thank you for contributing! 🎉
