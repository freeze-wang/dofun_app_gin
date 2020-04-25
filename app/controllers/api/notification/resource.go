package notification

import (
	"dofun/app/controllers"
	notificationModel "dofun/app/models/notification"
	userModel "dofun/app/models/user"

	"github.com/gin-gonic/gin"
)

// Index 通知列表
func Index(c *gin.Context, currentUser *userModel.User, tokenString string) {
	controllers.SendListResponse(c, 20, nil,
		notificationModel.AllCount,
		func(offset, limit, _, _ int) (interface{}, error) {
			return notificationModel.List(userModel.TableName, currentUser.ID, offset, limit)
		})
}
