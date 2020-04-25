package requests

import (
	userModel "dofun/app/models/user"
	"dofun/pkg/ginutils/captcha"
	"dofun/pkg/ginutils/validate"
)

// NameUniqueValidator name 唯一
func NameUniqueValidator(name string, id int) validate.ValidatorFunc {
	return func() (msg string) {
		u, err := userModel.GetByName(name)
		if err != nil || u.ID == uint(id) {
			return ""
		}

		return "用户名已经被注册过了"
	}
}

// EmailUniqueValidator 邮箱唯一
func EmailUniqueValidator(email string) validate.ValidatorFunc {
	return func() (msg string) {
		if _, err := userModel.GetByEmail(email); err != nil {
			return ""
		}
		return "邮箱已经被注册过了"
	}
}

// PhoneUniqueValidator 手机号唯一
func PhoneUniqueValidator(phone string) validate.ValidatorFunc {
	return func() (msg string) {
		if _, err := userModel.GetByPhone(phone); err != nil {
			return ""
		}
		return "手机已经被注册过了"
	}
}

// CaptchaValidator 验证码验证
func CaptchaValidator(captchaID, captchaVal string) validate.ValidatorFunc {
	return func() (msg string) {
		if ok := captcha.Verify(captchaID, captchaVal); ok {
			return ""
		}
		return "验证码错误"
	}
}
