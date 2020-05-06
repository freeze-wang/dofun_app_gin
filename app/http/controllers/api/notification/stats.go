package notification

import (
	"dofun/app/http/controllers"
	userModel "dofun/app/models/user"

	"github.com/gin-gonic/gin"
)

// Stats 通知统计
func Stats(c *gin.Context, currentUser *userModel.User, tokenString string) {
	controllers.SendOKResponse(c, map[string]int{
		"unread_count": 0,
	})
}
