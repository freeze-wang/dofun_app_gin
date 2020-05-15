package pw

import (
	"dofun/app/http/controllers"
	"dofun/app/services/dj"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Index topic list
func List(c *gin.Context) {
	var data interface{}

	page,_ := strconv.Atoi(c.DefaultQuery("page","0"))
	pageSize,_ := strconv.Atoi(c.DefaultQuery("pageSize","10"))

	data, err := dj.PwList(c.Query("class_id"), c.Query("attribute_id"), c.Query("sex"), c.Query("orderBy"), page, pageSize)
	if err != nil {
		controllers.SendErrorResponse(c, err)
		return
	}
	controllers.SendOKResponse(c, data)
	return
}
