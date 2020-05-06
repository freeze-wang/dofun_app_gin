package wrapper

import (
	"dofun/app/http/controllers"
	auth2 "dofun/app/http/controllers/auth"
	userModel "dofun/app/models/user"

	"github.com/gin-gonic/gin"
)

// GetUser 获取用户
func GetUser(handler func(*gin.Context, *userModel.User)) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, err := auth2.GetCurrentUserFromContext(c)
		if currentUser == nil || err != nil {
			controllers.RedirectRouter(c, "login")
			return
		}

		handler(c, currentUser)
	}
}
