package auth

import (
	"errors"
	userModel "dofun/app/models/user"
	"dofun/config"
	"dofun/pkg/ginutils"

	"github.com/gin-gonic/gin"
)

// SaveCurrentUserToContext : 保存用户数据到 context 中
func SaveCurrentUserToContext(c *gin.Context) *userModel.User {
	user, err := getCurrentUserFromSession(c)
	if err != nil {
		return nil
	}

	c.Keys[config.AppConfig.ContextCurrentUserDataKey] = user
	return user
}

// GetCurrentUserFromContext : 从 context 中获取用户模型
func GetCurrentUserFromContext(c *gin.Context) (*userModel.User, error) {
	err := errors.New("没有获取到用户数据")
	userDataFromContext := c.Keys[config.AppConfig.ContextCurrentUserDataKey]
	if userDataFromContext == nil {
		return nil, err
	}

	user, ok := userDataFromContext.(*userModel.User)
	if !ok {
		return nil, err
	}

	return user, nil
}

// GetUserFromContextOrDataBase : 从 context 或者从数据库中获取用户模型
func GetUserFromContextOrDataBase(c *gin.Context, id int) (*userModel.User, error) {
	// 当前用户存在并且就是想要获取的那个用户
	currentUser, err := GetCurrentUserFromContext(c)
	if currentUser != nil && err == nil {
		if int(currentUser.ID) == id {
			return currentUser, nil
		}
	}

	// 获取的是其他指定 id 的用户
	otherUser, err := userModel.Get(id)
	if err != nil {
		return nil, err
	}

	return otherUser, nil
}

// GetUserFromParamIDOrContext : params 中存在 id 参数时，可用该方法去获取 user model
func GetUserFromParamIDOrContext(c *gin.Context) (*userModel.User, error) {
	id, err := ginutils.GetIntParam(c, "id")
	if err != nil {
		return nil, err
	}

	return GetUserFromContextOrDataBase(c, id)
}
