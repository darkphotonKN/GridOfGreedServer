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
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type GameMessageResponse struct {
	Data interface{} `json:"data"`
}

/**
* Decodes message received as bytes from the websocket connection.
**/
func (ms *MessageService) DecodeMessage(msg []byte) (GameMessage, error) {
	fmt.Println("decoding msg...", msg)

	var jsonMsg GameMessage

	err := json.Unmarshal(msg, &jsonMsg)
	if err != nil {
		return GameMessage{}, err
	}

	fmt.Printf("Decoded msg: %+v", jsonMsg)

	return jsonMsg, nil
}

/**
* Write error message in pre-defined format back to connected user.
**/
func (ms *MessageService) WriteErrorMessage(err error) error {
	// exception, send back error
	errMsg, err := json.Marshal(err.Error())

	if err != nil {
		fmt.Println("Error when marshalling json:", err)
		// return that marshalling the error message itself errored
		return err
	}

	fmt.Println("Writing error back to user:", err)

	err = ms.ws.WriteMessage(websocket.TextMessage, errMsg)

	return nil
}

/**
* Writes message to user and handles any errors.
**/

func (ms *MessageService) WriteMessage(msg interface{}) {
	gameMsgRes := GameMessageResponse{
		Data: msg, // Keep msg as interface{}, it will handle arrays properly
	}

	// Marshal the response struct, including the array of booleans
	jsonMsg, err := json.Marshal(gameMsgRes)

	if err != nil {
		ms.WriteErrorMessage(fmt.Errorf("Error when marshaling response to JSON: %s", err))
		return
	}

	err = ms.ws.WriteMessage(websocket.TextMessage, jsonMsg)

	if err != nil {
		ms.WriteErrorMessage(err)
	}
}
