package system

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/minyee2913/dragonfly-baseSystem/system/camera"
	"github.com/minyee2913/dragonfly-baseSystem/system/command"
	"github.com/minyee2913/dragonfly-baseSystem/system/permission"
	"github.com/minyee2913/dragonfly-baseSystem/system/serv"
)

func SystemMain(srv *server.Server) {
	serv.SetServer(srv)
	command.RegisterCmd()

	permission.LoadPermission()

	OnPlayerJoin(func(p *player.Player) {
		camera.UpdatePreset(p)
	})
}
