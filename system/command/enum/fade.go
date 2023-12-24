package enum

import (
	"github.com/df-mc/dragonfly/server/cmd"
)

type Fade string

func (Fade) Type() string {
	return "fade"
}

func (Fade) Options(source cmd.Source) []string {
	return []string{"fade"}
}
