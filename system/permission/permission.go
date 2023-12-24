package permission

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/df-mc/dragonfly/server/player"
)

type Permission struct {
	Name      string
	Xuid      string
	PermLevel int //0 = member, 1 = operator
}

var permissions = make([]*Permission, 1)

var filePath = "./config/permission.json"

func FindByXuid(xuid string) *Permission {
	var val *Permission = nil

	for i := range permissions {
		permission := permissions[i]

		if permission.Xuid == xuid {
			val = permission
		}
	}

	return val
}

func IsPerm(p *player.Player, level int) bool {
	var perm = FindByXuid(p.XUID())

	if perm == nil {
		return (level == 0)
	}

	return level == perm.PermLevel
}

func SetPermission(p *player.Player, level int) {
	LoadPermission()

	var perm = FindByXuid(p.XUID())

	if perm == nil {
		data := Permission{
			Name:      p.Name(),
			Xuid:      p.XUID(),
			PermLevel: level,
		}

		permissions = append(permissions, &data)
	} else {
		perm.PermLevel = level
	}

	SavePermission()
}

func LoadPermission() {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	data, err := io.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &permissions)

	if err != nil {
		fmt.Println(err)
	}
}

func SavePermission() {
	data, err := json.Marshal(permissions)

	if err != nil {
		fmt.Println(err)
	}

	os.WriteFile(filePath, data, os.FileMode(0644))
}
