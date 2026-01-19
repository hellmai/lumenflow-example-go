.PHONY: all fmt lint vet test build clean gates

# Default target
all: gates build

# Format code
fmt:
	gofmt -w .

# Check format
fmt-check:
	@test -z "$$(gofmt -l . | tee /dev/stderr)"

# Run linter
lint:
	golangci-lint run

# Run go vet
vet:
	go vet ./...

# Run tests with coverage
test:
	go test -v -cover ./...

# Run all gates
gates: fmt-check lint vet test
	@echo "All gates passed!"

# Build the binary
build:
	go build -o bin/calculator ./cmd/calculator

# Clean build artifacts
clean:
	rm -rf bin/
	go clean -testcache
