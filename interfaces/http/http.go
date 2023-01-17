package http

import (
	"github.com/viafcccy/garden-be/cmd"
	"github.com/viafcccy/garden-be/global"
	"github.com/viafcccy/garden-be/interfaces/http/router"
)

func InitHttp(app *cmd.App) {
	go func() {
		global.GLog.Infof("garden-be http 服务启动... 端口：%s，服务名：%s，模式：%s",
			global.Gconfig.Server.Address, global.Gconfig.Server.ServerName, global.Gconfig.Server.Mode)
		router.Routers(app).Run(global.Gconfig.Server.Address) // 启动 web
	}()
}
