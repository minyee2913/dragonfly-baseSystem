package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type GetPosCommand struct {
}

func (c GetPosCommand) Run(source cmd.Source, output *cmd.Output) {
	pos := source.(*player.Player).Position()

	output.Printf("내 좌표: %.1f %.1f %.1f", pos.X(), pos.Y(), pos.Z())
}
