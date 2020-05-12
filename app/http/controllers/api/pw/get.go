package pw

import (
	"dofun/app/http/controllers"
	"dofun/app/services/dj"
	"github.com/gin-gonic/gin"
)

// Index topic list
func List(c *gin.Context) {
	var data interface{}

	data = dj.PwList("","","","",1,10)
	controllers.SendOKResponse(c, data)
	return
}
