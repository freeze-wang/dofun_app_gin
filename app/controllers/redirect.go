package controllers

import (
	"dofun/config"
	"dofun/pkg/ginutils"
	"dofun/pkg/ginutils/router"

	"dofun/pkg/ginutils/last"

	"github.com/gin-gonic/gin"
)

// Redirect : 路由重定向 use path
func Redirect(c *gin.Context, redirectPath string, withRoot bool) {
	path := redirectPath
	if withRoot {
		path = config.AppConfig.URL + redirectPath
	}

	ginutils.Redirect(c, path)
}

// RedirectRouter : 路由重定向 use router name
func RedirectRouter(c *gin.Context, routerName string, args ...interface{}) {
	ginutils.Redirect(c, router.G(routerName, args...))
}

// RedirectBack : 重定向到之前 path
func RedirectBack(c *gin.Context, defaultRouter string, args ...interface{}) {
	lastData := last.Read(c)

	if lastData != nil && lastData.Path != "" && c.Request.URL.Path != lastData.Path {
		Redirect(c, lastData.Path, false)
	} else {
		RedirectRouter(c, defaultRouter, args...)
	}
}
