package entityType

import (
	"github.com/bedrock-gophers/living/living"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
)

type DefaultHandler struct {
	living.NopHandler
	e *living.Living
}

func (DefaultHandler) HandleHurt(ctx *event.Context, damage float64, src world.DamageSource) {
}
