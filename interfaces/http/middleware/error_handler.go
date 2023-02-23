package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viafcccy/garden-be/infrastructure/pkg/response"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	// for _, err := range c.Errors {
	// 	c.JSON(http.StatusInternalServerError, err)
	// 	return
	// }
	// 兜底：context 错误队列中还有未处理的错误，返回通用错误
	if len(c.Errors) > 0 {
		c.JSON(http.StatusInternalServerError, response.Err)
	}

}
