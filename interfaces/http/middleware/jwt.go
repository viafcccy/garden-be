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

		token := c.Query("token")
		// 未携带 token
		if token == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrInvalidParams)
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
	}
}