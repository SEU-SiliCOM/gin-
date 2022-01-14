package response

import (
	"github.com/gin-gonic/gin"
	"webtest/settings/response_code"
)

func APIResponse(c *gin.Context, code int, msg string, result gin.H) {
	c.JSON(200, gin.H{"code": code, "msg": msg, "result": result})
}

func InvalidParamsResponse(c *gin.Context) {
	c.JSON(200, gin.H{"code": response_code.InvalidParams, "msg": "无效的参数"})
}
