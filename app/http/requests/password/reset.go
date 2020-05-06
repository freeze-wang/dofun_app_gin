package passsword

import (
	"dofun/pkg/ginutils/validate"
)

type PassWordResetForm struct {
	validate.Validate
	Email                string
	Token                string
	Password             string
	PasswordConfirmation string
}




// RegisterMessages 注册错误信息
func (*PassWordResetForm) RegisterMessages() validate.MessagesMap {
	return validate.MessagesMap{
		"password": {
			"密码不能为空",
			"密码长度不能小于 6 个字符",
			"两次输入的密码不一致",
		},
		"token": {
			"token 不能为空",
			"该 token 不存在",
		},
	}
}
