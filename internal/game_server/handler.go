package gameserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type GameServer struct {
	addr    string
	players map[*websocket.Conn]string
	ws      *websocket.Conn
}

type WebsocketServer interface {
	Connect() error
	HandleConnections(w http.ResponseWriter, r *http.Request)
}

func NewGameServer(addr string) WebsocketServer {
	return &GameServer{addr: addr}
}

func (g *GameServer) Connect() error {
	http.HandleFunc("/ws", g.HandleConnections)

	log.Printf("Server started on %s.", g.addr)

	err := http.ListenAndServe(g.addr, nil)
	if err != nil {
		return err
	}

	return nil
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins TODO: update for production
	},
}

/**
* Handles initial incoming connections.
**/
func (g *GameServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = ws.WriteMessage(websocket.TextMessage, []byte("Successfully connected to the server."))
	if err != nil {
		fmt.Println("Error when sending initial websocket server message.")
	}

	// set websocket connection to the instance
	g.ws = ws

	// start read message loop for connected player on its own go-routine
	go g.HandlePlayer()
}

/**
* Starts a goroutine to handle incoming client messages.
**/
func (g *GameServer) HandlePlayer() {
	defer g.ws.Close()

	for {
		_, msg, err := g.ws.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		fmt.Printf("Received message: %s\n", msg)

		if string(msg) == "startgame" {

			// provide board state to app
			boardState := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false}

			boardStateJson, err := json.Marshal(boardState)

			if err != nil {
				fmt.Println("Error when connverting to json: ", err)
			}

			err = g.ws.WriteMessage(websocket.TextMessage, boardStateJson)
			if err != nil {
				log.Println("Error when providing board to player:", err)
				break
			}
		}

		// Echo message back to the client
		err = g.ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}
