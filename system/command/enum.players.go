package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/plots/system/serv"
)

type PlayerEnum string

func (PlayerEnum) Type() string {
	return "Players"
}

func (PlayerEnum) Options(source cmd.Source) []string {
	playerNames := []string{}

	players := serv.GetServer().Players()
	for i := range players {
		playerNames = append(playerNames, players[i].Name())
	}

	return playerNames
}
