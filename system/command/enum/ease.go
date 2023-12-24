package enum

import (
	"github.com/df-mc/dragonfly/server/cmd"
)

type Ease string

func (Ease) Type() string {
	return "ease"
}

func (Ease) Options(source cmd.Source) []string {
	return []string{"ease"}
}
