package main

import (
	"fmt"
	"os"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/biome"
	"github.com/df-mc/dragonfly/server/world/generator"
	"github.com/df-mc/dragonfly/server/world/mcdb"
	"github.com/minyee2913/dragonfly-baseSystem/system"
	"github.com/minyee2913/dragonfly-baseSystem/system/console"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	conf, err := readConfig(log)
	if err != nil {
		log.Fatalln(err)
	}

	srv := conf.New()
	srv.CloseOnProgramEnd()

	server.New()

	srv.Listen()

	system.SystemMain(srv)

	worldProvider, err := mcdb.Config{Log: log}.Open("worlds/lumiluni")

	conf2, err := readConfig(log)

	confs := world.Config{
		Log:             log,
		Dim:             world.Overworld,
		Provider:        worldProvider,
		Generator:       generator.NewFlat(biome.Plains{}, []world.Block{block.Grass{}, block.Dirt{}, block.Dirt{}, block.Bedrock{}}),
		RandomTickSpeed: conf2.RandomTickSpeed,
		ReadOnly:        conf2.ReadOnlyWorld,
		Entities:        conf2.Entities,
		PortalDestination: func(dim world.Dimension) *world.World {
			return srv.World()
		},
	}
	w := confs.New()

	go console.HandleConsole(srv)

	for srv.Accept(func(p *player.Player) {
		go system.FirePlayerJoin(p)
		w.AddEntity(p)
		w.RemoveEntity(p)

	}) {
	}

	srv.End()
}

// readConfig reads the configuration from the config.toml file, or creates the
// file if it does not yet exist.
func readConfig(log server.Logger) (server.Config, error) {
	c := server.DefaultConfig()
	var zero server.Config
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}
		return zero, nil
	}
	data, err := os.ReadFile("config.toml")
	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}
	return c.Config(log)
}
