package routes

import (
	_ "dofun/docs" // swagger docs
	"time"

	"dofun/pkg/ginutils/csrf"
	"dofun/pkg/ginutils/last"
	"dofun/pkg/ginutils/oldvalue"
	"dofun/pkg/ginutils/session"

	"dofun/app/http/middleware"
	"dofun/pkg/constants"
	"dofun/pkg/errno"
	"dofun/pkg/ginutils/router"

	"dofun/app/http/controllers"
	"dofun/config"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

const (
	// APIRoot -
	APIRoot = "/api"

	// AdminWebRoot -
	AdminWebRoot = "/admin"
	// AdminAPIRoot -
	AdminAPIRoot = "/admin-api"
)

// Register 注册路由和中间件
func Register(g *gin.Engine) *gin.Engine {
	// ---------------------------------- 注册全局中间件 ----------------------------------
	g.Use(gin.Recovery())
	//if config.AppConfig.RunMode != config.RunmodeRelease {
		g.Use(gin.Logger())
	//}
	g.Use(last.LastMiddleware()) // 记录上一次请求信息

	// ---------------------------------- 注册路由 ----------------------------------
	r := &router.MyRoute{Router: g}

	// +++++++++++++++++++ swagger +++++++++++++++++++
	// 需全局安装 go get -u github.com/swaggo/swag/cmd/swag 然后 swag init 生成文档
	if config.AppConfig.RunMode != config.RunmodeRelease {
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// +++++++++++++++++++ web +++++++++++++++++++
	registerWeb(r,
		// session
		session.SessionMiddleware(),
		// csrf
		csrf.Middleware(func(c *gin.Context, _ bool) {
			if constants.IsApiRequest(c) {
				controllers.SendErrorResponse(c, errno.SessionError)
			} else {
				controllers.Render403(c, "很抱歉！您的 Session 已过期，请刷新后再试一次。")
			}
		}),
		// 记忆上次表单提交的内容，消费即消失
		oldvalue.OldValueMiddleware(),
		// 中间件中会从 session 中获取到 current user model
		middleware.CurrentUserMiddleware(),
	)

	// +++++++++++++++++++ admin +++++++++++++++++++
	registerAdmin(r)

	// +++++++++++++++++++ api +++++++++++++++++++
	registerAPI(r,
		middleware.RateLimiter(1*time.Minute, 60), // 1 分钟最多 60 次请求
	)

	// ---------------------------------- error ----------------------------------
	g.NoRoute(func(c *gin.Context) {
		if constants.IsApiRequest(c) {
			controllers.SendErrorResponse(c, errno.NotFoundError)
		} else {
			controllers.Render404(c)
		}
	})
	g.NoMethod(func(c *gin.Context) {
		if constants.IsApiRequest(c) {
			controllers.SendErrorResponse(c, errno.NotFoundError)
		} else {
			controllers.Render404(c)
		}
	})

	return g
}
