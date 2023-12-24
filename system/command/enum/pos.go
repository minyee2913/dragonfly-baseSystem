package enum

import (
	"github.com/df-mc/dragonfly/server/cmd"
)

type Pos string

func (Pos) Type() string {
	return "pos"
}

func (Pos) Options(source cmd.Source) []string {
	return []string{"pos"}
}
