package middleware

import (
	"dofun/app/auth"
	"dofun/app/helpers"
	"dofun/app/http/controllers"

	"strings"

	"github.com/gin-gonic/gin"
)

var whitePathList = [...]string{
	"email",
	"logout",
}

// CurrentUserMiddleware : 从 session 中获取 user model 的 middleware
func CurrentUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.SaveCurrentUserToContext(c)
		if user != nil {
			helpers.RecordLastActivedAt(user)
		}

		// 如果用户已经登录，并且未有耨恒 email
		// 并且访问的不是 email 验证或退出的 url，会重定向到激活页
		if user != nil && !user.IsActivated() {
			if !inWhitePathList(c.Request.URL.Path) {
				controllers.RedirectRouter(c, "verification.notice")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

// path 是否存在于白名单中
func inWhitePathList(path string) bool {
	for _, v := range whitePathList {
		if strings.Contains(path, v) {
			return true
		}
	}

	return false
}
