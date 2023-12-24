package enum

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/minyee2913/dragonfly-baseSystem/system/camera"
)

type CameraPreset string

func (CameraPreset) Type() string {
	return "CameraPreset"
}

func (CameraPreset) Options(source cmd.Source) []string {
	presets := []string{}

	for _, preset := range camera.Presets {
		presets = append(presets, preset.Name)
	}

	return presets
}
