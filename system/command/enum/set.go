package enum

import (
	"github.com/df-mc/dragonfly/server/cmd"
)

type Set string

func (Set) Type() string {
	return "set"
}

func (Set) Options(source cmd.Source) []string {
	return []string{"set"}
}
