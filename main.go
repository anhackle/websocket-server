package main

import (
	"fmt"
	"net/http"

	"example.com/websocket/process"
	"github.com/gorilla/websocket"
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection to websocket")
		return
	}

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error receiving message from client")
			return
		}

		if messageType == websocket.BinaryMessage {
			process.ProcessChunk(conn, message)
		}
	}
}

func main() {
	http.HandleFunc("/", handleConnections)
	err := http.ListenAndServeTLS(":8081", "./certs/cert.pem", "./certs/key.pem", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
