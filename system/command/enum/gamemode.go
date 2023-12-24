package enum

import (
	"github.com/df-mc/dragonfly/server/cmd"
)

type Gamemode string

func (Gamemode) Type() string {
	return "Gamemode"
}

func (Gamemode) Options(source cmd.Source) []string {
	return []string{"creative", "survival", "adventure", "spectator"}
}
