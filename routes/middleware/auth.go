package middleware

import (
	"dofun/app/auth"
	"dofun/app/controllers"
	"dofun/app/helpers"

	"github.com/gin-gonic/gin"
)

// Auth 用户已登录才能访问
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, err := auth.GetCurrentUserFromContext(c)
		if currentUser == nil || err != nil {
			controllers.RedirectRouter(c, "login")
			c.Abort()
			return
		}

		helpers.RecordLastActivedAt(currentUser)
		c.Next()
	}
}
