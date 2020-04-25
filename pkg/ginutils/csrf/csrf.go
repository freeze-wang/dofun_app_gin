package csrf

import (
	"fmt"
	"dofun/pkg/ginutils"
	"dofun/pkg/ginutils/utils"
	"html/template"

	"github.com/gin-gonic/gin"
)

// CsrfInput csrf input html
func CsrfInput(c *gin.Context) (template.HTML, string, bool) {
	inputName := ginutils.GetGinUtilsConfig().CsrfParamName
	token := c.Keys[inputName]
	tokenStr, ok := token.(string)
	if !ok {
		return "", "", false
	}

	return template.HTML(fmt.Sprintf(`<input type="hidden" name="%s" value="%s">`, inputName, tokenStr)), tokenStr, true
}

// CsrfMeta csrf meta html
func CsrfMeta(c *gin.Context) (template.HTML, string, bool) {
	inputName := ginutils.GetGinUtilsConfig().CsrfParamName
	token := c.Keys[inputName]
	tokenStr, ok := token.(string)
	if !ok {
		return "", "", false
	}

	return template.HTML(fmt.Sprintf(`<meta name="csrf-token" content="%s">`, tokenStr)), tokenStr, true
}

// 从 cookie 中获取 csrf token
func getCsrfTokenFromCookie(c *gin.Context) (token string) {
	keyName := ginutils.GetGinUtilsConfig().CsrfParamName

	if s, err := c.Request.Cookie(keyName); err == nil {
		token = s.Value
	}

	if token == "" {
		token = string(utils.RandomCreateBytes(32))
		c.SetCookie(keyName, token, 0, "/", "", false, true)
	}
	c.Keys[keyName] = token

	return token
}

// 从 params 或 headers 中获取 csrf token
func getCsrfTokenFromParamsOrHeader(c *gin.Context) (token string, inHeader bool) {
	req := c.Request

	// 从 params 中获取
	token = req.PostFormValue(ginutils.GetGinUtilsConfig().CsrfParamName)
	if token == "" {
		// 从 headers 中获取
		token = req.Header.Get(ginutils.GetGinUtilsConfig().CsrfHeaderName)
		if token != "" {
			inHeader = true
		}
	}

	return
}
