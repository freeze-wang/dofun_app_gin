package notification

import (
	"dofun/app/controllers"
	notificationModel "dofun/app/models/notification"
	userModel "dofun/app/models/user"
	"dofun/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Read 标记消息通知为已读
func Read(c *gin.Context, currentUser *userModel.User, tokenString string) {
	if err := currentUser.Notification(0); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.DatabaseError, err))
		return
	}
	if err := notificationModel.Read(userModel.TableName, currentUser.ID); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.DatabaseError, err))
		return
	}

	controllers.SendOKResponse(c, nil)
}
