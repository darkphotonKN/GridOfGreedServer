package gameserver

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type MessageService struct {
	ws *websocket.Conn
}

func NewMessageService(ws *websocket.Conn) *MessageService {
	return &MessageService{
		ws: ws,
	}
}

type GameMessage struct {
	Type  string `json:"type"`
	Value any    `json:"value"`
}

/**
* Decodes message received as bytes from the websocket connection.
**/
func (ms *MessageService) DecodeMessage(msg []byte) (GameMessage, error) {
	fmt.Println("msg:", msg)
	return GameMessage{}, nil
}

/**
* Write error message in pre-defined format back to connected user.
**/
func (ms *MessageService) WriteErrorMessage(err error) error {
	fmt.Println("Error when providing board to player:", err)

	// exception, send back error
	errMsg, err := json.Marshal("Error when attempting to send back game state.")

	if err != nil {
		return err
	}

	err = ms.ws.WriteMessage(websocket.TextMessage, errMsg)

	return nil
}
