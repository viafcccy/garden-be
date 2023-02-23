package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dto "github.com/viafcccy/garden-be/application/dto/user"
	service "github.com/viafcccy/garden-be/application/service/user"
	"github.com/viafcccy/garden-be/infrastructure/pkg/response"
)

type UserHandler struct {
	UserSrv *service.UserService
	LogIn   *service.UserService
}

// ApiGetSimpleUser 查询测试用例用户信息 GET /api/v1/user/getSimpleUserInfo
func (u *UserHandler) ApiGetSimpleUser(c *gin.Context) {
	simpleUserReq := &dto.SimpleUserInfoReq{}
	// 处理请求参数
	if err := c.ShouldBindQuery(simpleUserReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrInvalidParams)
		return
	}

	dtoSimpleUserInfo, err := u.UserSrv.GetSimpleUserInfo(c.Request.Context(), simpleUserReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(dtoSimpleUserInfo))

}
