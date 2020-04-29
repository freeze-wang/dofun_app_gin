package viewmodels

import (
	topicModel "dofun/app/models/topic"
	"dofun/pkg/constants"
)

// NewTopicViewModelSerializer -
func NewTopicViewModelSerializer(t *topicModel.Topic) map[string]interface{} {
	return map[string]interface{}{
		"ID":              t.ID,
		"CreatedAt":       t.CreatedAt,
		"UpdatedAt":       t.UpdatedAt,
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
