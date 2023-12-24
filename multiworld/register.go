package multiworld

import (
	"fmt"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

type MultiWorldHandler struct {
	world.NopHandler
}

func (MultiWorldHandler) HandleClose() {
	fmt.Println("close multiworld")
	for _, val := range Worlds {
		val.Close()
	}
}

func Register(srv *server.Server) {
	cmd.Register(cmd.New("transferworld", "change world", nil, TransferWorldCommand{}))
	cmd.Register(cmd.New("worldlist", "worlds list", nil, WorldListCommand{}))

	srv.World().Handle(MultiWorldHandler{})
}
