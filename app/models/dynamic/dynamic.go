package dynamic

import (
	"dofun/app/models/topic"
	"dofun/app/models/user"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"
)

var (
	dynamicCache = cache.New(10*time.Minute, 30*time.Minute)
)

// Dynamic 动态
type Dynamic struct {
	ID             int         `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;not null" binding:"required"`
	TopicId        int         `json:"topic_id" gorm:"column:topic_id;not null" binding:"required"`
	Content        string      `json:"content" gorm:"column:content;not null" binding:"required"`
	UserType       string      `json:"user_type" gorm:"column:user_type;not null" binding:"required"`
	UserId         int         `json:"user_id" gorm:"column:user_id;not null" binding:"required"`
	AdminId        int         `json:"admin_id" gorm:"column:admin_id;" binding:"required"`
	PraisePoints   int         `json:"praise_points" gorm:"column:praise_points;" binding:"required"`
	ShareNumber    int         `json:"share_number" gorm:"column:share_number;" binding:"required"`
	IsTop          int         `json:"is_top" gorm:"column:is_top;" binding:"required"`
	Status         int         `json:"status" gorm:"column:status;" binding:"required"`
	Remark         string      `json:"remark" gorm:"column:remark;not null" binding:"required"`
	JumpIconUrl    string      `json:"jump_icon_url" gorm:"column:jump_icon_url;not null" binding:"required"`
	JumpCopy       string      `json:"jump_copy" gorm:"column:jump_copy;not null" binding:"required"`
	JumpType       string      `json:"jump_type" gorm:"column:jump_type;not null" binding:"required"`
	JumpLocation   string      `json:"jump_location" gorm:"column:jump_location;not null" binding:"required"`
	ClickNumber    int         `json:"click_number" gorm:"column:click_number;not null" binding:"required"`
	ClickFrequency string      `json:"click_frequency" gorm:"column:click_frequency;not null" binding:"required"`
	DeleteStatus   int         `json:"delete_status" gorm:"column:delete_status;not null" binding:"required"`
	HotAt          time.Time   `json:"hot_at" gorm:"column:hot_at;not null" binding:"required"`
	CreatedAt      time.Time   `json:"created_at" gorm:"column:created_at;not null" binding:"required"`
	UpdatedAt      time.Time   `json:"updated_at" gorm:"column:updated_at;not null" binding:"required"`
	User           user.User   `json:"user" gorm:"ForeignKey:user_id"`
	Topic          topic.Topic `json:"topic" gorm:"ForeignKey:topic_id"`
}

// TableName 表名
func (Dynamic) TableName() string {
	return "app_topic_dynamic"
}

// BeforeSave - hook
func (t *Dynamic) BeforeSave() error {

	return nil
}

// AfterSave - hook
func (t *Dynamic) AfterSave() error {

	return nil
}

// BeforeDelete - hook
func (t *Dynamic) BeforeDelete(tx *gorm.DB) (err error) {

	return
}

// ------------ private
func makeExcerpt(value string, length int) string {
	r := regexp.MustCompile(`\r\n|\r|\n+|\<[\S\s]+?\>`)
	v := string(r.ReplaceAll([]byte(value), []byte("")))
	v = strings.TrimSpace(v)
	ru := []rune(v) // utf8 字符串需先转 rune 才可 [:]

	if len(ru) < length {
		return v
	}
	return string(ru[:length])
}

func setToCache(dynamic *Dynamic) {
	key := strconv.Itoa(int(dynamic.ID))
	dynamicCache.Set(key, dynamic, cache.DefaultExpiration)
}

func getFromCache(id int) (*Dynamic, bool) {
	cachedDynamic, ok := dynamicCache.Get(strconv.Itoa(id))
	if !ok {
		return nil, false
	}

	t, ok := cachedDynamic.(*Dynamic)
	if !ok {
		return nil, false
	}

	return t, true
}

func delCache(id int) {
	dynamicCache.Delete(strconv.Itoa(id))
}
