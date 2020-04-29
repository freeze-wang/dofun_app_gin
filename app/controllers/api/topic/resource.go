package topic

import (
	"dofun/app/controllers"
	topicModel "dofun/app/models/topic"
	userModel "dofun/app/models/user"
	"dofun/app/policies"
	request "dofun/app/requests/api/topic"
	"dofun/app/services"
	"dofun/app/viewmodels"
	"dofun/pkg/errno"
	"dofun/pkg/ginutils"

	"github.com/gin-gonic/gin"
)

// Index topic list
func Index(c *gin.Context) {
	controllers.SendListResponse(c, 20, nil,
		topicModel.Count,
		func(offset, limit, _, _ int) (interface{}, error) {
			return services.TopicListAPIService(func() ([]*topicModel.Topic, error) {
				return topicModel.List(offset, limit, c.DefaultQuery("order", "default"))
			})
		})
}

// UserIndex topic list
func UserIndex(c *gin.Context) {
	id, err := ginutils.GetIntParam(c, "user_id")
	if err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}

	user, err := userModel.Get(id)
	if err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ResourceNotFoundError, err))
		return
	}

	controllers.SendListResponse(c, 20, nil,
		func() (int, error) { return topicModel.CountByUserID(int(user.ID)) },
		func(offset, limit, _, _ int) (interface{}, error) {
			return services.TopicListAPIService(func() ([]*topicModel.Topic, error) {
				return topicModel.GetByUserID(int(user.ID), offset, limit, c.DefaultQuery("order", "default"))
			})
		})
}

// Show topic detail
func Show(c *gin.Context) {
	topic, _, ok := getTopic(c, nil)
	if !ok {
		return
	}

	controllers.SendOKResponse(c, viewmodels.TopicApi(topic))
}

// Store 发布 topic
func Store(c *gin.Context, currentUser *userModel.User, tokenString string) {
	var req request.Store
	if err := c.ShouldBind(&req); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}

	topic, err := req.Run(currentUser.ID)
	if err != nil {
		controllers.SendErrorResponse(c, err)
		return
	}

	controllers.SendOKResponse(c, viewmodels.TopicApi(topic))
}

// Update 修改 topic
func Update(c *gin.Context, currentUser *userModel.User, tokenString string) {
	topic, _, ok := getTopic(c, currentUser)
	if !ok {
		return
	}

	var req request.Update
	if err := c.ShouldBind(&req); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.ParamsError, err))
		return
	}

	if err := req.Run(topic); err != nil {
		controllers.SendErrorResponse(c, err)
		return
	}

	controllers.SendOKResponse(c, viewmodels.TopicApi(topic))
}

// Destroy 删除 topic
func Destroy(c *gin.Context, currentUser *userModel.User, tokenString string) {
	topic, id, ok := getTopic(c, currentUser)
	if !ok {
		return
	}

	// 权限
	if ok := policies.TopicPolicyOwner(c, currentUser, int(topic.ID)); !ok {
		return
	}

	if err := topicModel.Delete(id); err != nil {
		controllers.SendErrorResponse(c, errno.New(errno.DatabaseError, err))
		return
	}

	controllers.SendOKResponse(c, map[string]interface{}{
		"id": id,
	})
}
