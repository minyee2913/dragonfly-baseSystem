package multiworld

import (
	"time"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/minyee2913/dragonfly-baseSystem/system/serv"
)

func TransferWorld(p *player.Player, w *world.World, pos mgl64.Vec3) {
	go func() {
		p.World().RemoveEntity(p)
		serv.GetServer().Nether().AddEntity(p)

		time.Sleep(1200 * time.Millisecond)

		p.World().RemoveEntity(p)
		w.AddEntity(p)

		p.Teleport(pos)
	}()
}
