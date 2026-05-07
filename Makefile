.PHONY: test benchmark lint format tidy vet clean help

# Default target
.DEFAULT_GOAL := help

# Go parameters
GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
GOFMT=gofmt
GOMOD=$(GOCMD) mod

help: ## Display this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

test: ## Run all tests
	@echo "Running tests..."
	$(GOTEST) -v -race -coverprofile=coverage.out ./...
	@echo "Coverage report:"
	$(GOCMD) tool cover -func=coverage.out

test-short: ## Run tests without race detector (faster)
	@echo "Running tests (short)..."
	$(GOTEST) -v ./...

benchmark: ## Run benchmark tests
	@echo "Running benchmarks..."
	$(GOTEST) -bench=. -benchmem ./...

benchmark-cpu: ## Run CPU benchmarks with profiling
	@echo "Running CPU benchmarks..."
	$(GOTEST) -bench=. -benchmem -cpuprofile=cpu.prof ./...

lint: ## Run golangci-lint (requires golangci-lint installation)
	@echo "Running linter..."
	@which golangci-lint > /dev/null || (echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; exit 1)
	golangci-lint run ./... --fix

format: ## Format all Go files
	@echo "Formatting code..."
	$(GOFMT) -s -w .
	@echo "Done formatting"

format-check: ## Check if code is formatted
	@echo "Checking formatting..."
	@test -z "$$($(GOFMT) -s -l . | tee /dev/stderr)"

tidy: ## Tidy go.mod and go.sum
	@echo "Tidying go.mod..."
	$(GOMOD) tidy
	@echo "Done tidying"

vet: ## Run go vet
	@echo "Running go vet..."
	$(GOVET) ./...
	@echo "Done vetting"

clean: ## Clean build artifacts and test cache
	@echo "Cleaning..."
	$(GOCMD) clean
	rm -f coverage.out cpu.prof mem.prof
	@echo "Done cleaning"

check: vet format-check test ## Run vet, format check, and tests

all: format tidy vet test ## Run format, tidy, vet, and test
