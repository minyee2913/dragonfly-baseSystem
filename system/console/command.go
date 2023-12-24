package console

import (
	"fmt"
	"strconv"

	"github.com/df-mc/dragonfly/server"
)

func Command(serv *server.Server, input string) {
	if input == "/stop" {
		serv.Close()
	} else if input == "/list" {
		fmt.Println("players: " + strconv.Itoa(len(serv.Players())) + "/" + strconv.Itoa(serv.MaxPlayerCount()))
	}
}
