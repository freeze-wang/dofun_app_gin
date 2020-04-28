package dynamic

import (
	"dofun/app/controllers"
	"dofun/app/models/dynamic"
	SystemMenu "dofun/app/models/system"
	"dofun/pkg/errno"
	"dofun/pkg/ginutils"
	"github.com/gin-gonic/gin"
)

// Index topic list
func Index(c *gin.Context) {
	/*menu, _, status := SystemMenu.GetMenu(c)
	if !status {
		controllers.SendErrorResponse(c, errno.Base(errno.ParamsError, "不存在"))
		return
	}
	controllers.SendOKResponse(c, SystemMenu.MenuApi(menu))*/

	/*id, _ := ginutils.GetIntParam(c, "id")

	dynamic, err := dynamic.Get(id)
	if err != nil{
		controllers.SendErrorResponse(c,err)
		return
	}
	controllers.SendOKResponse(c, dynamic)*/

	dynamic, status := SystemMenu.GetDynamic(c)
	if !status {
		controllers.SendErrorResponse(c, errno.Base(errno.ParamsError, "不存在"))
		return
	}
	controllers.SendOKResponse(c, dynamic)
	return
}
// Index topic list
func Detail(c *gin.Context) {

	id, _ := ginutils.GetIntParam(c, "id")

	d, err := dynamic.Get(id)
	if err != nil{
		controllers.SendErrorResponse(c,err)
		return
	}
	controllers.SendOKResponse(c, d)

	return
}
