package password

import (
	userModel "dofun/app/models/user"
	"dofun/pkg/ginutils/validate"
)

func emailExistValidator(email string) validate.ValidatorFunc {
	return func() string {
		if _, err := userModel.GetByEmail(email); err == nil {
			return ""
		}
		return "该邮箱不存在"
	}
}
