# Testing Guide

This guide covers different ways to test the WebSocket server.

## Table of Contents

1. [Testing with Go Client](#testing-with-go-client)
2. [Testing with Browser](#testing-with-browser)
3. [Testing with Postman](#testing-with-postman)
4. [Testing with Command Line Tools](#testing-with-command-line-tools)
5. [Testing with Envoy](#testing-with-envoy)

## Prerequisites

Start the server first:

```bash
# Option 1: Run locally
go run cmd/server/main.go

# Option 2: Run with Docker
docker-compose up

# Option 3: Run with Make
make server
```

The server will start on `ws://localhost:8080/ws`

---

## Testing with Go Client

The project includes a Go client for testing.

```bash
# Run the client
go run cmd/client/main.go

# Or with custom server address
go run cmd/client/main.go -addr localhost:8080 -path /ws
```

The client will:
1. Connect to the server
2. Send an initial "Hello" message
3. Allow you to type and send messages
4. Display received messages
5. Handle graceful shutdown with Ctrl+C

---

## Testing with Browser

The server includes a built-in test page.

1. Open your browser
2. Navigate to: `http://localhost:8080`
3. Click "Connect" to establish WebSocket connection
4. Type messages and click "Send Test Message"
5. Observe the messages log

This is the easiest way to quickly test if the server is working.

---

## Testing with Postman

Postman has built-in WebSocket support.

### Steps:

1. Open Postman
2. Click "New" → "WebSocket Request"
3. Enter URL: `ws://localhost:8080/ws`
4. Click "Connect"
5. In the message field at the bottom, type your message
6. Click "Send"
7. View responses in the messages panel

### What to Expect:

- Upon connection, you'll receive a welcome message with your client ID
- Each message you send will be echoed back with a timestamp
- You can see connection status and message history

---

## Testing with Command Line Tools

### Using websocat

```bash
# Install websocat (if not installed)
cargo install websocat

# Connect to the server
websocat ws://localhost:8080/ws

# Type messages and press Enter
```

### Using wscat

```bash
# Install wscat (if not installed)
npm install -g wscat

# Connect to the server
wscat -c ws://localhost:8080/ws

# Type messages and press Enter
```

### Using the provided script

```bash
# Uses websocat or wscat if available
./scripts/test-connection.sh

# Or specify a different host
./scripts/test-connection.sh localhost:8080
```

---

## Testing with Envoy

Test the WebSocket connection through Envoy proxy.

### Start with Envoy:

```bash
# Using Docker Compose
docker-compose --profile envoy up

# Or using Make
make envoy-up
```

### Connect through Envoy:

The Envoy proxy listens on port 10000 and forwards to the WebSocket server.

```bash
# Using Go client through Envoy
go run cmd/client/main.go -addr localhost:10000

# Using websocat through Envoy
websocat ws://localhost:10000/ws

# Using browser
# Open http://localhost:10000 (if Envoy is configured for HTTP)
# Or use Postman with ws://localhost:10000/ws
```

### Verify Envoy is Working:

1. **Check Envoy admin interface:**
   ```bash
   curl http://localhost:9901/stats | grep websocket
   ```

2. **Check health:**
   ```bash
   curl http://localhost:9901/clusters
   ```

3. **View logs:**
   ```bash
   docker-compose logs -f envoy
   ```

### Architecture with Envoy:

```
Client (port 10000) → Envoy → WebSocket Server (port 8080)
```

---

## Health Check

The server provides a health endpoint:

```bash
curl http://localhost:8080/health
```

Expected response:
```json
{"status":"ok"}
```

---

## Expected Behavior

### On Connection:
- Client receives: `Welcome! Your client ID is: <address>-<timestamp>`

### On Message:
- Client sends: `Hello World`
- Server responds: `[HH:MM:SS] Echo: Hello World`

### On Disconnect:
- Connection closes gracefully
- Server logs the disconnection

---

## Troubleshooting

### Connection Refused

**Problem:** Cannot connect to the server

**Solutions:**
- Verify the server is running: `docker-compose ps` or check if process is running
- Check if port 8080 is in use: `lsof -i :8080` (macOS/Linux) or `netstat -ano | findstr :8080` (Windows)
- Try with explicit localhost: `ws://127.0.0.1:8080/ws`

### WebSocket Upgrade Failed

**Problem:** HTTP 400 or upgrade error

**Solutions:**
- Ensure you're using `ws://` protocol, not `http://`
- Check that the path is `/ws`
- Verify the client is sending proper WebSocket upgrade headers

### Envoy Connection Issues

**Problem:** Cannot connect through Envoy

**Solutions:**
- Verify Envoy is running: `docker-compose ps`
- Check Envoy logs: `docker-compose logs envoy`
- Verify server is healthy: `curl http://localhost:8080/health`
- Check Envoy admin interface: `curl http://localhost:9901/stats`

### Message Not Received

**Problem:** Messages sent but not received

**Solutions:**
- Check server logs for errors
- Verify message format (should be text)
- Ensure connection is still active
- Check for firewall or proxy issues

---

## Load Testing

For load testing, you can use tools like:

- **Artillery**: `npm install -g artillery`
- **ws-benchmark**: For WebSocket-specific benchmarking

Example Artillery config (`artillery-test.yml`):

```yaml
config:
  target: "ws://localhost:8080"
  phases:
    - duration: 60
      arrivalRate: 10

scenarios:
  - name: "WebSocket test"
    engine: ws
    flow:
      - send: "Hello from load test"
      - think: 1
```

Run: `artillery run artillery-test.yml`

---

## Automated Testing

You can create automated tests:

```bash
# Create a test file
cat > test_websocket.sh << 'EOF'
#!/bin/bash
echo "Test Message" | websocat ws://localhost:8080/ws -1
EOF

chmod +x test_websocket.sh
./test_websocket.sh
```

---

## Next Steps

After confirming the WebSocket connection works:

1. Test with Envoy proxy enabled
2. Experiment with different message types
3. Test connection stability under load
4. Implement custom message handlers
5. Add authentication/authorization
6. Implement broadcast functionality

---

## Additional Resources

- [WebSocket Protocol RFC](https://tools.ietf.org/html/rfc6455)
- [Gorilla WebSocket Documentation](https://pkg.go.dev/github.com/gorilla/websocket)
- [Envoy WebSocket Support](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/http/upgrades)
