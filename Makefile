.PHONY: test benchmark install-tools lint build

# Define the default target
default: test

BINARY_NAME = flake
BUILD_DIR = ./bin
GO_BUILD = go build

# Build the flake CLI tool
build:
	@echo "Building the flake CLI..."
	@mkdir -p $(BUILD_DIR)
	$(GO_BUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/flake.go

# Run unit tests
test:
	@echo "=== Running unit tests ==="
	go test ./...

# Run benchmarks
benchmark:
	@echo "=== Running benchmarks ==="
	go test -bench=. ./...

# Clean any generated files or binaries (if necessary)
clean:
	@echo "=== Cleaning up ==="
	# Add commands to clean generated files or binaries if needed

# Run golint
lint: install-tools
	@echo "=== Running golint ==="
	go vet ./...

# Install local tools
install-tools:
	@echo "=== Installing local tools ==="
	go get -u golang.org/x/lint/golint
