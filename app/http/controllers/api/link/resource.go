package link

import (
	"dofun/app/http/controllers"
	linkModel "dofun/app/models/link"
	"dofun/app/viewmodels"

	"github.com/gin-gonic/gin"
)

// Index 资源链接列表
func Index(c *gin.Context) {
	all, _ := linkModel.All()
	controllers.SendOKResponse(c, controllers.ListData{
		List: viewmodels.LinkAPIList(all),
	})
}
