package main

import (
	"log"

	"github.com/darkphotonKN/GridOfGreedWsServer/internal/game_server"
)

func main() {
	gs := gameserver.NewGameServer(":6666")

	err := gs.Connect()
	if err != nil {
		log.Fatal("ListenAndServe failed, error: ", err)
	}
}
