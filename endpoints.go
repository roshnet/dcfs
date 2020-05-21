package main

import (
	"fmt"
	"log"
	"net/http"
)

// rootHandler handles the default route
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have reached /")
}

// wsHandler handles the WebSocket connections
func wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(">>>Error in Upgrader<<<", err)
	}

	// The client has connected, send an ack
	err = ws.WriteMessage(1, []byte("Hello client, I am the server."))
	if err != nil {
		log.Println(">>>Error in sending message to client<<<", err)
	}

	reader(ws)
}
