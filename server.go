package gosocket

import (
	"net"
)

// Server represents a socket server
type Server struct {
	Socket
	listener net.Listener
}

// NewServer create a new server instance
func NewServer(address string) (*Server, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Server{
		Socket: Socket{
			listeners: make(map[string][]func(data interface{})),
		},
		listener: listener,
	}, nil
}

// AcceptConnections acept entries connections
func (s *Server) AcceptConnections() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			continue
		}
		go s.handleConnection(conn)
	}
}

// handleConnection handle a connection
func (s *Server) handleConnection(conn net.Conn) {
	clientSocket := Socket{
		conn:      conn,
		listeners: s.listeners,
	}
	go clientSocket.listen()
}
