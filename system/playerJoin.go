package system

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/title"
	"github.com/df-mc/dragonfly/server/world"
)

var events = []func(*player.Player){}

func OnPlayerJoin(ev func(*player.Player)) {
	events = append(events, ev)
}

type entityTypeZombie struct{}

func (entityTypeZombie) EncodeEntity() string {
	return "minecraft:zombie"
}
func (entityTypeZombie) BBox(world.Entity) cube.BBox {
	return cube.Box(-0.3, 0, -0.3, 0.3, 2.9, 0.3)
}

func FirePlayerJoin(player *player.Player) {
	player.SendTitle(title.New("welcome", "good"))

	for i := range events {
		event := events[i]

		event(player)
	}
}
