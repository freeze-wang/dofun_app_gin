package requests

import (
	"dofun/app/http/controllers"
	"dofun/pkg/errno"
	"dofun/pkg/ginutils/validate"

	"github.com/gin-gonic/gin"
)

// RunPhoneValidate 验证手机
func RunPhoneValidate(c *gin.Context) (string, bool) {
	var p struct {
		Phone string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&p); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return "", false
	}

	ok, _, errMap := validate.RunSingle("phone",
		[]validate.ValidatorFunc{
			validate.RequiredValidator(p.Phone),
			validate.PhoneValidator(p.Phone),
			PhoneUniqueValidator(p.Phone),
		},
		[]string{"手机号不能为空", "手机号格式错误"})

	if !ok {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, errMap))
		return "", false
	}

	return p.Phone, true
}
