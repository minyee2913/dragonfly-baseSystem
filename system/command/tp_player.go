package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/plots/system/permission"
	"github.com/df-mc/plots/system/serv"
)

type TpPlayerCommand struct {
	Player PlayerEnum
}

func (c TpPlayerCommand) Run(source cmd.Source, output *cmd.Output) {
	p, _ := serv.GetServer().PlayerByName(string(c.Player))
	pos := p.Position()

	source.(*player.Player).Teleport(pos)

	output.Printf("%s 님께 순간이동 하였습니다", c.Player)
}

func (TpPlayerCommand) Allow(source cmd.Source) bool {
	return permission.IsPerm(source.(*player.Player), 1)
}
