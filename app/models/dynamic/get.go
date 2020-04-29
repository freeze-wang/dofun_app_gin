package dynamic

import (
	"dofun/database"
)

// Get -
func Get(id int) (*Dynamic, error) {
	r := &Dynamic{}
	if err := database.DB.Preload("User").First(&r, id).Error; err != nil {
		return r, err
	}

	return r, nil
}
// List -
func List(notifiableType string, notifiableID uint, offset, limit int) ([]interface{}, error) {
	ns := make([]*Dynamic, 0)
	result := make([]interface{}, 0)

	if err := database.DB.Where("notifiable_type = ? AND notifiable_id = ?",
		notifiableType,
		notifiableID,
	).Offset(offset).Limit(limit).Order("created_at").Find(&ns).Error; err != nil {
		return result, err
	}

	return result, nil
}