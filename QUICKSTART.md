# Quick Start Guide

Get the WebSocket Playground up and running in 5 minutes!

## Prerequisites

Choose one of the following options:

### Option A: Run Locally (Requires Go)
- Go 1.21 or higher

### Option B: Run with Docker (Recommended)
- Docker
- Docker Compose

## Quick Start Options

### 🚀 Fastest: Using Docker

```bash
# 1. Clone or navigate to the project
cd /path/to/websocket-playground

# 2. Start the server
docker-compose up --build

# 3. Open browser
open http://localhost:8080
```

That's it! You now have a WebSocket server running.

---

### 🔧 Alternative: Run Locally with Go

```bash
# 1. Clone or navigate to the project
cd /path/to/websocket-playground

# 2. Install dependencies
go mod download

# 3. Run the server
go run cmd/server/main.go

# 4. Open browser
open http://localhost:8080
```

---

## Testing the Connection

### Method 1: Browser (Easiest)

1. Open `http://localhost:8080` in your browser
2. Click the "Connect" button
3. Type a message and click "Send Test Message"
4. Watch messages appear in real-time!

### Method 2: Go Client

```bash
# In a new terminal
go run cmd/client/main.go
```

Type messages and see them echoed back by the server.

### Method 3: Postman

1. Open Postman
2. New → WebSocket Request
3. URL: `ws://localhost:8080/ws`
4. Click "Connect"
5. Send messages!

---

## Testing with Envoy Proxy

Want to test with Envoy between client and server?

```bash
# Start everything including Envoy
docker-compose --profile envoy up --build
```

Now connect to:
- Direct: `ws://localhost:8080/ws` (bypasses Envoy)
- Via Envoy: `ws://localhost:10000/ws` (goes through Envoy)

You can test both and verify the connection works seamlessly through Envoy!

---

## Verifying Everything Works

### Check Health Endpoint

```bash
curl http://localhost:8080/health
# Expected: {"status":"ok"}
```

### Check Envoy Admin (if using Envoy profile)

```bash
curl http://localhost:9901/stats | grep websocket
```

### View Logs

```bash
# Docker logs
docker-compose logs -f

# Just server logs
docker-compose logs -f websocket-server

# Just Envoy logs (if using Envoy profile)
docker-compose logs -f envoy
```

---

## Common Commands

```bash
# Start services
docker-compose up -d

# Stop services
docker-compose down

# Rebuild and start
docker-compose up --build

# View logs
docker-compose logs -f

# Check status
docker-compose ps
```

Or use the Makefile:

```bash
make server          # Run server locally
make client          # Run client locally
make docker-run      # Run with Docker
make docker-stop     # Stop Docker services
make envoy-up        # Run with Envoy
make help            # See all commands
```

---

## What's Next?

1. ✅ **Test basic connection** - Use browser or Go client
2. ✅ **Test with Envoy** - Verify proxy works transparently
3. 🔍 **Explore the code** - Check out `internal/handler/websocket.go`
4. 🛠️ **Customize** - Modify message handling, add features
5. 📚 **Read full docs** - Check `README.md` and `TESTING.md`

---

## Troubleshooting

### Port 8080 already in use?

```bash
# Change the port
PORT=9000 go run cmd/server/main.go

# Or in Docker
# Edit docker-compose.yml and change "8080:8080" to "9000:8080"
```

### Docker build fails?

```bash
# Clean up and rebuild
docker-compose down
docker system prune -f
docker-compose up --build
```

### Can't connect?

1. Make sure server is running: `curl http://localhost:8080/health`
2. Check logs: `docker-compose logs`
3. Try `127.0.0.1` instead of `localhost`
4. Check firewall settings

---

## Project Structure

```
websocket-playground/
├── cmd/
│   ├── server/          # WebSocket server
│   └── client/          # Test client
├── internal/
│   ├── handler/         # WebSocket handlers
│   └── config/          # Configuration
├── deployments/
│   ├── docker/          # Dockerfiles
│   └── envoy/           # Envoy config
├── docker-compose.yml   # Docker orchestration
├── Makefile            # Convenient commands
└── README.md           # Full documentation
```

---

## Success Checklist

- [ ] Server starts without errors
- [ ] Health endpoint returns `{"status":"ok"}`
- [ ] Browser test page loads
- [ ] Can connect and send messages via browser
- [ ] Go client can connect and exchange messages
- [ ] (Optional) Envoy proxy routes WebSocket traffic

---

## Need Help?

- 📖 Full documentation: See `README.md`
- 🧪 Testing guide: See `TESTING.md`
- 🤝 Contributing: See `CONTRIBUTING.md`
- 🐛 Issues: Open an issue on GitHub

---

**Congratulations! 🎉** You now have a working WebSocket server with Envoy proxy support!
