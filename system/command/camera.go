package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/minyee2913/dragonfly-baseSystem/system/camera"
	"github.com/minyee2913/dragonfly-baseSystem/system/command/enum"
	"github.com/minyee2913/dragonfly-baseSystem/system/permission"
	"github.com/minyee2913/dragonfly-baseSystem/system/serv"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

type CameraSetCommand struct {
	Player   enum.PlayerSelector
	Set      enum.Set
	Preset   enum.CameraPreset
	Pos      enum.Pos
	Position mgl64.Vec3
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func (c CameraSetCommand) Run(source cmd.Source, output *cmd.Output) {
	p, _ := serv.GetServer().PlayerByName(string(c.Player))

	pre := indexOf(string(c.Preset), camera.GetPresetNames())
	camera.Set(p, protocol.CameraInstructionSet{
		Preset: uint32(pre),
		Position: protocol.Option(mgl32.Vec3{
			float32(c.Position.X()),
			float32(c.Position.Y()),
			float32(c.Position.Z()),
		}),
	})
}

func (CameraSetCommand) Allow(source cmd.Source) bool {
	return permission.IsPerm(source.(*player.Player), 1)
}

type CameraClearCommand struct {
	Player enum.PlayerSelector
	Clear  enum.Clear
}

func (c CameraClearCommand) Run(source cmd.Source, output *cmd.Output) {
	p, _ := serv.GetServer().PlayerByName(string(c.Player))
	camera.Clear(p)
}

func (CameraClearCommand) Allow(source cmd.Source) bool {
	return permission.IsPerm(source.(*player.Player), 1)
}
