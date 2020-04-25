package viewmodels

import (
	replyModel "dofun/app/models/reply"
	"dofun/pkg/constants"
	gintime "dofun/pkg/ginutils/time"
)

// NewReplyViewModelSerializer -
func NewReplyViewModelSerializer(r *replyModel.Reply) map[string]interface{} {
	return map[string]interface{}{
		"ID":        r.ID,
		"Content":   r.Content,
		"UserID":    r.UserID,
		"TopicID":   r.TopicID,
		"CreatedAt": gintime.SinceForHuman(r.CreatedAt),
	}
}

// ReplyApi -
func ReplyApi(r *replyModel.Reply) map[string]interface{} {
	return map[string]interface{}{
		"id":         r.ID,
		"content":    r.Content,
		"user_id":    r.UserID,
		"topic_id":   r.TopicID,
		"created_at": r.CreatedAt.Format(constants.DateTimeLayout),
		"updated_at": r.UpdatedAt.Format(constants.DateTimeLayout),
	}
}
