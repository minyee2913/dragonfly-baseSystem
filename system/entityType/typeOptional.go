package entityType

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
)

type Optional struct {
	Identifier string
}

func (e Optional) EncodeEntity() string {
	return e.Identifier
}
func (Optional) BBox(world.Entity) cube.BBox {
	return cube.Box(-0.3, 0, -0.3, 0.3, 2, 0.3)
}
