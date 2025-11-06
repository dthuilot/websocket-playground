package main

import (
	"bufio"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var (
	addr = flag.String("addr", "localhost:8080", "WebSocket server address")
	path = flag.String("path", "/ws", "WebSocket path")
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: *path}
	log.Printf("Connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	// Goroutine to read messages from the server
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("Received: %s", message)
		}
	}()

	// Goroutine to read messages from stdin and send to server
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		log.Println("Type messages to send (or Ctrl+C to quit):")
		for scanner.Scan() {
			message := scanner.Text()
			if message == "" {
				continue
			}

			err := c.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("write:", err)
				return
			}
			log.Printf("Sent: %s", message)
		}
	}()

	// Send initial message
	err = c.WriteMessage(websocket.TextMessage, []byte("Hello from Go client!"))
	if err != nil {
		log.Println("write:", err)
		return
	}

	// Wait for interrupt or done signal
	select {
	case <-done:
		log.Println("Connection closed by server")
	case <-interrupt:
		log.Println("Interrupt received, closing connection...")

		// Send close message to server
		err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Println("write close:", err)
			return
		}

		// Wait for close acknowledgment or timeout
		select {
		case <-done:
		case <-time.After(time.Second):
		}
	}
}
