package system

import (
	"dofun/app/auth/token"
	"dofun/app/controllers"
	"dofun/app/models/dynamic"
	"dofun/app/models/topic"
	"dofun/database"
	"dofun/pkg/errno"
	"dofun/pkg/ginutils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const (
	orderDefault = "default"
	orderRecent  = "recent"
)

// 获取要编辑的 topic
func GetMenu(c *gin.Context) (*Menu, int, bool) {
	id, err := ginutils.GetIntParam(c, "id")
	if err != nil {
		return nil, id, false
	}

	menu, err := GetMenuModel(id)
	if err != nil {
		return nil, id, false
	}

	return menu, id, true
}

func GetTopic(c *gin.Context) (interface{}, bool) {
	topics := make([]*topic.Topic, 0)
	id, err := ginutils.GetIntParam(c, "id")
	if err != nil {
		return nil, false
	}
	menu, err := GetMenuModel(id)
	database.DB.Model(&menu).Related(&topics, "topic")
	return topics, true
}
func GetRecommendDynamic(c *gin.Context, menu *Menu) (interface{}, bool) {
	data := make(map[string]interface{})
	dynamics := make([]*dynamic.Dynamic, 0)
	banners := make([]*Banner, 0)
	//user := token.User(c)

	database.DB.Model(menu).Preload("Topic", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,topic_name,topic_type")
	}).Preload("User").Related(&dynamics, "dynamic")

	database.DB.Model(menu).Where("bannerable_type = ? and status = ? ", "App\\Models\\V1\\System\\Menu", 1).Limit(5).Order("weight DESC").Related(&banners, "banner")
	data["type"] = menu.MenuType
	data["banner"] = banners
	data["list"] = dynamics
	return data, true
}

func GetFollowDynamic(c *gin.Context) (interface{}, bool) {
	data := make([]*dynamic.Dynamic, 0)
	brief := &dynamic.Brief{}
	user := token.User(c)
	if user == nil {
		controllers.SendErrorResponse(c, errno.SessionError)
		return nil, false
	}
	if err := database.DB.First(&brief, user.ID).Error; err != nil {
		return nil, false
	}

	database.DB.Model(brief).Preload("User").Preload("Topic").Related(&data, "follow_dynamic")

	return data, true
}

func GetMatchDynamic(c *gin.Context, menu *Menu) (interface{}, bool) {
	data := make(map[string]interface{})
	dynamics := make([]*dynamic.Dynamic, 0)
	banners := make([]*Banner, 0)
	//user := token.User(c)

	if database.DB.Model(menu).Preload("Topic", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,topic_name,topic_type")
	}).Preload("User").Related(&dynamics, "dynamic").RecordNotFound() {
		return nil, false
	}

	database.DB.Model(menu).Where("bannerable_type = ? and status = ? ", "App\\Models\\V1\\System\\Menu", 1).Limit(5).Order("weight DESC").Related(&banners, "banner")
	data["type"] = menu.MenuType
	data["banner"] = banners
	data["list"] = dynamics
	return data, true
}

// Get -
func GetMenuModel(id int) (*Menu, error) {

	t := &Menu{}
	if err := database.DB.First(&t, id).Error; err != nil {
		return t, err
	}

	return t, nil
}

// List -
func List(offset, limit int, order string) (topics []*Menu, err error) {
	topics = make([]*Menu, 0)

	if order == orderRecent {
		order = "created_at"
	} else {
		order = "updated_at"
	}

	if err = database.DB.Offset(offset).Limit(limit).Order(order + " desc").Find(&topics).Error; err != nil {
		return topics, err
	}

	return topics, nil
}

// Count -
func Count() (count int, err error) {
	err = database.DB.Model(&Menu{}).Count(&count).Error
	return
}

// CountByCategoryID -
func CountByCategoryID(categoryID int) (count int, err error) {
	err = database.DB.Model(&Menu{}).Where("category_id = ?", categoryID).Count(&count).Error
	return
}

// CountByUserID -
func CountByUserID(userID int) (count int, err error) {
	err = database.DB.Model(&Menu{}).Where("user_id = ?", userID).Count(&count).Error
	return
}

// All -
func All() (topics []*Menu, err error) {
	topics = make([]*Menu, 0)

	if err = database.DB.Order("created_at desc").Find(&topics).Error; err != nil {
		return topics, err
	}

	return topics, nil
}

// AllID -
func AllID() (ids []uint, err error) {
	ids = make([]uint, 0)
	topics, err := All()
	if err != nil {
		return ids, err
	}

	for _, t := range topics {
		ids = append(ids, t.ID)
	}

	return ids, nil
}

// GetByCategoryID 根据 category_id 获取 topics
func GetByCategoryID(categoryID, offset, limit int, order string) (topics []*Menu, err error) {
	topics = make([]*Menu, 0)

	if order == orderRecent {
		order = "created_at"
	} else {
		order = "updated_at"
	}

	if err = database.DB.Where("category_id = ?", categoryID).Offset(offset).Limit(limit).Order(order + " desc").Find(&topics).Error; err != nil {
		return topics, err
	}

	return topics, nil
}

// GetByUserID -
func GetByUserID(userID, offset, limit int, order string) (topics []*Menu, err error) {
	topics = make([]*Menu, 0)

	if order == orderRecent {
		order = "created_at"
	} else {
		order = "updated_at"
	}

	if err = database.DB.Where("user_id = ?", userID).Offset(offset).Limit(limit).Order(order + " desc").Find(&topics).Error; err != nil {
		return topics, err
	}

	return topics, nil
}
