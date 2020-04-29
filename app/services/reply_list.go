package services

import (
	replyModel "dofun/app/models/reply"
	topicModel "dofun/app/models/topic"
	userModel "dofun/app/models/user"
	"dofun/app/policies"
	"dofun/app/viewmodels"
	"dofun/database"
	"dofun/pkg/ginutils/utils"
)

// RpleyListService -
func RpleyListService(getReplyFunc func() ([]*replyModel.Reply, error), currentUser *userModel.User) (interface{}, error) {
	var (
		result = make([]interface{}, 0) // 最终结果

		replyIDs   = make([]uint, 0) // 存储所有 reply id (用于最后排序 topic)
		replyIDMap = make(map[uint]map[string]interface{})

		usersIDs = make([]uint, 0) // 存储所有 user id
		users    = make([]*userModel.User, 0)

		topicsIDs = make([]uint, 0) // 存储所有 topic id
		topics    = make([]*topicModel.Topic, 0)
	)

	// 获取 reply
	replies, err := getReplyFunc()
	if err != nil {
		return result, err
	}

	for _, r := range replies {
		replyIDMap[r.ID] = viewmodels.NewReplyViewModelSerializer(r)
		replyIDs = append(replyIDs, r.ID)
		usersIDs = append(usersIDs, r.UserID)
		topicsIDs = append(topicsIDs, r.TopicID)
	}

	// 获取 user 和 topic
	if err = database.DB.Where("id in (?)", utils.UniqueUintSlice(usersIDs)).Find(&users).Error; err != nil {
		return result, err
	}
	if err = database.DB.Where("id in (?)", utils.UniqueUintSlice(topicsIDs)).Find(&topics).Error; err != nil {
		return result, err
	}

	// 整理数据
	for _, r := range replies {
		for _, u := range users {
			if r.UserID == u.ID {
				replyIDMap[r.ID]["User"] = viewmodels.NewUserViewModelSerializer(u)
			}
		}
		for _, t := range topics {
			if r.TopicID == t.ID {
				replyIDMap[r.ID]["Topic"] = viewmodels.NewTopicViewModelSerializer(t)
				// 是否可删除
				if policies.CheckReplyPolicy(currentUser, r, t) {
					replyIDMap[r.ID]["CanDelete"] = true
				}
			}
		}

	}

	for _, id := range replyIDs {
		result = append(result, replyIDMap[id])
	}

	return result, nil
}

// RpleyListApiService -
func RpleyListApiService(getReplyFunc func() ([]*replyModel.Reply, error)) (interface{}, error) {
	var (
		result = make([]interface{}, 0) // 最终结果

		replyIDs   = make([]uint, 0) // 存储所有 reply id (用于最后排序 topic)
		replyIDMap = make(map[uint]map[string]interface{})

		usersIDs = make([]uint, 0) // 存储所有 user id
		users    = make([]*userModel.User, 0)

		topicsIDs = make([]uint, 0) // 存储所有 topic id
		topics    = make([]*topicModel.Topic, 0)
	)

	// 获取 reply
	replies, err := getReplyFunc()
	if err != nil {
		return result, err
	}

	for _, r := range replies {
		replyIDMap[r.ID] = viewmodels.ReplyApi(r)
		replyIDs = append(replyIDs, r.ID)
		usersIDs = append(usersIDs, r.UserID)
		topicsIDs = append(topicsIDs, r.TopicID)
	}

	// 获取 user 和 topic
	if err = database.DB.Where("id in (?)", utils.UniqueUintSlice(usersIDs)).Find(&users).Error; err != nil {
		return result, err
	}
	if err = database.DB.Where("id in (?)", utils.UniqueUintSlice(topicsIDs)).Find(&topics).Error; err != nil {
		return result, err
	}

	// 整理数据
	for _, r := range replies {
		for _, u := range users {
			if r.UserID == u.ID {
				replyIDMap[r.ID]["user"] = viewmodels.NewUserAPISerializer(u)
			}
		}

		for _, t := range topics {
			if r.TopicID == t.ID {
				replyIDMap[r.ID]["topic"] = viewmodels.TopicApi(t)
			}
		}
	}

	for _, id := range replyIDs {
		result = append(result, replyIDMap[id])
	}

	return result, nil
}
