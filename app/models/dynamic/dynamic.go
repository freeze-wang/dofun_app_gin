package dynamic

import (
	"dofun/app/models"
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

const (
	STATUS_DISABLE = 0
	STATUS_ABLE = 1
	DELEET_STATUS_DEFAULT = 0
	DELETE_STATUS_USER_OPERATE = 1
	DELETE_STATUS_ADMIN_OPERATE = 2
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
	HotAt          models.Time `json:"hot_at" gorm:"column:hot_at;not null" binding:"required"`
	CreatedAt      models.Time `json:"created_at" gorm:"column:created_at;not null" binding:"required"`
	UpdatedAt      models.Time `json:"updated_at" gorm:"column:updated_at;not null" binding:"required"`
	User           *Brief   `json:"user" gorm:"ForeignKey:user_id"`
	Topic          *topic.Topic `json:"topic" gorm:"ForeignKey:topic_id"`
	AppTopicDynamicDetail []AppTopicDynamicDetail	`json:"detail" gorm:"foreignkey:dynamic_id"`
}
// 动态媒体内容表
type AppTopicDynamicDetail struct {
	Id              int32  `json:"id" gorm:"id"`
	DynamicId       int32  `json:"dynamic_id" gorm:"dynamic_id"`               // 动态id
	MediaType       string `json:"media_type" gorm:"media_type"`               // 内容媒体类型
	MediaUrl        string `json:"media_url" gorm:"media_url"`                 // 内容链接
	MediaTimeLength string `json:"media_time_length" gorm:"media_time_length"` // 媒体播放时长
	CoverImgUrl     string `json:"cover_img_url" gorm:"cover_img_url"`         // 封面图片地址

}
// User 简略版用户模型
type Brief struct {
	ID            uint    `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;not null" binding:"required"`
	ZhwId         int     `json:"zhw_id" gorm:"column:zhw_id;not null" binding:"required"`
	ZhwUsername   string  `json:"zhw_username" gorm:"column:zhw_username;not null" binding:"required"`
	Phone         string  `json:"phone" gorm:"column:phone;not null" binding:"required"`
	Avatar        string  `json:"avatar" gorm:"column:avatar;not null" binding:"required"`
	Nickname      string  `json:"nickname" gorm:"column:nickname;not null" binding:"required"`
	FollowDynamic []Dynamic `gorm:"many2many:app_dynamic_follow;association_jointable_foreignkey:dynamic_id;jointable_foreignkey:user_id;" json:"follow_dynamic"`
	//Dynamic       Dynamic `json:"user" gorm:"ForeignKey:user_id" `
}

// TableName 表名
func (Brief) TableName() string {
	return user.TableName
}

// TableName 表名
func (Dynamic) TableName() string {
	return "app_topic_dynamic"
}
// TableName 表名
func (AppTopicDynamicDetail) TableName() string {
	return "app_topic_dynamic_detail"
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
