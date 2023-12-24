package system

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/plots/system/command"
	"github.com/df-mc/plots/system/permission"
	"github.com/df-mc/plots/system/serv"
)

func SystemMain(srv *server.Server) {
	serv.SetServer(srv)
	command.RegisterCmd()

	permission.LoadPermission()
}
