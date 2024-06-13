package connection

import (
	"encoding/json"
	"net"
)

func SendMessage(conn net.Conn, message map[string]interface{}) error {
	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}
	_, err = conn.Write(jsonData)
	return err
}

func ReceiveMessage(conn net.Conn) (map[string]interface{}, error) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	var message map[string]interface{}
	if err := json.Unmarshal(buffer[:n], &message); err != nil {
		return nil, err
	}
	return message, nil
}
