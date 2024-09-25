package gameserver

/**
* Acts as the hub to handle all concurrent websocket messages sent to the server.
**/
func (g *GameServer) MessageHub() {

	select {
	case msg := <-msgChan:
		g.handleGameMove(msg)

	}
}
