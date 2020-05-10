package main

import (
	"dofun/app/http/controllers/api/authorization"
	"dofun/app/http/controllers/api/dynamic"
	"dofun/app/http/middleware"
	"dofun/bootstrap"
	"dofun/config"
	"dofun/database"
	"github.com/gin-contrib/pprof"
	_ "github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

var (
	// 需要 mock data，注意该操作会覆盖数据库；只在非 release 时生效
	needMock = pflag.BoolP("mock", "m", false, "need mock data")
)

func main() {
	// 初始化配置
	config.InitConfig("", true)
	r := setupRouter()
	pprof.Register(r)	 // 性能监控
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	bootstrap.SetupGin(r)
	database.InitDB()
	/*r.Use(middleware.TokenAuth(),middleware.RateLimiter(1*time.Minute, 10))
	{
		r.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
	}*/
	vt := r.Group("api/v1/").Use(middleware.TokenRefresh())
	{
		vt.GET("index/dynamic/:id", dynamic.Index)
	}
	v1 := r.Group("api/v1/")
	{
		v1.GET("dynamic/detail/:id", dynamic.Detail)
	}
	r.POST("/login", authorization.Store)

	/*server := endless.NewServer(config.AppConfig.Addr, r)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}*/
	return r
}
