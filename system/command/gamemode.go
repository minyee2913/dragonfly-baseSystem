package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/minyee2913/dragonfly-baseSystem/system/command/enum"
	"github.com/minyee2913/dragonfly-baseSystem/system/permission"
	"github.com/minyee2913/dragonfly-baseSystem/system/serv"
)

type GamemodeCommand struct {
	GameMode enum.Gamemode
	Player   cmd.Optional[enum.PlayerSelector]
}

func (c GamemodeCommand) Run(source cmd.Source, output *cmd.Output) {
	selector, has := c.Player.Load()

	var mode world.GameMode

	switch c.GameMode {
	case "creative":
		mode = world.GameModeCreative
		break
	case "adventure":
		mode = world.GameModeAdventure
		break
	case "spectator":
		mode = world.GameModeSpectator
		break
	default:
		mode = world.GameModeSurvival
	}

	if has {
		player, _ := serv.GetServer().PlayerByName(string(selector))

		player.SetGameMode(mode)

		output.Printf("%s님의 게임모드를 %s(으)로 변경하였습니다", player.Name(), c.GameMode)
	} else {
		source.(*player.Player).SetGameMode(mode)

		output.Printf("내 게임모드를 %s(으)로 변경하였습니다", c.GameMode)
	}
}

func (GamemodeCommand) Allow(source cmd.Source) bool {
	return permission.IsPerm(source.(*player.Player), 1)
}
