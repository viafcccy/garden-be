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
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("getSimpleUserInfo", hand.ApiGetSimpleUser) // 查找用户
	}
}
