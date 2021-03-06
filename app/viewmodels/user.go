package viewmodels

import (
	"dofun/app/handlers"
	permissionModel "dofun/app/models/permission"
	userModel "dofun/app/models/user"
	"dofun/pkg/constants"
	gintime "dofun/pkg/ginutils/time"
)

// UserViewModel 用户
type UserViewModel struct {
	ID                int
	Name              string
	Email             string
	Avatar            string
	Introduction      string
	CreatedAt         string
	LastActivedAt     string
	NotificationCount int
}

// NewUserViewModelSerializer 用户数据展示
func NewUserViewModelSerializer(u *userModel.User) *UserViewModel {
	data := &UserViewModel{
		ID:                int(u.ID),
		Name:              u.Nickname,
		Email:             u.Email,
		Avatar:            u.Avatar,
		Introduction:      u.Sign,
		CreatedAt:         gintime.SinceForHuman(u.CreatedAt.Time),
	}
	t := handlers.GetUserActivedLastActivedAt(u)
	if t != nil {
		data.LastActivedAt = gintime.SinceForHuman(*t)
	}

	return data
}

// NewUserAPISerializer api data
func NewUserAPISerializer(u *userModel.User) map[string]interface{} {
	r := map[string]interface{}{
		"id":           u.ID,
		"name":         u.Nickname,
		"email":        u.Email,
		"avatar":       u.Avatar,
		"introduction": u.Sign,
		"bound_phone":  false,
		"bound_wechat": false,
		"created_at":   u.CreatedAt.Format(constants.DateTimeLayout),
		"updated_at":   u.UpdatedAt.Format(constants.DateTimeLayout),
	}

	t := handlers.GetUserActivedLastActivedAt(u)
	if t != nil {
		r["last_actived_at"] = t.Format(constants.DateTimeLayout)
	}
	if u.Phone != "" {
		r["bound_phone"] = true
	}
	if u.WeixinUnionid != "" {
		r["bound_wechat"] = true
	}

	return r
}

// NewUserAPIHasRoles -
func NewUserAPIHasRoles(u *userModel.User, rs []*permissionModel.Role) map[string]interface{} {
	uvm := NewUserAPISerializer(u)
	uvm["roles"] = RoleAPIList(rs)

	return uvm
}
