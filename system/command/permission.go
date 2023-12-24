package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/minyee2913/dragonfly-baseSystem/system/command/enum"
	"github.com/minyee2913/dragonfly-baseSystem/system/permission"
	"github.com/minyee2913/dragonfly-baseSystem/system/serv"
)

type PermissionCommand struct {
	Player enum.PlayerSelector
	Level  int
}

func (c PermissionCommand) Run(source cmd.Source, output *cmd.Output) {
	player, _ := serv.GetServer().PlayerByName(string(c.Player))

	if player != nil {
		permission.SetPermission(player, c.Level)
	}
}

func (PermissionCommand) Allow(source cmd.Source) bool {
	return permission.IsPerm(source.(*player.Player), 1)
}
