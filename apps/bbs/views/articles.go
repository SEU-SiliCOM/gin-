package views

import (
	"github.com/gin-gonic/gin"
	"webtest/core/response"
	"webtest/core/view"
)

type ArticlesView struct {
	view.BaseView
}

// url	/bbs/articles

// GET 获取一页文章
func (*ArticlesView) GET(c *gin.Context) {
	/*
		一段逻辑
	*/
	response.APIResponse(c, 100, "成功获取第一页的文章", gin.H{})
}

// POST 上传一篇文章
func (*ArticlesView) POST(c *gin.Context) {
	/*
		一段逻辑
	*/
	response.APIResponse(c, 100, "成功上传文章", gin.H{})
}
