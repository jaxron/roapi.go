set dotenv-load

# Build all packages
build:
    go build ./...

# Run linter (matches CI)
lint:
    golangci-lint run --timeout=30m

# Run linter and auto-fix issues
lint-fix:
    golangci-lint run --timeout=30m --fix

# Run all tests
test:
    go test ./...

# Run a specific test by name
test-run name *args:
    go test -run {{name}} {{args}}

# Format code
fmt:
    gofumpt -w .

# Generate enums
generate:
    go generate ./...
