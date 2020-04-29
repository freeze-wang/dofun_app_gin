package viewmodels

import (
	topicModel "dofun/app/models/topic"
	"dofun/pkg/constants"
	gintime "dofun/pkg/ginutils/time"
)

// NewTopicViewModelSerializer -
func NewTopicViewModelSerializer(t *topicModel.Topic) map[string]interface{} {
	return map[string]interface{}{
		"ID":              t.ID,
		"CreatedAt":       gintime.SinceForHuman(t.CreatedAt),
		"UpdatedAt":       gintime.SinceForHuman(t.UpdatedAt),
	}
}

// Topic -
func TopicApi(t *topicModel.Topic) map[string]interface{} {
	return map[string]interface{}{
		"id":                 t.ID,
		"created_at":         t.CreatedAt.Format(constants.DateTimeLayout),
		"updated_at":         t.UpdatedAt.Format(constants.DateTimeLayout),
	}
}
