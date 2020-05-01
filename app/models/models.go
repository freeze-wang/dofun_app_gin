package models

import (
	"database/sql/driver"
	"dofun/pkg/ginutils/pagination"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

const (
	// TrueTinyint true
	TrueTinyint = 1
	// FalseTinyint false
	FalseTinyint = 0
)

// ListData 带列表时的 data
type ListData struct {
	Page      int         `json:"page,omitempty"`      // 当前页数
	PageLine  int         `json:"per_page,omitempty"`  // 每页条数
	Total     int         `json:"total"`               // 总数
	Data      interface{} `json:"data"`                // 列表数据 (无数据时，默认返回一个 [])
	TotalPage interface{} `json:"totalPage,omitempty"` // 其他数据 (可选)
}

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

// Paginate 分页数据
func Paginate(c *gin.Context,
	returnDataFunc func(int, int) interface{},
	totalCount int, defaultPageLine int) ListData {

	// 从 request query 中获取分页参数
	offset, limit, currentPage, totalPage := pagination.GetPageQuery(c, defaultPageLine, totalCount)

	// 得到列表数据
	items := returnDataFunc(offset, limit)

	listData := ListData{
		Page:      currentPage,
		PageLine:  pagination.GetPageLine(c, defaultPageLine),
		Total:     totalCount,
		TotalPage: totalPage,
		Data:      items,
	}
	return listData
}
