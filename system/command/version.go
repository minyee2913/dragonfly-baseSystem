package command

import (
	"fmt"
	"os"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"golang.org/x/mod/modfile"
)

type VersionCommand struct {
}

func (c VersionCommand) Run(source cmd.Source, output *cmd.Output) {
	buf, err := os.ReadFile("go.mod")
	if err != nil {
		fmt.Print(err)
	}

	f, err := modfile.Parse("go.mod", buf, nil)
	if err != nil {
		panic(err)
	}

	txt := "§bGolang §fv" + f.Go.Version + "\n"

	for _, require := range f.Require {
		if require.Mod.Path == "github.com/df-mc/dragonfly" {
			txt += "§adf-mc/dragonfly §f" + require.Mod.Version + "\n"
		}
	}

	source.(*player.Player).Message(txt + "\n§7Using github.com/minyee2913/dragonfly-baseSystem")
}
