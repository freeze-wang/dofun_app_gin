package policies

import (
	replyModel "dofun/app/models/reply"
	topicModel "dofun/app/models/topic"
	userModel "dofun/app/models/user"

	"github.com/gin-gonic/gin"
)

// ReplyPolicy : 是否有更新、删除 reply 的权限
func ReplyPolicy(c *gin.Context, currentUser *userModel.User, reply *replyModel.Reply, topic *topicModel.Topic) bool {
	if CheckReplyPolicy(currentUser, reply, topic) {
		return true
	}

	Unauthorized(c)
	return false
}

// CheckReplyPolicy -
func CheckReplyPolicy(currentUser *userModel.User, reply *replyModel.Reply, topic *topicModel.Topic) bool {
	if currentUser == nil {
		return false
	}
	if before(currentUser) {
		return true
	}

	if reply.UserID == currentUser.ID || topic.UserID == currentUser.ID {
		return true
	}

	return false
}
