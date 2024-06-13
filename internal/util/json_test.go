package util_test

import (
	"encoding/json"
	"testing"
)

func TestJSONSerialization(t *testing.T) {
	message := map[string]interface{}{
		"event": "test_event",
		"data": map[string]interface{}{
			"message": "Hello, JSON!",
		},
	}

	data, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	var received map[string]interface{}
	if err := json.Unmarshal(data, &received); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if received["event"] != "test_event" {
		t.Errorf("Expected event to be 'test_event', got '%v'", received["event"])
	}

	dataMap, ok := received["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("Expected data to be a map, got '%T'", received["data"])
	}

	if dataMap["message"] != "Hello, JSON!" {
		t.Errorf("Expected message to be 'Hello, JSON!', got '%v'", dataMap["message"])
	}
}
