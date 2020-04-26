package policies

import (
	"dofun/app/controllers"
	permissionModel "dofun/app/models/permission"
	userModel "dofun/app/models/user"
	"dofun/pkg/constants"
	"dofun/pkg/errno"

	"github.com/gin-gonic/gin"
)

func before(currentUser *userModel.User) bool {
	if currentUser == nil {
		return false
	}
	hasContentManagePermission := permissionModel.UserHasPermission(currentUser, permissionModel.PermissionNameManageContents)
	return hasContentManagePermission
}

// Unauthorized : 无权限时
func Unauthorized(c *gin.Context) {
	if constants.IsApiRequest(c) {
		controllers.SendErrorResponse(c, errno.AuthError)
		return
	}

	controllers.RenderUnauthorized(c)
}

// CheckPolicy 检查权限
func CheckPolicy(c *gin.Context, currentUser *userModel.User, cond func() bool) bool {
	if before(currentUser) {
		return true
	}

	if cond() {
		return true
	}

	Unauthorized(c)
	return false
}
