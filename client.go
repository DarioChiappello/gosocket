package gosocket

import (
	"net"
)

// Client represents a socket client
type Client struct {
	Socket
}

// NewClient create new client instance
func NewClient(address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Client{
		Socket: Socket{
			conn:      conn,
			listeners: make(map[string][]func(data interface{})),
		},
	}, nil
}

// Send a message
func (c *Client) Send(event string, data interface{}) error {
	message := map[string]interface{}{
		"event": event,
		"data":  data,
	}
	return c.Socket.sendMessage(message)
}
