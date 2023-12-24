package command

import "github.com/df-mc/dragonfly/server/cmd"

func RegisterCmd() {
	cmd.Register(cmd.New("stop", "close Server", []string{}, StopCommand{}))
}
