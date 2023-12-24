package packets

import (
	_ "unsafe"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

//go:linkname PlayerSession github.com/df-mc/dragonfly/server/player.(*Player).session
func PlayerSession(*player.Player) *session.Session

//go:linkname WritePacket github.com/df-mc/dragonfly/server/session.(*Session).writePacket
func WritePacket(*session.Session, packet.Packet)
