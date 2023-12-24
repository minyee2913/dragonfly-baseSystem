package multiworld

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/minyee2913/dragonfly-baseSystem/system/permission"
	"github.com/minyee2913/dragonfly-baseSystem/system/serv"
)

type TransferWorldCommand struct {
	Id  string
	Pos cmd.Optional[mgl64.Vec3]
}

func (c TransferWorldCommand) Run(source cmd.Source, output *cmd.Output) {
	pos, has := c.Pos.Load()

	if !has {
		pos = source.(*player.Player).Position()
	}

	w := GetWorld(c.Id)

	if c.Id == "main" {
		w = serv.GetServer().World()
	}

	if w == nil {
		return
	}

	TransferWorld(source.(*player.Player), w, pos)
}

func (TransferWorldCommand) Allow(source cmd.Source) bool {
	return permission.IsPerm(source.(*player.Player), 1)
}
