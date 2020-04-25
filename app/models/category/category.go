package category

import (
	"dofun/app/models"
)

// Category 分类
type Category struct {
	models.BaseModel
	Name        string `gorm:"column:name;type:varchar(255);not null" sql:"index"`
	Description string `gorm:"column:description;type:text"`
	PostCount   int    `gorm:"column:post_count;type:int;default:0"` // 分类下的帖子数量
}

// TableName 表名
func (Category) TableName() string {
	return "categories"
}
