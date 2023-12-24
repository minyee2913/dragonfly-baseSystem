package serv

import "github.com/df-mc/dragonfly/server"

var defaultServer *server.Server

func SetServer(serv *server.Server) {
	defaultServer = serv
}

func GetServer() *server.Server {
	return defaultServer
}
