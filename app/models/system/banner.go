package system

import (
	"dofun/app/models"
)

type Banner struct {
	ID int `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;not null" binding:"required"`
	BannerTitle string `json:"banner_title" gorm:"column:banner_title;not null" binding:"required"`
	BannerImage string `json:"banner_image" gorm:"column:banner_image;not null" binding:"required"`
	BannerType string `json:"banner_type" gorm:"column:banner_type;not null" binding:"required"`
	RelationId int `json:"relation_id" gorm:"column:relation_id;not null" binding:"required"`
	RelationUrl string `json:"relation_url" gorm:"column:relation_url;not null" binding:"required"`
	BannerableId int `json:"bannerable_id" gorm:"column:bannerable_id;not null" binding:"required"`
	BannerableType string `json:"bannerable_type" gorm:"column:bannerable_type;not null" binding:"required"`
	Weight int `json:"weight" gorm:"column:weight;not null" binding:"required"`
	Status int `json:"status" gorm:"column:status;not null" binding:"required"`
	ClickFrequency string `json:"click_frequency" gorm:"column:click_frequency;not null" binding:"required"`
	CreatedAt models.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt models.Time `json:"updated_at" gorm:"column:updated_at"`
}
// TableName 表名
func (Banner) TableName() string {
	return "app_banner"
}