package topic

import (
	catModel "dofun/app/models/category"
	"dofun/pkg/ginutils/validate"
)

func categoryIDValidator(catid uint) validate.ValidatorFunc {
	return func() string {
		if catid == 0 {
			return "分类 id 不可为空"
		}
		_, err := catModel.Get(int(catid))
		if err != nil {
			return "该分类不存在"
		}

		return ""
	}
}
