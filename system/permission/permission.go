package permission

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Permission struct {
	name      string
	xuid      string
	permLevel int
}

var permissions []Permission

func LoadPermission() {
	data, err := os.Open("./config/permission.json")

	if err != nil {
		fmt.Println(err)
	}

	byteVal, _ := io.ReadAll(data)

	json.Unmarshal(byteVal, &permissions)

	fmt.Println(permissions)
}
