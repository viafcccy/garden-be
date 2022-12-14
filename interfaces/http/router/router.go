package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/viafcccy/garden-be/cmd"
	"github.com/viafcccy/garden-be/interfaces/http/middleware"
	"github.com/viafcccy/garden-be/interfaces/http/router/user"
)

func Routers(app *cmd.App) *gin.Engine {
	Router := gin.Default()

	// 健康检查
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	//配置跨域
	Router.Use(middleware.Cors())

	ApiGroup := Router.Group("/api/v1")
	user.InitUserRouter(ApiGroup, app) //注入用户模块

	return Router
}
