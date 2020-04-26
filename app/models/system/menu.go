package system

import (
	"dofun/app/controllers"
	"dofun/pkg/constants"
	"dofun/pkg/errno"
	"dofun/pkg/ginutils"
	"dofun/pkg/ginutils/utils"
	"github.com/gin-gonic/gin"
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
	ID uint `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	MenuName  string    `gorm:"column:menu_name;type:varchar(30);not null"`  // 菜单名
	MenuType  string    `gorm:"column:menu_type;type:varchar(100);not null"` // 菜单类型
	IsDefault int       `gorm:"column:is_default;not null"`                  // 是否默认
	Status    uint      `gorm:"column:status;not null"`                      // 状态
	Weight    int       `gorm:"column:weight;not null;default:0"`            // 回复数量
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName 表名
func (Menu) TableName() string {
	return "app_system_menu"
}

// 获取要编辑的 topic
func GetMenu(c *gin.Context) (*Menu, int, bool) {
	id, err := ginutils.GetIntParam(c, "id")
	if err != nil {
		controllers.SendErrorResponse(c, errno.Base(errno.ParamsError, "id 不存在"))
		return nil, id, false
	}

	menu, err := GetMenuModel(id)
	if err != nil {
		controllers.SendErrorResponse(c, err)
		return nil, id, false
	}

	return menu, id, true
}
// Menu -
func MenuApi(t *Menu) map[string]interface{} {
	return map[string]interface{}{
		"id":                 t.ID,
		"menu_name":              t.MenuName,
		"menu_type":               t.MenuType,
		"is_default":            t.IsDefault,
		"status":        t.Status,
		"weight":        t.Weight,
		"created_at":         t.CreatedAt.Format(constants.DateTimeLayout),
		"updated_at":         t.UpdatedAt.Format(constants.DateTimeLayout),
	}
}

// BeforeSave - hook
func (t *Menu) BeforeSave() error {
	t.MenuName = utils.XSSClean(t.MenuName)

	return nil
}

// AfterSave - hook
func (t *Menu) AfterSave() error {
	// if t.Slug == "" {

	// }

	return nil
}

// BeforeDelete - hook
func (t *Menu) BeforeDelete(tx *gorm.DB) (err error) {

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
