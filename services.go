package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// reader handles the incoming messages from the client
func reader(conn *websocket.Conn) {
	for {
		msgType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("[INCOMING] ", string(p))

		err = conn.WriteMessage(msgType, p)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
