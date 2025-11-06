package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dthuilot/websocket-playground/internal/config"
	"github.com/dthuilot/websocket-playground/internal/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	// Setup logger
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	
	// Load configuration
	cfg := config.Load()
	
	// Set log level
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	log.SetLevel(level)
	
	log.WithFields(logrus.Fields{
		"port":             cfg.Port,
		"read_buffer_size":  cfg.ReadBufferSize,
		"write_buffer_size": cfg.WriteBufferSize,
	}).Info("Starting WebSocket server")
	
	// Create WebSocket handler
	wsHandler := handler.NewWebSocketHandler(log, cfg)
	
	// Setup HTTP routes
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", wsHandler.HandleWebSocket)
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/", handleRoot)
	
	// Create server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	// Start server in a goroutine
	go func() {
		log.Infof("Server listening on :%s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()
	
	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	log.Info("Shutting down server...")
	
	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	
	log.Info("Server exited")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>WebSocket Server</title>
</head>
<body>
    <h1>WebSocket Server</h1>
    <p>Server is running. Connect to <code>ws://localhost:8080/ws</code> to establish a WebSocket connection.</p>
    <h2>Quick Test</h2>
    <div>
        <button onclick="connect()">Connect</button>
        <button onclick="disconnect()">Disconnect</button>
        <button onclick="sendMessage()">Send Test Message</button>
    </div>
    <div>
        <input type="text" id="messageInput" placeholder="Enter message" style="width: 300px;" />
    </div>
    <div style="margin-top: 20px;">
        <h3>Status: <span id="status">Disconnected</span></h3>
        <h3>Messages:</h3>
        <pre id="messages" style="background: #f4f4f4; padding: 10px; min-height: 200px;"></pre>
    </div>
    <script>
        let ws;
        const messages = document.getElementById('messages');
        const status = document.getElementById('status');
        const messageInput = document.getElementById('messageInput');

        function connect() {
            ws = new WebSocket('ws://' + window.location.host + '/ws');
            
            ws.onopen = function() {
                status.textContent = 'Connected';
                status.style.color = 'green';
                addMessage('Connected to server');
            };
            
            ws.onmessage = function(event) {
                addMessage('Received: ' + event.data);
            };
            
            ws.onclose = function() {
                status.textContent = 'Disconnected';
                status.style.color = 'red';
                addMessage('Disconnected from server');
            };
            
            ws.onerror = function(error) {
                addMessage('Error: ' + error);
            };
        }

        function disconnect() {
            if (ws) {
                ws.close();
            }
        }

        function sendMessage() {
            if (ws && ws.readyState === WebSocket.OPEN) {
                const message = messageInput.value || 'Hello from browser!';
                ws.send(message);
                addMessage('Sent: ' + message);
                messageInput.value = '';
            } else {
                addMessage('Not connected');
            }
        }

        function addMessage(message) {
            const timestamp = new Date().toLocaleTimeString();
            messages.textContent += timestamp + ' - ' + message + '\n';
            messages.scrollTop = messages.scrollHeight;
        }

        messageInput.addEventListener('keypress', function(e) {
            if (e.key === 'Enter') {
                sendMessage();
            }
        });
    </script>
</body>
</html>
`
	w.Write([]byte(html))
}
