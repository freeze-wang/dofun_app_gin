package middleware

import (
	"dofun/app/handlers"
	"dofun/app/http/controllers"
	"dofun/app/http/controllers/auth/token"

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

		handlers.RecordLastActivedAt(currentUser)
		c.Next()
	}
}
//  TokenRefresh token 自动刷新,非强制登录
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