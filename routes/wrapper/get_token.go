package wrapper

import (
	"dofun/app/auth/token"
	"dofun/app/http/controllers"
	userModel "dofun/app/models/user"
	"dofun/pkg/errno"

	"github.com/gin-gonic/gin"
)

// GetToken 获取 token
func GetToken(handler func(*gin.Context, *userModel.User, string)) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, user, ok := token.GetTokenUserFromContext(c)
		if !ok {
			controllers.SendErrorResponse(c, errno.TokenError)
			return
		}

		handler(c, user, tokenStr)
	}
}
