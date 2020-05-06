package category

import (
	"dofun/app/http/controllers"
	categoryModel "dofun/app/models/category"
	"dofun/app/viewmodels"

	"github.com/gin-gonic/gin"
)

// Index category 列表
func Index(c *gin.Context) {
	cats, _ := categoryModel.All()

	controllers.SendOKResponse(c, controllers.ListData{
		List: viewmodels.CategoryList(cats),
	})
}
