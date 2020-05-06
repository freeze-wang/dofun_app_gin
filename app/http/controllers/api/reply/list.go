package reply

import (
	"dofun/app/http/controllers"
	replyModel "dofun/app/models/reply"
	topicModel "dofun/app/models/topic"
	userModel "dofun/app/models/user"
	"dofun/app/services"
	"dofun/pkg/errno"
	"dofun/pkg/ginutils"

	"github.com/gin-gonic/gin"
)

// TopicReplies topic 的 reply list
func TopicReplies(c *gin.Context) {
	topicID, err := ginutils.GetIntParam(c, "topic_id")
	if err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}

	topic, err := topicModel.Get(topicID)
	if err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.DatabaseError, err))
		return
	}

	// 获取回复
	replies, _ := services.RpleyListApiService(func() ([]*replyModel.Reply, error) {
		return replyModel.TopicReplies(int(topic.ID))
	})

	controllers.SendOKResponse(c, controllers.ListData{List: replies})
}

// UserReplies user 的 reply list
func UserReplies(c *gin.Context) {
	userID, err := ginutils.GetIntParam(c, "user_id")
	if err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}

	user, err := userModel.Get(userID)
	if err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.DatabaseError, err))
		return
	}

	// 获取回复
	replies, _ := services.RpleyListApiService(func() ([]*replyModel.Reply, error) {
		return replyModel.UserReplies(int(user.ID), 0, 0)
	})

	controllers.SendOKResponse(c, controllers.ListData{List: replies})
}
