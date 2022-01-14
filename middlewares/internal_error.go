package middlewares

import (
	"github.com/gin-gonic/gin"
	"webtest/core/response"
)

func InternalError(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			response.APIResponse(c, -1, "未知错误", gin.H{})
		}
		c.Abort()
	}()
	c.Next()
}
