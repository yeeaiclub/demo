#!/bin/bash

# A2A Go SDK Demo Startup Script
# This script automatically starts the server and client for demonstration

set -e

# Color definitions
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Print colored messages
print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

# Check if Go is installed
check_go() {
    if ! command -v go &> /dev/null; then
        print_error "Go is not installed. Please install Go 1.19 or higher first"
        exit 1
    fi
    
    GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
    print_success "Go version: $GO_VERSION"
}

# Check dependencies
check_deps() {
    print_info "Checking dependencies..."
    
    # Check server dependencies
    if [ ! -f "basic/server/go.mod" ]; then
        print_error "Server go.mod file does not exist"
        exit 1
    fi
    
    # Check client dependencies
    if [ ! -f "basic/client/go.mod" ]; then
        print_error "Client go.mod file does not exist"
        exit 1
    fi
    
    print_success "Dependency check completed"
}

# Install dependencies
install_deps() {
    print_info "Installing dependencies..."
    
    # Install server dependencies
    cd basic/server
    go mod tidy
    cd ../..
    
    # Install client dependencies
    cd basic/client
    go mod tidy
    cd ../..
    
    print_success "Dependencies installed"
}

# Start server
start_server() {
    print_info "Starting server..."
    
    cd basic/server
    go run . &
    SERVER_PID=$!
    cd ../..
    
    # Wait for server to start
    sleep 3
    
    # Check if server started successfully
    if curl -s http://localhost:8080/card > /dev/null; then
        print_success "Server started successfully (PID: $SERVER_PID)"
    else
        print_error "Server failed to start"
        kill $SERVER_PID 2>/dev/null || true
        exit 1
    fi
}

# Start client
start_client() {
    print_info "Starting client..."
    
    cd basic/client
    go run .
    cd ../..
}

# Cleanup function
cleanup() {
    print_info "Cleaning up resources..."
    
    if [ ! -z "$SERVER_PID" ]; then
        kill $SERVER_PID 2>/dev/null || true
        print_success "Server stopped"
    fi
}

# Set up signal handling
trap cleanup EXIT INT TERM

# Main function
main() {
    echo "🎯 A2A Go SDK Demo Startup Script"
    echo "================================"
    
    # Check Go
    check_go
    
    # Check dependencies
    check_deps
    
    # Install dependencies
    install_deps
    
    # Start server
    start_server
    
    # Start client
    start_client
    
    print_success "Demo completed!"
}

# Run main function
main "$@" 