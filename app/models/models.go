package models

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

const (
	// TrueTinyint true
	TrueTinyint = 1
	// FalseTinyint false
	FalseTinyint = 0
)
type Time struct { // 内嵌方式（推荐）
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	// tune := fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))
	tune := t.Format(`"2006-01-02 15:04:05"`)
	return []byte(tune), nil
}

// Value insert timestamp into mysql need this function.
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
// BaseModel model 基类
type BaseModel struct {
	ID uint `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	// MySQL的DATE/DATATIME类型可以对应Golang的time.Time
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	// 有 DeletedAt(类型需要是 *time.Time) 即支持 gorm 软删除
	//DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index"`
}

// GetIDstring 获取字符串形式的 id
func (m *BaseModel) GetIDstring() string {
	return strconv.Itoa(int(m.ID))
}
