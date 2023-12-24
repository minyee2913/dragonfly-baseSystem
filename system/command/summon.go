package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/minyee2913/dragonfly-baseSystem/system/permission"
)

type SummonCommand struct {
	identifier string
	pos        mgl64.Vec3
	health     cmd.Optional[int]
	moveSpeed  cmd.Optional[float64]
}

func (c SummonCommand) Run(source cmd.Source, output *cmd.Output) {
	// health := 20
	// moveSpeed := 0.3

	// hp, hasH := c.health.Load()
	// speed, hasS := c.moveSpeed.Load()

	// if hasH {
	// 	health = hp
	// }

	// if hasS {
	// 	moveSpeed = speed
	// }

	// mob := living.NewLivingEntity(
	// 	entityType.Optional{
	// 		Identifier: c.identifier,
	// 	},
	// 	float64(health),
	// 	moveSpeed,
	// 	nil,
	// )
}

func (SummonCommand) Allow(source cmd.Source) bool {
	return permission.IsPerm(source.(*player.Player), 1)
}
