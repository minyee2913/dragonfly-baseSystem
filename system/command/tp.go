package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/plots/system/permission"
	"github.com/go-gl/mathgl/mgl64"
)

type TpCommand struct {
	Pos mgl64.Vec3
}

func (c TpCommand) Run(source cmd.Source, output *cmd.Output) {
	source.(*player.Player).Teleport(c.Pos)

	output.Printf("%.1f %.1f %.1f (으)로 순간이동 하였습니다", c.Pos.X(), c.Pos.Y(), c.Pos.Z())
}

func (TpCommand) Allow(source cmd.Source) bool {
	return permission.IsPerm(source.(*player.Player), 1)
}
