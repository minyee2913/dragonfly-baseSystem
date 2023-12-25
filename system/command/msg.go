package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/minyee2913/dragonfly-baseSystem/system/command/enum"
	"github.com/minyee2913/dragonfly-baseSystem/system/serv"
)

type MsgCommand struct {
	Player  enum.PlayerSelector
	Message cmd.Varargs
}

func (c MsgCommand) Run(source cmd.Source, output *cmd.Output) {
	p, has := serv.GetServer().PlayerByName(string(c.Player))

	if has {
		p.Message(c.Message)
		output.Print("§7" + p.Name() + "님에게 >> " + string(c.Message))
	}
}
