package gosocket

import (
	"net"
	"sync"
)

// Socket main structure
type Socket struct {
	conn      net.Conn
	listeners map[string][]func(data interface{})
	mu        sync.Mutex
}

// On register a listener event for a specific event
func (s *Socket) On(event string, listener func(data interface{})) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.listeners[event]; !exists {
		s.listeners[event] = []func(data interface{}){}
	}
	s.listeners[event] = append(s.listeners[event], listener)
}

// Emit an event to all the registred listeners
func (s *Socket) Emit(event string, data interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if listeners, exists := s.listeners[event]; exists {
		for _, listener := range listeners {
			go listener(data)
		}
	}
}
