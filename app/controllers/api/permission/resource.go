package permission

import (
	. "dofun/app/controllers"
	permissionModel "dofun/app/models/permission"
	userModel "dofun/app/models/user"
	"dofun/app/viewmodels"

	"github.com/gin-gonic/gin"
)

// Index 用户权限列表
func Index(c *gin.Context, currentUser *userModel.User, tokenString string) {
	all, _ := permissionModel.GetUserAllPermission(currentUser)
	SendOKResponse(c, ListData{
		List: viewmodels.PermissionAPIList(all),
	})
}
