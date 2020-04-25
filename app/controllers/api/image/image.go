package image

import (
	"dofun/app/controllers"
	userModel "dofun/app/models/user"
	request "dofun/app/requests/api/image"
	"dofun/app/viewmodels"

	"github.com/gin-gonic/gin"
)

// Store 上传图片
func Store(c *gin.Context, currentUser *userModel.User, tokenString string) {
	img, _ := c.FormFile("image")
	req := &request.Upload{
		Type:  c.PostForm("type"),
		Image: img,
	}

	image, err := req.Run(currentUser)
	if err != nil {
		controllers.SendErrorResponse(c, err)
		return
	}

	controllers.SendOKResponse(c, viewmodels.Image(image))
}
