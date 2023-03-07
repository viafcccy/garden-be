package user

import (
	"github.com/gin-gonic/gin"

	"github.com/viafcccy/garden-be/cmd"
	handler "github.com/viafcccy/garden-be/interfaces/http/handler/user"
)

func InitUserRouter(Router *gin.RouterGroup, app *cmd.App) {
	hand := handler.UserHandler{
		UserSrv: app.UserSrv,
	}
	UserRouter := Router.Group("/user")
	{
		// 用户信息
		UserRouter.GET("/getSimpleUserInfo", hand.ApiGetSimpleUser) // 测试接口：查找 id=1 的用户

		// 登录
		UserRouter.POST("/login", hand.ApiLogin) // 用户登录
	}
}
