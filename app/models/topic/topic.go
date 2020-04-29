package topic

import (
	"dofun/app/models"
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

// Topic 话题
type Topic struct {
	ID        uint      `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;not null" binding:"required"`
	TopicName string    `json:"topic_name" gorm:"column:topic_name;not null" binding:"required"`
	TopicType string    `json:"topic_type" gorm:"column:topic_type;not null" binding:"required"`
	IsDefault int       `json:"is_default" gorm:"column:is_default;" binding:"required"`
	Status    int       `json:"status" gorm:"column:status;" binding:"required"`
	CreatedAt models.Time `json:"created_at" gorm:"column:created_at;" binding:"required"`
	UpdatedAt models.Time `json:"updated_at" gorm:"column:updated_at;" binding:"required"`
}

// TableName 表名
func (Topic) TableName() string {
	return "app_topic"
}

// BeforeSave - hook
func (t *Topic) BeforeSave() error {
	return nil
}

// AfterSave - hook
func (t *Topic) AfterSave() error {

	return nil
}

// BeforeDelete - hook
func (t *Topic) BeforeDelete(tx *gorm.DB) (err error) {

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

func setToCache(topic *Topic) {
	key := strconv.Itoa(int(topic.ID))
	topicCache.Set(key, topic, cache.DefaultExpiration)
}

func getFromCache(id int) (*Topic, bool) {
	cachedTopic, ok := topicCache.Get(strconv.Itoa(id))
	if !ok {
		return nil, false
	}

	t, ok := cachedTopic.(*Topic)
	if !ok {
		return nil, false
	}

	return t, true
}

func delCache(id int) {
	topicCache.Delete(strconv.Itoa(id))
}
