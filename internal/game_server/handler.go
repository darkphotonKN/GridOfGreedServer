package gameserver

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type GameServer struct {
	addr    string
	players map[*websocket.Conn]string
}

type WebsocketServer interface {
	Connect() error
	HandleConnections(w http.ResponseWriter, r *http.Request)
}

func NewGameServer(addr string) WebsocketServer {
	return &GameServer{addr: addr}
}

func (g *GameServer) Connect() error {
	log.Printf("Server started on %s.", g.addr)

	err := http.ListenAndServe(g.addr, nil)
	if err != nil {
		return err
	}

	http.HandleFunc("/ws", g.HandleConnections)
	return nil
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins TODO: update for production
	},
}

func (g *GameServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	err = ws.WriteMessage(websocket.TextMessage, []byte("Successfully connected to the server."))
	if err != nil {
		fmt.Println("Error when sending initial websocket server message.")
	}

	// start read message loop on its own go-routine
}

/**
* Starts a goroutine to handle incoming client messages.
**/
