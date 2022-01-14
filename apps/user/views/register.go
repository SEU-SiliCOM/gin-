package views

import (
	"github.com/gin-gonic/gin"
	"webtest/core/response"
	"webtest/core/view"
)

type RegisterView struct {
	view.BaseView
}

// url	/user/register

// POST 注册账号
func (*RegisterView) POST(c *gin.Context) {
	/*
		一段逻辑
	*/
	response.APIResponse(c, 100, "成功注册", gin.H{})
}

// DELETE 注销账号
func (*RegisterView) DELETE(c *gin.Context) {
	/*
		一段逻辑
	*/
	response.APIResponse(c, 100, "成功注销账户", gin.H{})
}
