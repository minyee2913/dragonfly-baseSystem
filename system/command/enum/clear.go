package enum

import (
	"github.com/df-mc/dragonfly/server/cmd"
)

type Clear string

func (Clear) Type() string {
	return "clear"
}

func (Clear) Options(source cmd.Source) []string {
	return []string{"clear"}
}
