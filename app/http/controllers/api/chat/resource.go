package chat

import (
	"dofun/app/cache/chat_server"
	"dofun/app/http/controllers"
	"dofun/app/services/grpcclient"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// Index topic list
func send(c *gin.Context) {

	// 获取参数
	appId,_ := strconv.Atoi(c.Param("appId"))
	userId := c.PostForm("userId")
	msgId := c.PostForm("msgId")
	message := c.PostForm("message")

	go func() {
		currentTime := uint64(time.Now().Unix())
		servers, err := chat_server.GetServerAll(currentTime)
		if err != nil {
			fmt.Println("给全体用户发消息", err)
			return
		}

		for _, server := range servers {
			grpcclient.SendMsgAll(server, msgId, uint32(appId), userId, "msg", message)
		}
	}()

	controllers.SendOKResponse(c, message)
	return
}
