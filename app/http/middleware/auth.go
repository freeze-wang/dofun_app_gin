package middleware

import (
	"dofun/app/handlers"
	"dofun/app/http/controllers"
	auth2 "dofun/app/http/controllers/auth"

	"github.com/gin-gonic/gin"
)

// Auth 用户已登录才能访问
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, err := auth2.GetCurrentUserFromContext(c)
		if currentUser == nil || err != nil {
			controllers.RedirectRouter(c, "login")
			c.Abort()
			return
		}

		handlers.RecordLastActivedAt(currentUser)
		c.Next()
	}
}
