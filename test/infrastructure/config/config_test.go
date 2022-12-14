package config

import (
	"fmt"
	"testing"

	"github.com/viafcccy/garden-be/global"
	_ "github.com/viafcccy/garden-be/test"
)

func TestNewConfig(t *testing.T) {

	if global.Gconfig.Server.Address == "" {
		t.Fatal("error: 请在 config.yaml 中配置端口号")
	}
	fmt.Println(global.Gconfig.Server.Address)
}
