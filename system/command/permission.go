package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/plots/system/permission"
	"github.com/df-mc/plots/system/serv"
)

type PermissionCommand struct {
	Player PlayerEnum
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
