# Variables
APP_NAME := nuitee-mohit-jain
CMD_DIR := .

# Load environment variables from .env file (if it exists)
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

# Targets

# Default target: build the application
all: build

# Build the application
build:
	@echo "Building the application..."
	@go build -o $(APP_NAME) $(CMD_DIR)/main.go

# Run the application
run: build
	@echo "Running the application..."
	@./$(APP_NAME)

# Clean up the build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f $(APP_NAME)

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy

# Phony targets
.PHONY: all build run clean test fmt vet lint deps