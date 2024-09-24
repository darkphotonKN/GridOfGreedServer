package gameserver

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

/**
* Acts as the hub to handle all concurrent websocket messages sent to the server.
**/
func (g *GameServer) MessageHub() {
	// init message service for encoding / decoding / writing errors
	msgService := NewMessageService(g.ws)

	select {
	case msg := <-msgChan:

		// decode message
		decodedMsg, err := msgService.DecodeMessage(msg)

		if err != nil {
			msgService.WriteErrorMessage(err)
		}

		fmt.Printf("Decoded message: %v", decodedMsg)

		fmt.Println("Received message in hub:", string(msg))

		if string(msg) == "startgame" {
			// provide board state to app
			boardStateJson, err := json.Marshal(g.gridState)

			if err != nil {
				fmt.Println("Error when connverting to json: ", err)
				break
			}

			// send message
			err = g.ws.WriteMessage(websocket.TextMessage, boardStateJson)
			if err != nil {
				err := msgService.WriteErrorMessage(err)

				if err != nil {
					fmt.Println("Error writing message:", err)
					break
				}
				break
			}
		}
	}
}
