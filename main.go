package main

import (
	"github.com/gin-gonic/gin"

	"github.com/spf13/pflag"
)
var (
	// 需要 mock data，注意该操作会覆盖数据库；只在非 release 时生效
	needMock = pflag.BoolP("mock", "m", false, "need mock data")
)

func main() {
	r := setupRouter()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}