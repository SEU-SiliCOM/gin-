package router

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func (router *Router) Assign(rootPath string, urlPatterns UrlPatterns, middlewares ...gin.HandlerFunc) {
	r := router.engine.Group(rootPath, middlewares...)
	for _, urlPattern := range urlPatterns {
		path, view, middlewares := urlPattern()
		_r := r.Group(path)

		if v := view.Middlewares("GET"); v != nil {
			_r.Handle("GET", "", append(v, view.GET)...)
		} else {
			_r.Handle("GET", "", append(middlewares, view.GET)...)
		}

		if v := view.Middlewares("POST"); v != nil {
			_r.Handle("POST", "", append(v, view.POST)...)
		} else {
			_r.Handle("POST", "", append(middlewares, view.POST)...)
		}

		if v := view.Middlewares("DELETE"); v != nil {
			_r.Handle("DELETE", "", append(v, view.DELETE)...)
		} else {
			_r.Handle("DELETE", "", append(middlewares, view.DELETE)...)
		}
	}
}

func New(engine *gin.Engine) Router {
	return Router{engine}
}
