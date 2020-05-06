package login

import (
	"dofun/app/auth"
	"dofun/app/http/controllers"
	"dofun/pkg/ginutils/flash"

	userRequest "dofun/app/http/requests/user"

	"github.com/gin-gonic/gin"
)

// 展示登录页面
func ShowLoginForm(c *gin.Context) {
	controllers.Render(c, "auth/login", gin.H{})
}

// 登录
func Login(c *gin.Context) {
	// 验证参数并且获取用户
	userLoginForm := &userRequest.UserLoginForm{
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}
	ok, user := userLoginForm.ValidateAndGetUser(c)

	if !ok || user == nil {
		controllers.RedirectRouter(c, "login")
		return
	}

	auth.Login(c, user)
	controllers.RedirectRouter(c, "root")
}

// 登出
func Logout(c *gin.Context) {
	auth.Logout(c)
	flash.NewSuccessFlash(c, "您已成功退出！")
	controllers.RedirectRouter(c, "login")
}
