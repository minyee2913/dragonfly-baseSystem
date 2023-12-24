package multiworld

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/minyee2913/dragonfly-baseSystem/system/permission"
	"github.com/minyee2913/dragonfly-baseSystem/system/serv"
)

type WorldListCommand struct{}

func (c WorldListCommand) Run(source cmd.Source, output *cmd.Output) {
	var txt = "-- world list --\n- " + serv.GetServer().World().Name() + " §amain" + "\n"
	for key, world := range Worlds {
		txt += "§f- " + world.Name() + " §7" + key + "\n"
	}

	output.Print(txt)
}

func (WorldListCommand) Allow(source cmd.Source) bool {
	return permission.IsPerm(source.(*player.Player), 1)
}
