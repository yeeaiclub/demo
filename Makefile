# A2A Go SDK Demo Makefile

.PHONY: help build-server build-client run-server run-client clean test

# Default target
help:
	@echo "🎯 A2A Go SDK Demo - Available Commands:"
	@echo ""
	@echo "📦 Build Commands:"
	@echo "  build-server    - Build server application"
	@echo "  build-client    - Build client application"
	@echo "  build          - Build all applications"
	@echo ""
	@echo "🚀 Run Commands:"
	@echo "  run-server     - Run server application"
	@echo "  run-client     - Run client application"
	@echo "  demo           - Run complete demo (requires two terminals)"
	@echo "  demo-auto      - Run automatic demo script"
	@echo ""
	@echo "🧹 Clean Commands:"
	@echo "  clean          - Clean build artifacts"
	@echo ""
	@echo "🧪 Test Commands:"
	@echo "  test           - Run all tests"
	@echo "  test-server    - Test server"
	@echo "  test-client    - Test client"

# Build server
build-server:
	@echo "🔨 Building server application..."
	cd basic/server && go build -o ../../bin/server .
	@echo "✅ Server build completed: bin/server"

# Build client
build-client:
	@echo "🔨 Building client application..."
	cd basic/client && go build -o ../../bin/client .
	@echo "✅ Client build completed: bin/client"

# Build all applications
build: build-server build-client
	@echo "🎉 All applications built!"

# Run server
run-server:
	@echo "🚀 Starting server..."
	cd basic/server && go run .

# Run client
run-client:
	@echo "🚀 Starting client..."
	cd basic/client && go run .

# Run complete demo
demo:
	@echo "🎯 Starting A2A Demo..."
	@echo "📋 Please follow these steps:"
	@echo ""
	@echo "1️⃣  In first terminal run: make run-server"
	@echo "2️⃣  In second terminal run: make run-client"
	@echo ""
	@echo "Or use these commands:"
	@echo "  Terminal 1: cd basic/server && go run ."
	@echo "  Terminal 2: cd basic/client && go run ."
	@echo ""
	@echo "💡 Or use automatic script:"
	@echo "  ./scripts/demo.sh"

# Run automatic demo script
demo-auto:
	@echo "🚀 Running automatic demo script..."
	@if [ -f "scripts/demo.sh" ]; then \
		chmod +x scripts/demo.sh; \
		./scripts/demo.sh; \
	else \
		echo "❌ Demo script not found: scripts/demo.sh"; \
		exit 1; \
	fi

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -rf bin/
	@echo "✅ Clean completed"

# Run all tests
test:
	@echo "🧪 Running all tests..."
	cd basic/server && go test ./...
	cd basic/client && go test ./...
	@echo "✅ Tests completed"

# Test server
test-server:
	@echo "🧪 Testing server..."
	cd basic/server && go test ./...

# Test client
test-client:
	@echo "🧪 Testing client..."
	cd basic/client && go test ./...

# Install dependencies
deps:
	@echo "📦 Installing dependencies..."
	cd basic/server && go mod tidy
	cd basic/client && go mod tidy
	@echo "✅ Dependencies installed"

# Format code
fmt:
	@echo "🎨 Formatting code..."
	cd basic/server && go fmt ./...
	cd basic/client && go fmt ./...
	@echo "✅ Code formatting completed"

# Code linting
lint:
	@echo "🔍 Code linting..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		cd basic/server && golangci-lint run; \
		cd basic/client && golangci-lint run; \
	else \
		echo "⚠️  golangci-lint not installed, skipping code linting"; \
		echo "   Install command: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi 