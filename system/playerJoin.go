package system

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/title"
	"github.com/df-mc/plots/system/permission"
)

var events = []func(*player.Player){}

func OnPlayerJoin(ev func(*player.Player)) {
	events = append(events, ev)
}

func FirePlayerJoin(player *player.Player) {
	player.SendTitle(title.New("welcome", "good"))

	permission.SetPermission(player, 1)

	for i := range events {
		event := events[i]

		event(player)
	}
}
