package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/plots/system/serv"
)

type StopCommand struct {
}

// Run will be called when the player runs the command. In this case, we will print the number back to the player
func (c StopCommand) Run(source cmd.Source, output *cmd.Output) {
	go func() {
		serv.GetServer().Close()
	}()
}
