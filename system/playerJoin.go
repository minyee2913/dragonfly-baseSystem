package system

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/title"
)

var events = []func(*player.Player){}

func OnPlayerJoin(ev func(*player.Player)) {
	events = append(events, ev)
}

func FirePlayerJoin(player *player.Player) {
	player.SendTitle(title.New("welcome", "good"))

	for i := range events {
		event := events[i]

		event(player)
	}
}
