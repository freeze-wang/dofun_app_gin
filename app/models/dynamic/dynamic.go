package dynamic

import (
	"dofun/app/helpers"
	"dofun/app/models"
	"dofun/database"
	"dofun/pkg/ginutils/utils"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"
)

var (
	topicCache = cache.New(10*time.Minute, 30*time.Minute)
)

// Menu 菜单
type Menu struct {
	models.BaseModel
	Title           string `gorm:"column:title;type:varchar(255);not null" sql:"index"` // 帖子标题
	Body            string `gorm:"column:body;type:text;not null"`                      // 帖子内容
	UserID          uint   `gorm:"column:user_id;not null" sql:"index"`                 // 用户 ID
	CategoryID      uint   `gorm:"column:category_id;not null" sql:"index"`             // 分类 ID
	ReplyCount      int    `gorm:"column:reply_count;not null;default:0"`               // 回复数量
	ViewCount       int    `gorm:"column:view_count;not null;default:0"`                // 查看总数
	LastReplyUserID uint   `gorm:"column:last_reply_user_id;not null;default:0"`        // 最后回复的用户 ID
	Order           int    `gorm:"column:order;not null;default:0"`                     // 排序
	Excerpt         string `gorm:"column:excerpt;type:text"`                            // 文章摘要，SEO 优化时使用
	Slug            string `gorm:"column:slug;type:varchar(255)"`                       // SEO 友好的 URI
}

// TableName 表名
func (Menu) TableName() string {
	return "app_system_menu"
}


// BeforeSave - hook
func (t *Menu) BeforeSave() error {
	t.Body = utils.XSSClean(t.Body)
	t.Excerpt = makeExcerpt(t.Body, 200)

	return nil
}

// AfterSave - hook
func (t *Menu) AfterSave() error {
	// if t.Slug == "" {
	// SlugTranslate 需要请求百度 api，放到协程中执行
	go func(t *Menu) {
		slug := helpers.SlugTranslate(t.Title)
		database.DB.Model(&t).UpdateColumn("slug", slug) // 这样更新可避免触发 gorm 钩子，从而导致死循环
	}(t)
	// }

	return nil
}

// BeforeDelete - hook
func (t *Menu) BeforeDelete(tx *gorm.DB) (err error) {
	// 话题删除时，删除其所属的回复
	tx.Exec("delete from replies where topic_id = ?", t.ID)

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

func setToCache(topic *Menu) {
	key := strconv.Itoa(int(topic.ID))
	topicCache.Set(key, topic, cache.DefaultExpiration)
}

func getFromCache(id int) (*Menu, bool) {
	cachedTopic, ok := topicCache.Get(strconv.Itoa(id))
	if !ok {
		return nil, false
	}

	t, ok := cachedTopic.(*Menu)
	if !ok {
		return nil, false
	}

	return t, true
}

func delCache(id int) {
	topicCache.Delete(strconv.Itoa(id))
}
