.PHONY: help build run test clean docker-build docker-run docker-stop envoy-up envoy-down client server

# Variables
BINARY_NAME=websocket-server
CLIENT_NAME=websocket-client
DOCKER_IMAGE=websocket-playground

help: ## Display this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build the server and client binaries
	@echo "Building binaries..."
	@go build -o bin/$(BINARY_NAME) cmd/server/main.go
	@go build -o bin/$(CLIENT_NAME) cmd/client/main.go
	@echo "Binaries built successfully!"

server: ## Run the WebSocket server
	@echo "Starting WebSocket server..."
	@go run cmd/server/main.go

client: ## Run the WebSocket client
	@echo "Starting WebSocket client..."
	@go run cmd/client/main.go

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@go clean

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker-compose build

docker-run: ## Run with Docker Compose (server only)
	@echo "Starting services with Docker Compose..."
	@docker-compose up -d

docker-stop: ## Stop Docker Compose services
	@echo "Stopping services..."
	@docker-compose down

docker-logs: ## Show Docker logs
	@docker-compose logs -f

envoy-up: ## Run with Envoy proxy enabled
	@echo "Starting services with Envoy proxy..."
	@docker-compose --profile envoy up -d

envoy-down: ## Stop all services including Envoy
	@echo "Stopping all services..."
	@docker-compose --profile envoy down

envoy-logs: ## Show Envoy logs
	@docker-compose logs -f envoy

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...

lint: ## Run linter (requires golangci-lint)
	@echo "Running linter..."
	@golangci-lint run

rebuild: clean docker-build ## Clean and rebuild Docker image

status: ## Show status of Docker services
	@docker-compose ps
