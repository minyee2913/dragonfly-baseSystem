package console

import "github.com/df-mc/dragonfly/server"

func Command(serv *server.Server, input string) {
	if input == "/stop" {
		serv.Close()
	} else if input == "/list" {

	}
}
