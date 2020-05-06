package middleware

import (
	"dofun/app/auth/token"
	"dofun/app/helpers"
	"dofun/app/http/controllers"

	"github.com/gin-gonic/gin"
)

// TokenAuth token 验证
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := token.GetTokenFromRequest(c)
		if err != nil || tokenStr == "" {
			controllers.SendErrorResponse(c, err)
			c.Abort()
			return
		}

		currentUser, err := token.ParseAndGetUser(c, tokenStr) // 会将用户数据和 token 存到 gin.Context 中
		if err != nil {
			controllers.SendErrorResponse(c, err)
			c.Abort()
			return
		}

		helpers.RecordLastActivedAt(currentUser)
		c.Next()
	}
}
// TokenAuth token 验证
func TokenRefresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := token.GetTokenFromRequest(c)
		if err != nil || tokenStr == "" {
			c.Next()
			return
		}
		_, _ = token.ParseAndGetUser(c, tokenStr) // 会将用户数据和 token 存到 gin.Context 中

		c.Next()
	}
}