# Setup Instructions for Local Machine

## Option 1: Copy Files from Container (Current Situation)

The project has been created in the container at `/home/claude/websocket-playground`. You need to copy it to your local machine at `/Users/dthuilot/Documents/Workspace/websocket-playground`.

### If using Docker Desktop or similar:

```bash
# Create the target directory on your Mac
mkdir -p /Users/dthuilot/Documents/Workspace/websocket-playground

# Copy from container to your machine
# Replace <container_id> with the actual container ID
docker cp <container_id>:/home/claude/websocket-playground/. /Users/dthuilot/Documents/Workspace/websocket-playground/
```

### To find the container ID:

```bash
docker ps -a | grep claude
```

---

## Option 2: Download All Files

I can provide all the files for you to recreate manually:

1. Create the directory structure on your Mac:
```bash
mkdir -p /Users/dthuilot/Documents/Workspace/websocket-playground
cd /Users/dthuilot/Documents/Workspace/websocket-playground
```

2. Initialize git:
```bash
git init
git branch -m main
```

3. Copy all the file contents I've created and save them to the appropriate paths

---

## After Copying to Your Local Machine

1. **Verify the files:**
```bash
cd /Users/dthuilot/Documents/Workspace/websocket-playground
ls -la
```

2. **Install Go dependencies:**
```bash
go mod download
```

3. **Test locally:**
```bash
# Run the server
go run cmd/server/main.go

# In another terminal, run the client
go run cmd/client/main.go
```

4. **Test with Docker:**
```bash
docker-compose up --build
```

5. **Test with Envoy:**
```bash
docker-compose --profile envoy up --build
```

---

## Verify Everything Works

### 1. Check health endpoint:
```bash
curl http://localhost:8080/health
```
Should return: `{"status":"ok"}`

### 2. Open browser test page:
```bash
open http://localhost:8080
```

### 3. Test WebSocket connection:
- Click "Connect" in the browser
- Send a test message
- Verify you receive an echo response

### 4. Test with Go client:
```bash
go run cmd/client/main.go
```
Type messages and verify they're echoed back.

---

## Set Up Git Remote (Optional)

If you want to push to GitHub:

```bash
# Create a new repository on GitHub first, then:
git remote add origin https://github.com/dthuilot/websocket-playground.git
git push -u origin main
```

---

## Troubleshooting

### "go: command not found"
Install Go from https://golang.org/dl/

### "docker: command not found"  
Install Docker Desktop from https://www.docker.com/products/docker-desktop

### Port 8080 already in use
```bash
# Find what's using it
lsof -i :8080

# Or change the port
PORT=9000 go run cmd/server/main.go
```

### Permission denied on scripts
```bash
chmod +x scripts/*.sh
```

---

## Quick Validation Checklist

- [ ] All files copied to local machine
- [ ] Git repository initialized
- [ ] Go dependencies downloaded
- [ ] Server runs locally with `go run cmd/server/main.go`
- [ ] Health endpoint returns OK
- [ ] Browser test page works
- [ ] Client can connect
- [ ] Docker build succeeds
- [ ] Docker containers start successfully
- [ ] Envoy proxy works (optional)

---

## What's Next?

1. Read **QUICKSTART.md** for immediate usage
2. Read **TESTING.md** for comprehensive testing guide
3. Read **README.md** for full documentation
4. Check **PROJECT_SUMMARY.md** for overview

Enjoy your WebSocket playground! 🚀
