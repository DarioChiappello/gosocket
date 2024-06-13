package connection_test

import (
	"testing"
	"time"

	"github.com/DarioChiappello/gosocket"
)

func TestServerClientConnection(t *testing.T) {
	server, err := gosocket.NewServer("localhost:8080")
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	go server.AcceptConnections()

	client, err := gosocket.NewClient("localhost:8080")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}

	messageReceived := make(chan bool)

	server.On("test_event", func(data interface{}) {
		messageReceived <- true
	})

	client.Send("test_event", map[string]interface{}{
		"message": "Hello, Server!",
	})

	select {
	case <-messageReceived:
		// Test passed
	case <-time.After(5 * time.Second):
		t.Fatal("Did not receive message within timeout")
	}
}
