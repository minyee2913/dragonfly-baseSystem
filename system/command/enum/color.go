package enum

import (
	"github.com/df-mc/dragonfly/server/cmd"
)

type Color string

func (Color) Type() string {
	return "color"
}

func (Color) Options(source cmd.Source) []string {
	return []string{"color"}
}
