package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/minyee2913/dragonfly-baseSystem/system/command/enum"
	"github.com/minyee2913/dragonfly-baseSystem/system/permission"
	"github.com/minyee2913/dragonfly-baseSystem/system/serv"
)

type KickCommand struct {
	Player  enum.PlayerSelector
	Message cmd.Optional[cmd.Varargs]
}

func (c KickCommand) Run(source cmd.Source, output *cmd.Output) {
	p, has := serv.GetServer().PlayerByName(string(c.Player))

	message := c.Message.LoadOr("You are kicked!")

	if has {
		p.Disconnect(message)
		chat.Global.WriteString(p.Name() + "is kicked by " + source.(*player.Player).Name())
	}
}

func (KickCommand) Allow(source cmd.Source) bool {
	return permission.IsPerm(source.(*player.Player), 1)
}
