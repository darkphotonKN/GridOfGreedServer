package main

import (
	"fmt"
	"log"

	"github.com/darkphotonKN/GridOfGreedWsServer/internal/game_server"
)

func main() {
	gs := gameserver.NewGameServer(":6666")

	// initiate message hub to constantly handle incoming messages
	go gs.MessageHub()

	err := gs.Connect()

	if err != nil {
		log.Fatal("ListenAndServe failed, error: ", err)
	}

	fmt.Println("exiting...")
}
