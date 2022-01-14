package main

import (
	"github.com/gin-gonic/gin"
	"webtest/middlewares"
	"webtest/router"
)

func main() {
	engine := gin.New()
	//无效路径时触发
	engine.NoRoute(middlewares.RouteError)
	//服务器内部出错时触发
	engine.Use(middlewares.InternalError)
	//路由分发
	router.Include(engine)
	//服务器运行
	err := engine.Run(":8000")
	if err != nil {
		return
	}
}
