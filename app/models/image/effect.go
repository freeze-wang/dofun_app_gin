package image

import (
	"dofun/database"

	"github.com/lexkong/log"
)

// Create -
func (i *Image) Create() (err error) {
	if err = database.DB.Create(&i).Error; err != nil {
		log.Warnf("image 创建失败: %v", err)
		return err
	}

	return nil
}
