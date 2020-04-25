package routes

import (
	"dofun/pkg/ginutils/captcha"
	"dofun/pkg/ginutils/router"
	"dofun/routes/middleware"
	"dofun/routes/wrapper"

	// "dofun/app/controllers/page"
	"dofun/app/controllers/auth/login"
	"dofun/app/controllers/auth/password"
	"dofun/app/controllers/auth/register"
	"dofun/app/controllers/auth/verification"
	"dofun/app/controllers/category"
	"dofun/app/controllers/notification"

	"dofun/app/controllers/reply"
	"dofun/app/controllers/topic"
	"dofun/app/controllers/user"
	"time"

	"github.com/gin-gonic/gin"
)

func registerWeb(r *router.MyRoute, middlewares ...gin.HandlerFunc) {
	r = r.Middleware(middlewares...)

	// r.Register("GET", "root", "", page.Root)
	r.Register("GET", "root", "", topic.Index)
	r.Register("GET", "captcha", "captcha/:id", captcha.Handler) // 验证码

	// ------------------------------------- Auth -------------------------------------
	// +++++++++++++++ 用户身份验证相关的路由 +++++++++++++++
	// 展示登录页面
	r.Register("GET", "login.show", "login", middleware.Guest(), login.ShowLoginForm)
	// 登录
	r.Register("POST", "login", "login", middleware.Guest(), login.Login)
	// 登出
	r.Register("POST", "logout", "logout", login.Logout)

	// +++++++++++++++ 用户注册相关路由 +++++++++++++++
	// 展示注册页面
	r.Register("GET", "register.show", "register", middleware.Guest(), register.ShowRegistrationForm)
	// 注册
	r.Register("POST", "register", "register", middleware.Guest(), register.Register)

	// +++++++++++++++ 密码重置相关路由 +++++++++++++++
	pwdRouter := r.Group("/password", middleware.Guest())
	{
		// 展示发送重置密码链接 email 的页面
		pwdRouter.Register("GET", "password.request", "/reset", password.ShowLinkRequestForm)
		// 发送重置密码链接的 email
		pwdRouter.Register("POST", "password.email", "/email", password.SendResetLinkEmail)
		// 展示重置密码的页面
		pwdRouter.Register("GET", "password.reset", "/reset/:token", password.ShowResetForm)
		// 重置密码
		pwdRouter.Register("POST", "password.update", "/reset", password.Reset)
	}

	// +++++++++++++++ Email 认证相关路由 +++++++++++++++
	verificationRouter := r.Group("/email", middleware.Auth())
	{
		// 展示发送激活用户链接邮件的页面
		// controller 中获取当前用户可使用 wrapper.GetUser 注入，或者使用 app/auth 里面的方法从 gin.Context 中获取
		verificationRouter.Register("GET", "verification.notice", "/verify", wrapper.GetUser(verification.Show))
		// 激活用户
		verificationRouter.Register("GET", "verification.verify", "/verify/:token",
			middleware.RateLimiter(1*time.Minute, 6), // 1 分钟最多 6 次请求
			verification.Verify)
		// 重新发送激活用户链接
		verificationRouter.Register("GET", "verification.resend", "/resend",
			middleware.RateLimiter(1*time.Minute, 6),
			wrapper.GetUser(verification.Resend))
	}

	// ------------------------------------- User -------------------------------------
	userRouter := r.Group("/users")
	{
		// 显示用户个人信息页面
		userRouter.Register("GET", "users.show", "/show/:id", user.Show)
		// 显示编辑个人资料页面
		userRouter.Register("GET", "users.edit", "/edit/:id", middleware.Auth(), wrapper.GetUser(user.Edit))
		// 处理 edit 页面提交的更改
		userRouter.Register("POST", "users.update", "/update/:id", middleware.Auth(), wrapper.GetUser(user.Update))
	}

	// ------------------------------------- topic -------------------------------------
	topicRouter := r.Group("/topics")
	{
		topicRouter.Register("GET", "topics.index", "", topic.Index)
		topicRouter.Register("GET", "topics.show_no_slug", "/show/:id", topic.Show)
		topicRouter.Register("GET", "topics.show", "/show/:id/*slug", topic.Show)
		topicRouter.Register("GET", "topics.create", "/create", middleware.Auth(), wrapper.GetUser(topic.Create))
		topicRouter.Register("POST", "topics.store", "", middleware.Auth(), wrapper.GetUser(topic.Store))
		topicRouter.Register("GET", "topics.edit", "/edit/:id", middleware.Auth(), wrapper.GetUser(topic.Edit))
		topicRouter.Register("POST", "topics.update", "/update/:id", middleware.Auth(), wrapper.GetUser(topic.Update))
		topicRouter.Register("POST", "topics.destroy", "/destroy/:id", middleware.Auth(), wrapper.GetUser(topic.Destroy))

		topicRouter.Register("POST", "topics.upload_image", "/upload_image", middleware.Auth(), wrapper.GetUser(topic.UploadImage))
	}

	// ------------------------------------- category -------------------------------------
	catRouter := r.Group("/categories")
	{
		catRouter.Register("GET", "categories.show", "/show/:id", category.Show)
	}

	// ------------------------------------- reply -------------------------------------
	replyRouter := r.Group("/replies", middleware.Auth())
	{
		replyRouter.Register("POST", "replies.store", "", wrapper.GetUser(reply.Store))
		replyRouter.Register("POST", "replies.destroy", "/destroy/:id", wrapper.GetUser(reply.Destroy))
	}

	// ------------------------------------- notification -------------------------------------
	notificationRouter := r.Group("/notifications", middleware.Auth())
	{
		notificationRouter.Register("GET", "notifications.index", "", wrapper.GetUser(notification.Index))
	}
}
