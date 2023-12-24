package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/minyee2913/dragonfly-baseSystem/system/permission"
	"github.com/minyee2913/dragonfly-baseSystem/system/serv"
)

type StopCommand struct {
}

func (c StopCommand) Run(source cmd.Source, output *cmd.Output) {
	go func() {
		serv.GetServer().Close()
	}()
}

func (StopCommand) Allow(source cmd.Source) bool {
	return permission.IsPerm(source.(*player.Player), 1)
}
