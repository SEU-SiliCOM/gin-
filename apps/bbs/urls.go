package bbs

import (
	"webtest/apps/bbs/views"
	"webtest/core/router"
)

var UrlPatterns = router.UrlPatterns{
	router.Path("/articles", &views.ArticlesView{}),
	router.Path("/article/:id/", &views.ArticleView{}),
}
