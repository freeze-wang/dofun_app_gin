package verification

import (
	"dofun/app/handlers"
	"dofun/app/http/controllers"
	auth2 "dofun/app/http/controllers/auth"
	userModel "dofun/app/models/user"
	"dofun/pkg/ginutils/flash"

	"github.com/gin-gonic/gin"
)

// Show 展示发送激活邮件的页面
func Show(c *gin.Context, currentUser *userModel.User) {
	if currentUser.IsActivated() {
		controllers.RedirectBack(c, "root")
	} else {
		controllers.Render(c, "auth/verify", gin.H{})
	}
}

// Verify 激活
func Verify(c *gin.Context) {
	token := c.Param("token")
	user, err := userModel.GetByActivationToken(token)
	if user == nil || err != nil {
		controllers.Render404(c)
		return
	}

	// 更新用户

	if err = user.Update(); err != nil {
		flash.NewSuccessFlash(c, "用户激活失败: "+err.Error())
		controllers.RedirectRouter(c, "verification.notice")
		return
	}

	auth2.Login(c, user)
	flash.NewSuccessFlash(c, "邮箱验证成功 ^_^")
	controllers.RedirectRouter(c, "root")
}

// Resend 重新发送激活邮件
func Resend(c *gin.Context, currentUser *userModel.User) {
	if currentUser.IsActivated() {
		controllers.RedirectBack(c, "root")
		return
	}

	if err := handlers.SendVerifyEmail(currentUser); err != nil {
		flash.NewDangerFlash(c, "邮件发送失败: "+err.Error())
	} else {
		flash.NewSuccessFlash(c, "新的验证链接已发送到您的 E-mail")
	}

	controllers.RedirectRouter(c, "verification.notice")
}
