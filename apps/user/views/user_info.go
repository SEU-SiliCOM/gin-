package views

import (
	"github.com/gin-gonic/gin"
	"webtest/core/response"
	"webtest/core/view"
)

type UserInfoView struct {
	view.BaseView
}

// url	/user/user-info/:username

// GET 获取用户信息
func (*UserInfoView) GET(c *gin.Context) {
	/*
		一段逻辑
	*/
	response.APIResponse(c, 100, "已获取用户"+c.Param("username")+"的信息", gin.H{})
}

// POST 修改用户信息
func (*UserInfoView) POST(c *gin.Context) {
	/*
		一段逻辑
	*/
	response.APIResponse(c, 100, "已修改用户"+c.Param("username")+"的信息", gin.H{})
}
