package user

import (
	. "dofun/app/controllers"
	"dofun/app/helpers"
	"dofun/app/viewmodels"

	"github.com/gin-gonic/gin"
)

// ActivedIndex 活跃用户列表
func ActivedIndex(c *gin.Context) {
	activeUsersVM := make([]map[string]interface{}, 0)
	activeUsers := helpers.NewActiveUser().GetActiveUsers()
	for _, v := range activeUsers {
		activeUsersVM = append(activeUsersVM, viewmodels.NewUserAPISerializer(v))
	}

	SendOKResponse(c, ListData{
		List: activeUsersVM,
	})
}
