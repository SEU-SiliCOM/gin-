package router

import "github.com/gin-gonic/gin"

type IView interface {
	GET(c *gin.Context)
	POST(c *gin.Context)
	DELETE(c *gin.Context)

	Middlewares(method string) M
}

type M []gin.HandlerFunc

type UrlPatterns []func() (string, IView, M)

func Path(path string, view IView, middlewares ...gin.HandlerFunc) func() (string, IView, M) {
	return func() (string, IView, M) {
		return path, view, middlewares
	}
}
