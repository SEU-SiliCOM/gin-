package user

import (
	"webtest/apps/user/views"
	"webtest/core/router"
)

var UrlPatterns = router.UrlPatterns{
	router.Path("/user-info/:username", &views.UserInfoView{}),
	router.Path("/login", &views.LoginView{}),
	router.Path("/register", &views.RegisterView{}),
}
