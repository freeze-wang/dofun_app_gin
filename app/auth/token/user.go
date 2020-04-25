package token

import (
	userModel "dofun/app/models/user"
	"dofun/pkg/errno"
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	tokenParamsKeyName          = "token"
	tokenHeaderKeyName          = "Authorization"
	tokenInHeaderIdentification = "Bearer"
)

// GetTokenFromRequest 从请求中获取 token
func GetTokenFromRequest(c *gin.Context) (string, *errno.Errno) {
	if token, ok := getTokenFromHeader(c); ok {
		return token, nil
	}
	if token, ok := getTokenFromParams(c); ok {
		return token, nil
	}

	return "", errno.TokenMissingError
}

// ParseAndGetUser 解析 token 获取 user
func ParseAndGetUser(c *gin.Context, token string) (*userModel.User, *errno.Errno) {
	claims, err := parseToken(token)
	if err == errno.TokenExpireError {
		token, _, _ = refresh(token)
	}

	if err != nil {
		return nil, err
	}

	user, e := userModel.Get(int(claims.UserID))
	if e != nil {
		return nil, errno.New(errno.DatabaseError, e)
	}

	c.Set(tokenHeaderKeyName+"User", user)
	c.Set(tokenHeaderKeyName+"Token", token)
	return user, nil
}

// GetTokenUserFromContext -
func GetTokenUserFromContext(c *gin.Context) (string, *userModel.User, bool) {
	user, ok := c.Get(tokenHeaderKeyName + "User")
	if !ok {
		return "", nil, false
	}
	t, ok := c.Get(tokenHeaderKeyName + "Token")
	if !ok {
		return "", nil, false
	}

	u, ok := user.(*userModel.User)
	s, ok := t.(string)
	if !ok {
		return "", nil, false
	}

	return s, u, true
}

// ---------------- private
func getTokenFromHeader(c *gin.Context) (string, bool) {
	header := c.Request.Header.Get(tokenHeaderKeyName)
	if header == "" {
		return "", false
	}

	var token string
	fmt.Sscanf(header, tokenInHeaderIdentification+" %s", &token)
	if token == "" {
		return "", false
	}
	return token, true
}

func getTokenFromParams(c *gin.Context) (string, bool) {
	token := c.Query(tokenParamsKeyName)
	if token != "" {
		return token, true
	}
	token = c.PostForm(tokenParamsKeyName)
	if token != "" {
		return token, true
	}

	return "", false
}
