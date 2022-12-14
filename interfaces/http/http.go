package http

import (
	"github.com/viafcccy/garden-be/cmd"
	"github.com/viafcccy/garden-be/global"
	"github.com/viafcccy/garden-be/interfaces/http/router"
)

func InitHttp(app *cmd.App) {
	go func() {
		router.Routers(app).Run(global.Gconfig.Server.Address) // 启动web
	}()
}
