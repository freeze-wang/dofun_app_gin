package user

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"
)

const (
	// TableName 表名
	TableName = "df_user"
)

var (
	userCache = cache.New(30*time.Minute, 1*time.Hour)
)

// User 用户模型
type User struct {
	ID             uint `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;not null" binding:"required"`
	UserType       string `json:"user_type" gorm:"column:user_type;not null" binding:"required"`
	ZhwId          int `json:"zhw_id" gorm:"column:zhw_id;not null" binding:"required"`
	ZhwUsername    string `json:"zhw_username" gorm:"column:zhw_username;not null" binding:"required"`
	DjIdentify     string `json:"dj_identify" gorm:"column:dj_identify;not null" binding:"required"`
	ChannelNumber  string `json:"channel_number" gorm:"column:channel_number;not null" binding:"required"`
	BeansBalance   float64 `json:"beans_balance" gorm:"column:beans_balance;not null" binding:"required"`
	IntegralTotal  int `json:"integral_total" gorm:"column:integral_total;not null" binding:"required"`
	Phone          string `json:"phone" gorm:"column:phone;not null" binding:"required"`
	Avatar         string `json:"avatar" gorm:"column:avatar;not null" binding:"required"`
	Nickname       string `json:"nickname" gorm:"column:nickname;not null" binding:"required"`
	Gender         string `json:"gender" gorm:"column:gender;not null" binding:"required"`
	Email          string `json:"email" gorm:"column:email;not null" binding:"required"`
	Salt           string `json:"salt" gorm:"column:salt;not null" binding:"required"`
	Password       string `json:"password" gorm:"column:password;not null" binding:"required"`
	QqUnionid      string `json:"qq_unionid" gorm:"column:qq_unionid;not null" binding:"required"`
	WeixinUnionid  string `json:"weixin_unionid" gorm:"column:weixin_unionid;not null" binding:"required"`
	RegisterIp     string `json:"register_ip" gorm:"column:register_ip;not null" binding:"required"`
	LastToken      string `json:"last_token" gorm:"column:last_token;not null" binding:"required"`
	Status         int `json:"status" gorm:"column:status;not null" binding:"required"`
	IsBlack        int `json:"is_black" gorm:"column:is_black;not null" binding:"required"`
	InviteCode     string `json:"invite_code" gorm:"column:invite_code;not null" binding:"required"`
	RegisterSource string `json:"register_source" gorm:"column:register_source;not null" binding:"required"`
	ProductSource  string `json:"product_source" gorm:"column:product_source;not null" binding:"required"`
	Sign           string `json:"sign" gorm:"column:sign;not null" binding:"required"`
	ZhwLoginData   string `json:"zhw_login_data" gorm:"column:zhw_login_data;not null" binding:"required"`
	CreatedAt      time.Time `json:"created_at" gorm:"column:created_at;not null" binding:"required"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"column:updated_at;not null" binding:"required"`
	DeletedAt      time.Time `json:"deleted_at" gorm:"column:deleted_at;not null" binding:"required"`
}

// TableName 表名
func (User) TableName() string {
	return TableName
}

// BeforeCreate - hook
func (u *User) BeforeCreate() (err error) {
	if u.Password != "" {
		if isEncrypted := passwordEncrypted(u.Password); !isEncrypted {
			if err = u.Encrypt(); err != nil {
				return errors.New("User Model 创建失败: passwordEncrypted")
			}
		}
	}


	// 生成用户头像
	if u.Avatar == "" {
		hash := md5.Sum([]byte(u.Email))
		u.Avatar = "http://www.gravatar.com/avatar/" + hex.EncodeToString(hash[:])
	}

	return err
}

// BeforeUpdate - hook
func (u *User) BeforeUpdate() (err error) {
	if isEncrypted := passwordEncrypted(u.Password); !isEncrypted {
		if err = u.Encrypt(); err != nil {
			return errors.New("User Model 更新失败")
		}
	}

	return
}

// BeforeDelete - hook
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	// 当用户删除时，删除其发布的话题
	tx.Exec("delete from topics where user_id = ?", u.ID)
	// 当用户删除时，删除其发布的回复
	tx.Exec("delete from replies where user_id = ?", u.ID)

	return
}

// ------------ private
func passwordEncrypted(pwd string) (status bool) {
	return len(pwd) == 60 // 长度等于 60 说明加密过了
}

func setToCache(user *User) {
	key := strconv.Itoa(int(user.ID))
	userCache.Set(key, user, cache.DefaultExpiration)
}

func getFromCache(id int) (*User, bool) {
	cachedUser, ok := userCache.Get(strconv.Itoa(id))
	if !ok {
		return nil, false
	}

	u, ok := cachedUser.(*User)
	if !ok {
		return nil, false
	}

	return u, true
}

func delCache(id int) {
	userCache.Delete(strconv.Itoa(id))
}
