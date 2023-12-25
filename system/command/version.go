package command

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"golang.org/x/mod/modfile"
)

type VersionCommand struct {
}

var ignores = []string{
	"github.com/go-gl/mathgl",
	"github.com/google/uuid",
	"github.com/pelletier/go-toml",
	"github.com/sirupsen/logrus",
	"github.com/brentp/intintmap",
	"github.com/cespare/xxhash",
	"github.com/cespare/xxhash/v2",
	"github.com/df-mc/atomic",
	"github.com/go-jose/go-jose/v3",
	"github.com/golang/protobuf",
	"github.com/golang/snappy",
	"github.com/klauspost/compress",
	"github.com/muhammadmuzzammil1998/jsonc",
	"github.com/rogpeppe/go-internal",
	"github.com/segmentio/fasthash",
	"github.com/df-mc/goleveldb",
	"github.com/sandertv/go-raknet",
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

	txt := "§bGolang §fv" + f.Go.Version + "\n§6Minecraft §fv" + protocol.CurrentVersion + "\n\n"

	i := 0
	for _, require := range f.Require {
		if slices.Contains(ignores, require.Mod.Path) {
			continue
		}

		if strings.HasPrefix(require.Mod.Path, "github.com/") {
			txt += "§a" + strings.Replace(require.Mod.Path, "github.com/", "", 1) + "§f " + require.Mod.Version + "\n"
			i++
		}

		if i > 12 {
			txt += "...\n"
			break
		}
	}

	source.(*player.Player).Message(txt + "\n§7Using github.com/minyee2913/dragonfly-baseSystem")
}
