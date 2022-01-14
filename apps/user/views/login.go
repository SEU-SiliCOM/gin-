package views

import (
	"github.com/gin-gonic/gin"
	"webtest/core/response"
	"webtest/core/view"
)

type LoginView struct {
	view.BaseView
}

// url	/user/login

// POST 登录
func (*LoginView) POST(c *gin.Context) {
	/*
		一段逻辑
	*/
	response.APIResponse(c, 100, "成功登录", gin.H{})
}
