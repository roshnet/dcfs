package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// initEndpoints registers the required endpoints
func initEndpoints() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ws", wsHandler)
}

func main() {

	// Register endpoints
	initEndpoints()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
