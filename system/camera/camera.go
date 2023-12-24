package camera

import (
	"image/color"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/minyee2913/dragonfly-baseSystem/system/packets"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

var Presets = []protocol.CameraPreset{
	{
		Name:   "minecraft:free",
		Parent: "",
	},
}

func GetPresetNames() []string {
	presets := []string{}

	for _, preset := range Presets {
		presets = append(presets, preset.Name)
	}

	return presets
}

func UpdatePreset(p *player.Player) {
	packets.WritePacket(packets.PlayerSession(p), &packet.CameraPresets{
		Presets: []protocol.CameraPreset{
			{
				Name:   "minecraft:free",
				Parent: "",
			},
		},
	})
}

func Set(p *player.Player, options protocol.CameraInstructionSet) {
	packets.WritePacket(packets.PlayerSession(p), &packet.CameraInstruction{
		Set: protocol.Option(options),
	})
}

func Clear(p *player.Player) {
	packets.WritePacket(packets.PlayerSession(p), &packet.CameraInstruction{
		Clear: protocol.Option(true),
	})
}

func Fade(p *player.Player, fadeIn float32, wait float32, fadeOut float32, color color.RGBA) {
	packets.WritePacket(packets.PlayerSession(p), &packet.CameraInstruction{
		Fade: protocol.Option(protocol.CameraInstructionFade{
			FadeInDuration:  fadeIn,
			WaitDuration:    wait,
			FadeOutDuration: fadeOut,
			Colour:          color,
		}),
	})
}
