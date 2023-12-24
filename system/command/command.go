package command

import "github.com/df-mc/dragonfly/server/cmd"

func RegisterCmd() {
	cmd.Register(cmd.New("stop", "close Server", nil, StopCommand{}))
	cmd.Register(cmd.New("permission", "set players permission", nil, PermissionCommand{}))
	cmd.Register(cmd.New("tp", "teleport", nil, TpCommand{}, TpPlayerCommand{}))
	cmd.Register(cmd.New("getpos", "get position", nil, GetPosCommand{}))
	cmd.Register(cmd.New("gamemode", "change gamemode", nil, GamemodeCommand{}))
}
