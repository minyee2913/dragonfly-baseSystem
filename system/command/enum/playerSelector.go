package enum

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/minyee2913/dragonfly-baseSystem/system/serv"
)

type PlayerSelector string

func (PlayerSelector) Type() string {
	return "Players"
}

func (PlayerSelector) Options(source cmd.Source) []string {
	playerNames := []string{}

	players := serv.GetServer().Players()
	for i := range players {
		playerNames = append(playerNames, players[i].Name())
	}

	return playerNames
}
