package dynamic

import (
	"dofun/app/controllers"
	SystemMenu "dofun/app/models/system"
	"github.com/gin-gonic/gin"
)

// Index topic list
func Index(c *gin.Context) {
	menu, _, status := SystemMenu.GetMenu(c)
	if status==false{
		return
	}
	controllers.SendOKResponse(c,SystemMenu.MenuApi(menu))
	return
}
