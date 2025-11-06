# WebSocket Playground

A WebSocket service implementation in Go for testing and experimenting with WebSocket connections, including Envoy proxy integration.

## Overview

This project provides:
- A WebSocket server implementation in Go
- A simple WebSocket client for testing
- Docker containerization
- Envoy proxy configuration for WebSocket routing
- Example configurations and best practices

## Architecture

```
Client <-> Envoy Proxy <-> WebSocket Server
```

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- (Optional) Postman for WebSocket testing

## Project Structure

```
.
├── cmd/
│   ├── server/          # WebSocket server application
│   └── client/          # WebSocket client application
├── internal/
│   ├── handler/         # WebSocket handlers
│   └── config/          # Configuration management
├── deployments/
│   ├── docker/          # Dockerfiles
│   └── envoy/           # Envoy proxy configuration
├── scripts/             # Utility scripts
├── go.mod
├── go.sum
├── README.md
├── .gitignore
└── docker-compose.yml
```

## Quick Start

### 1. Run WebSocket Server Locally

```bash
# Install dependencies
go mod download

# Run the server
go run cmd/server/main.go
```

The server will start on `ws://localhost:8080/ws`

### 2. Test with Client

```bash
# In another terminal
go run cmd/client/main.go
```

### 3. Run with Docker

```bash
# Build and run
docker-compose up --build
```

### 4. Run with Envoy Proxy

```bash
# Start all services including Envoy
docker-compose --profile envoy up --build
```

## Testing with Postman

1. Open Postman
2. Create a new WebSocket Request
3. Enter URL: `ws://localhost:8080/ws`
4. Click "Connect"
5. Send messages and observe responses

## API Endpoints

### WebSocket Endpoint

- **URL**: `/ws`
- **Protocol**: WebSocket
- **Description**: Main WebSocket connection endpoint

### Health Check

- **URL**: `/health`
- **Method**: GET
- **Description**: Health check endpoint for monitoring

## Configuration

### Environment Variables

- `PORT`: Server port (default: 8080)
- `LOG_LEVEL`: Logging level (default: info)
- `READ_BUFFER_SIZE`: WebSocket read buffer size (default: 1024)
- `WRITE_BUFFER_SIZE`: WebSocket write buffer size (default: 1024)

## Development

### Running Tests

```bash
go test ./...
```

### Running with Hot Reload

```bash
# Install air
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

### Building

```bash
# Build server
go build -o bin/server cmd/server/main.go

# Build client
go build -o bin/client cmd/client/main.go
```

## Envoy Integration

The Envoy proxy is configured to:
- Route WebSocket traffic to the backend service
- Provide load balancing capabilities
- Enable monitoring and observability
- Support WebSocket upgrade mechanisms

Configuration file: `deployments/envoy/envoy.yaml`

## Troubleshooting

### Connection Refused

- Ensure the server is running
- Check if the port is not already in use
- Verify firewall settings

### WebSocket Upgrade Failed

- Check if the client is sending proper upgrade headers
- Verify Envoy configuration allows WebSocket upgrades
- Review server logs for error messages

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

MIT License - See LICENSE file for details

## Resources

- [Gorilla WebSocket Documentation](https://pkg.go.dev/github.com/gorilla/websocket)
- [Envoy WebSocket Documentation](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/http/upgrades)
- [WebSocket Protocol RFC](https://tools.ietf.org/html/rfc6455)
