package gameserver

import (
	"fmt"
	"github.com/darkphotonKN/GridOfGreedWsServer/internal/types"
)

/**
* Primary Game Logic
**/

/**
* Handles all game move logic.
**/
func (g *GameServer) handleGameMove(msg []byte) {
	// init message service for encoding / decoding / writing errors
	msgService := NewMessageService(g.ws)

	decodedMsg, err := msgService.DecodeMessage(msg)
	fmt.Printf("Decoded message: %v\n", decodedMsg)

	if err != nil {
		msgService.WriteErrorMessage(err)
	}

	// deduce game move by casting type from message
	move := types.GameMove(decodedMsg.Type)

	fmt.Println("Game Move after cast:", move)

	switch move {
	case types.ActivateGrid:
		// NOTE: from json decoded numbers are float64
		activateIndex, ok := decodedMsg.Value.(float64)

		if !ok {
			msgService.WriteErrorMessage(fmt.Errorf("Incorrect value sent when attempting to make move \"Activate_Grid\" "))
			break
		}

		g.gridState[int(activateIndex)] = !g.gridState[int(activateIndex)]

		fmt.Println("Grid state after activate:", g.gridState)

		msgService.WriteMessage(g.gridState)

		break

	case types.StartGame:
		msgService.WriteMessage(g.gridState)
		break

	}

}
