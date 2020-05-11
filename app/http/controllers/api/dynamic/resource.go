package dynamic

import (
	"dofun/app/http/controllers"
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
	id, err := ginutils.GetIntParam(c, "id")
	if err != nil {
		controllers.SendErrorResponse(c, errno.Base(errno.ParamsError, "菜单不存在"))
		return
	}
	menu, _ := SystemMenu.GetMenuModel(id)
	if menu == nil {
		controllers.SendErrorResponse(c, errno.Base(errno.ParamsError, "菜单不存在"))
		return
	}
	var data interface{}
	switch menu.MenuType {
	case SystemMenu.MENU_TYPE_DEFAULT:
		data, _ = SystemMenu.GetRecommendDynamic(c, menu)
	case SystemMenu.MENU_TYPE_FOLLOW:
		if data, _ = SystemMenu.GetFollowDynamic(c, menu); data == nil {
			return
		}
	case SystemMenu.MENU_TYPE_MATCH:
		data, _ = SystemMenu.GetMatchDynamic(c, menu)
	case SystemMenu.MENU_TYPE_RECOMMEND:
		data, _ = SystemMenu.GetRecommendDynamic(c, menu)
	default:

	}

	controllers.SendOKResponse(c, data)
	return
}

// Index topic list
func Detail(c *gin.Context) {

	id, _ := ginutils.GetIntParam(c, "id")

	d, err := dynamic.Get(id)
	if err != nil {
		controllers.SendErrorResponse(c, err)
		return
	}
	controllers.SendOKResponse(c, d)

	return
}
