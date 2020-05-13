package pw

import (
	"dofun/app/http/controllers"
	"dofun/app/services/dj"
	"github.com/gin-gonic/gin"
)

// Index topic list
func List(c *gin.Context) {
	var data interface{}

	data, err := dj.PwList(c.Param("class_id"), c.Param("attribute_id"),  c.Param("sex"),  c.Param("orderBy"), c.GetInt("page"), c.GetInt("pageSize"))
	if err!=nil{
		controllers.SendErrorResponse(c, err)
		return
	}
	controllers.SendOKResponse(c, data)
	return
}
