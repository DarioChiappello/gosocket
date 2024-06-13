package main

import (
	"log"

	"github.com/DarioChiappello/gosocket"
)

func main() {
	server, err := gosocket.NewServer("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	server.On("message", func(data interface{}) {
		log.Printf("New client message: %+v\n", data)
	})

	go server.AcceptConnections()

	select {} // app is now running
}
