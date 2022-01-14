package router

import (
	"github.com/gin-gonic/gin"
	"webtest/apps/bbs"
	"webtest/apps/user"
	"webtest/core/router"
)

func Include(engine *gin.Engine) {
	r := router.New(engine)
	//论坛相关	/bbs/...
	r.Assign("/bbs", bbs.UrlPatterns)
	//用户相关	/user/...
	r.Assign("/user", user.UrlPatterns)
}
