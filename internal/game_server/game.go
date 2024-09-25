package gameserver

import (
	"encoding/json"
	"fmt"
	"github.com/darkphotonKN/GridOfGreedWsServer/internal/types"
)

/**
* Primary Game Logic
**/

/**
* Handles all game move logic.
**/
func (g *GameServer) handleGameMove(msg []byte) error {
	// init message service for encoding / decoding / writing errors
	msgService := NewMessageService(g.ws)

	// decode message
	decodedMsg, err := msgService.DecodeMessage(msg)

	// deduce game move by casting type from message
	move := types.GameMove(decodedMsg.Type)

	fmt.Println("Game move:", move)

	if err != nil {
		msgService.WriteErrorMessage(err)
	}

	fmt.Printf("Decoded message: %v\n", decodedMsg)

	switch move {
	case types.ActivateGrid:

		activateIndex, ok := decodedMsg.Value.(int)

		fmt.Println("Activating Grid at index:", activateIndex)

		if !ok {
			return fmt.Errorf("Incorrect value sent when attempting to make move \"Activate_Grid\" ")
		}

		g.gridState[activateIndex] = !g.gridState[activateIndex]

		fmt.Println("Grid state after activate:", g.gridState)

		boardStateJson, err := json.Marshal(g.gridState)

		if err != nil {
			fmt.Println("Error when connverting to json: ", err)
			return err
		}

		msgService.WriteMessage(boardStateJson)

	case types.StartGame:
		boardStateJson, err := json.Marshal(g.gridState)

		if err != nil {
			fmt.Println("Error when connverting to json: ", err)
			return err
		}

		msgService.WriteMessage(boardStateJson)

	default:
		return nil
	}

	return nil
}
