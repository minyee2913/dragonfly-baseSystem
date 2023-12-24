package multiworld

import (
	"fmt"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/biome"
	"github.com/df-mc/dragonfly/server/world/generator"
	"github.com/df-mc/dragonfly/server/world/mcdb"
	"github.com/sirupsen/logrus"
)

var Worlds = make(map[string](*world.World))

func GetWorld(id string) *world.World {
	return Worlds[id]
}

func LoadWorld(id string, path string, log *logrus.Logger) *world.World {

	worldProvider, err := mcdb.Config{Log: log}.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	conf2, err := server.DefaultConfig().Config(log)

	confs := world.Config{
		Log:               log,
		Dim:               world.Overworld,
		Provider:          worldProvider,
		Generator:         generator.NewFlat(biome.Plains{}, []world.Block{block.Grass{}, block.Dirt{}, block.Dirt{}, block.Bedrock{}}),
		RandomTickSpeed:   conf2.RandomTickSpeed,
		ReadOnly:          conf2.ReadOnlyWorld,
		Entities:          conf2.Entities,
		PortalDestination: nil,
	}

	w := confs.New()

	Worlds[id] = w

	return w
}
