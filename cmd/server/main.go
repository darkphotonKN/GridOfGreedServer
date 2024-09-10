package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	err = ws.WriteMessage(websocket.TextMessage, []byte("Welcome to my Websocket Server meow."))
	if err != nil {
		fmt.Println("Error when sending initial websocket server message.")
	}

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		fmt.Printf("Received message: %s\n", msg)

		// Echo message back to the client
		err = ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)

	log.Println("WebSocket server started on :6666")
	err := http.ListenAndServe(":6666", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
