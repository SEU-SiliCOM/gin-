package middlewares

import (
	"github.com/gin-gonic/gin"
	"webtest/core/response"
)

func RouteError(c *gin.Context) {
	response.APIResponse(c, 0, "404", gin.H{})
	c.Abort()
}
