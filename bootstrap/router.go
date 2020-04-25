package bootstrap

import (
	"dofun/config"
	"dofun/pkg/ginutils/router"
	"dofun/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter(g *gin.Engine) {
	routes.Register(g)
	printRoute()
}

// 打印命名路由
func printRoute() {
	// 只有非 release 时才可用该函数
	if config.AppConfig.RunMode == config.RunmodeRelease {
		return
	}

	router.PrintRoutes()
}
