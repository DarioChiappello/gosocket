package main

import (
	"fmt"
	"log"

	"github.com/DarioChiappello/gosocket"
)

func main() {
	client, err := gosocket.NewClient("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	client.On("message", func(data interface{}) {
		fmt.Printf("New message: %+v\n", data)
	})

	go client.Listen()

	client.Send("message", map[string]interface{}{
		"text": "Hello World",
		"num":  123,
		"list": []string{"one", "two", "three"},
	})

	select {} // app is now running
}
