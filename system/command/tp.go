package command

import (
	"time"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/minyee2913/dragonfly-baseSystem/system/permission"
	"github.com/minyee2913/dragonfly-baseSystem/system/serv"
)

type TpCommand struct {
	Pos mgl64.Vec3
}

func (c TpCommand) Run(source cmd.Source, output *cmd.Output) {
	//source.(*player.Player).Teleport(c.Pos)

	go func() {
		time.Sleep(50 * time.Millisecond)
		serv.GetServer().Nether().AddEntity(source.(*player.Player))
		time.Sleep(3 * time.Second)
		serv.GetServer().Nether().RemoveEntity(source.(*player.Player))
		time.Sleep(50 * time.Millisecond)
		serv.GetServer().World().AddEntity(source.(*player.Player))
	}()

	//output.Printf("%.1f %.1f %.1f (으)로 순간이동 하였습니다", c.Pos.X(), c.Pos.Y(), c.Pos.Z())
}

func (TpCommand) Allow(source cmd.Source) bool {
	return permission.IsPerm(source.(*player.Player), 1)
}
