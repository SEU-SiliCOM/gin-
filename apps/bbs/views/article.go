package views

import (
	"github.com/gin-gonic/gin"
	"webtest/core/response"
	"webtest/core/router"
	"webtest/core/view"
	"webtest/middlewares"
)

type ArticleView struct {
	view.BaseView
}

// url	/bbs/article/:id

// GET 获取单篇文章
func (*ArticleView) GET(c *gin.Context) {
	/*
		一段逻辑
	*/
	response.APIResponse(c, 100, "成功获取id为"+c.Param("id")+"的文章", gin.H{})
}

// POST 修改单篇文章
func (*ArticleView) POST(c *gin.Context) {
	/*
		一段逻辑
	*/
	response.APIResponse(c, 100, "已修改id为"+c.Param("id")+"的文章", gin.H{})
}

// DELETE 删除单篇文章
func (*ArticleView) DELETE(c *gin.Context) {
	/*
		一段逻辑
	*/
	response.APIResponse(c, 100, "已删除id为"+c.Param("id")+"的文章", gin.H{})
}

func (*ArticleView) Middlewares(method string) router.M {
	if method == "POST" {
		return router.M{middlewares.RouteError}
	}
	return nil
}
