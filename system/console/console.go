package console

import (
	"fmt"
	"strings"

	"github.com/df-mc/dragonfly/server"
)

func HandleConsole(serv *server.Server) {
	var input string
	fmt.Scan(&input)

	if strings.HasPrefix(input, "/") {
		Command(serv, input)
	} else {
		players := serv.Players()

		for i := range players {
			players[i].Message("<Server> " + input)
		}
	}

	HandleConsole(serv)
}
