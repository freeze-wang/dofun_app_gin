package topic

import (
	"dofun/database"
)

const (
	orderDefault = "default"
	orderRecent  = "recent"
)

// Get -
func Get(id int) (*Topic, error) {
	if cachedTopic, ok := getFromCache(id); ok {
		return cachedTopic, nil
	}

	t := &Topic{}
	if err := database.DB.First(&t, id).Error; err != nil {
		return t, err
	}
	setToCache(t)

	return t, nil
}

// List -
func List(offset, limit int, order string) (topics []*Topic, err error) {
	topics = make([]*Topic, 0)

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
	err = database.DB.Model(&Topic{}).Count(&count).Error
	return
}

// CountByCategoryID -
func CountByCategoryID(categoryID int) (count int, err error) {
	err = database.DB.Model(&Topic{}).Where("category_id = ?", categoryID).Count(&count).Error
	return
}

// CountByUserID -
func CountByUserID(userID int) (count int, err error) {
	err = database.DB.Model(&Topic{}).Where("user_id = ?", userID).Count(&count).Error
	return
}

// All -
func All() (topics []*Topic, err error) {
	topics = make([]*Topic, 0)

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
func GetByCategoryID(categoryID, offset, limit int, order string) (topics []*Topic, err error) {
	topics = make([]*Topic, 0)

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
func GetByUserID(userID, offset, limit int, order string) (topics []*Topic, err error) {
	topics = make([]*Topic, 0)

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

