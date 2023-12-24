package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/minyee2913/dragonfly-baseSystem/system/command/enum"
	"github.com/minyee2913/dragonfly-baseSystem/system/permission"
	"github.com/minyee2913/dragonfly-baseSystem/system/serv"
)

type TpPlayerCommand struct {
	Player enum.PlayerSelector
	target cmd.Optional[enum.PlayerSelector]
}

func (c TpPlayerCommand) Run(source cmd.Source, output *cmd.Output) {
	p, _ := serv.GetServer().PlayerByName(string(c.Player))
	pos := p.Position()

	target, hasTarget := c.target.Load()

	if hasTarget {
		p2, _ := serv.GetServer().PlayerByName(string(target))

		source.(*player.Player).Teleport(p2.Position())

		output.Printf("%s 님을 %s님께 순간이동 시켰습니다", c.Player, p2.Name())
	} else {
		source.(*player.Player).Teleport(pos)

		output.Printf("%s 님께 순간이동 하였습니다", c.Player)
	}
}

func (TpPlayerCommand) Allow(source cmd.Source) bool {
	return permission.IsPerm(source.(*player.Player), 1)
}
