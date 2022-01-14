package view

import (
	"github.com/gin-gonic/gin"
	"webtest/core/response"
	"webtest/core/router"
)

type BaseView struct{}

func (*BaseView) GET(c *gin.Context) {
	response.APIResponse(c, 101, "不支持GET方式", gin.H{})
}

func (*BaseView) POST(c *gin.Context) {
	response.APIResponse(c, 101, "不支持POST方式", gin.H{})
}

func (*BaseView) DELETE(c *gin.Context) {
	response.APIResponse(c, 101, "不支持DELETE方式", gin.H{})
}

func (*BaseView) Middlewares(method string) router.M {
	return nil
}
