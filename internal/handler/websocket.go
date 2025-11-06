package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dthuilot/websocket-playground/internal/config"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

const (
	// Time allowed to write a message to the peer
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer
	pongWait = 60 * time.Second

	// Send pings to peer with this period (must be less than pongWait)
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer
	maxMessageSize = 512
)

// WebSocketHandler handles WebSocket connections
type WebSocketHandler struct {
	log      *logrus.Logger
	upgrader websocket.Upgrader
}

// NewWebSocketHandler creates a new WebSocket handler
func NewWebSocketHandler(log *logrus.Logger, cfg *config.Config) *WebSocketHandler {
	return &WebSocketHandler{
		log: log,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  cfg.ReadBufferSize,
			WriteBufferSize: cfg.WriteBufferSize,
			CheckOrigin: func(r *http.Request) bool {
				// Allow all origins for development
				// In production, implement proper origin checking
				return true
			},
		},
	}
}

// HandleWebSocket handles WebSocket connection requests
func (h *WebSocketHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.log.WithError(err).Error("Failed to upgrade connection")
		return
	}

	clientID := fmt.Sprintf("%s-%d", r.RemoteAddr, time.Now().Unix())
	
	h.log.WithFields(logrus.Fields{
		"client_id":   clientID,
		"remote_addr": r.RemoteAddr,
	}).Info("New WebSocket connection")

	// Create client
	client := &Client{
		conn:     conn,
		send:     make(chan []byte, 256),
		log:      h.log,
		clientID: clientID,
	}

	// Start client goroutines
	go client.writePump()
	go client.readPump()

	// Send welcome message
	welcomeMsg := fmt.Sprintf("Welcome! Your client ID is: %s", clientID)
	client.send <- []byte(welcomeMsg)
}

// Client represents a WebSocket client
type Client struct {
	conn     *websocket.Conn
	send     chan []byte
	log      *logrus.Logger
	clientID string
}

// readPump pumps messages from the WebSocket connection to the hub
func (c *Client) readPump() {
	defer func() {
		c.conn.Close()
		c.log.WithField("client_id", c.clientID).Info("Client disconnected")
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.log.WithError(err).WithField("client_id", c.clientID).Error("Unexpected close error")
			}
			break
		}

		c.log.WithFields(logrus.Fields{
			"client_id":    c.clientID,
			"message_type": messageType,
			"message":      string(message),
		}).Info("Received message")

		// Echo the message back with a timestamp
		response := fmt.Sprintf("[%s] Echo: %s", time.Now().Format("15:04:05"), string(message))
		c.send <- []byte(response)
	}
}

// writePump pumps messages from the hub to the WebSocket connection
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued messages to the current WebSocket message
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}

			c.log.WithFields(logrus.Fields{
				"client_id": c.clientID,
				"message":   string(message),
			}).Debug("Sent message")

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
