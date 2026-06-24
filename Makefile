## Variables
BINARY_NAME=chill
BUILD_DIR=bin

.PHONY: all build run test clean tidy fmt linter help

all: build

## build: Build the Go binary
build:
	@echo "Building binary..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) cmd/chill/main.go   

## install: Build and install the application
install: build
	@echo "Installing $(BINARY_NAME)..."
	@go install ./cmd/chill

# test: Run all unit tests
test:
	@echo "Running tests..."
	@go test -v -race ./...

## clean: Remove build artifacts
clean:
	@echo "Cleaning build directory..."
	@rm -rf $(BUILD_DIR)

## tidy: Add missing and remove unused modules
tidy:
	@go mod tidy

## fmt: Run go fmt against all packages
fmt:
	@go fmt ./...

## linter: Run golangci-lint
linter:
	@golangci-lint run ./...