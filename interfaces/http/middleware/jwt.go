package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	myJwt "github.com/viafcccy/garden-be/infrastructure/pkg/jwt"
	"github.com/viafcccy/garden-be/infrastructure/pkg/response"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {

		// login api
		if c.Request.URL.Path == "/api/v1/user/login" {
			// 前置 进入下一个中间件
			c.Next()
			// 后置 再次回来时直接销毁函数
			return
		}

		// 其它 api
		// 前置
		token := c.Request.Header.Get("token")
		// 未携带 token
		if token == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, response.ErrNotCarryToken)
		} else {
			claims, err := myJwt.ParseJwtToken(token)
			// 验证失败
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrAuthCheckTokenFail)
				// token 超时
			} else if jwt.NewNumericDate(time.Now()).Unix() > claims.ExpiresAt.Unix() {
				c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrAuthCheckTokenTimeout)
			}
		}

		c.Next()
		// 后置 ...
	}
}
