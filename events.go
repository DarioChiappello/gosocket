package gosocket

import (
	"encoding/json"
)

func (s *Socket) sendMessage(message map[string]interface{}) error {
	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}
	_, err = s.conn.Write(jsonData)
	return err
}

func (s *Socket) listen() {
	buffer := make([]byte, 1024)
	for {
		n, err := s.conn.Read(buffer)
		if err != nil {
			return
		}
		message := buffer[:n]
		var received map[string]interface{}
		if err := json.Unmarshal(message, &received); err == nil {
			if event, exists := received["event"].(string); exists {
				s.Emit(event, received["data"])
			}
		}
	}
}
