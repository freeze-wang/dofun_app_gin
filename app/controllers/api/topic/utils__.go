package topic

import (
	"dofun/app/controllers"
	topicModel "dofun/app/models/topic"
	userModel "dofun/app/models/user"
	"dofun/app/policies"
	"dofun/pkg/errno"
	"dofun/pkg/ginutils"

	"github.com/gin-gonic/gin"
)

// 获取要编辑的 topic
func getTopic(c *gin.Context, currentUser *userModel.User) (*topicModel.Topic, int, bool) {
	id, err := ginutils.GetIntParam(c, "id")
	if err != nil {
		controllers.SendErrorResponse(c, errno.Base(errno.ParamsError, "id 不存在"))
		return nil, id, false
	}

	topic, err := topicModel.Get(id)
	if err != nil {
		controllers.SendErrorResponse(c, errno.ResourceNotFoundError)
		return nil, id, false
	}

	// 权限
	if currentUser != nil {
		if ok := policies.TopicPolicyOwner(c, currentUser, int(topic.UserID)); !ok {
			return nil, id, false
		}
	}

	return topic, id, true
}
