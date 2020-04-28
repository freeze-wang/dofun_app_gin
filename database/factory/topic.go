package factory

import (
	"dofun/app/models/category"
	"dofun/app/models/topic"
	"dofun/app/models/user"
	"fmt"
)

func topicFactory(uids, cids []uint) *topic.Topic {
	/*title := randomdata.Country(randomdata.FullCountry)
	paragraph := randomdata.Paragraph()
	excerpt := paragraph
	if len(excerpt) > 20 {
		excerpt = excerpt[:20]
	}
	ur := utils.RandInt(0, len(uids)-1)
	cr := utils.RandInt(0, len(cids)-1)
	randUID := uids[ur]
	randCID := cids[cr]

	return &topic.Topic{
		Title:      title,
		Body:       paragraph,
		Excerpt:    excerpt,
		UserID:     randUID,
		CategoryID: randCID,
	}*/
	return nil
}

// TopicTableSeeder -
func TopicTableSeeder(needCleanTable bool) {
	if needCleanTable {
		dropAndCreateTable(&topic.Topic{})
	}

	userIDs, _ := user.AllID()
	categoryIDs, _ := category.AllID()

	for i := 0; i < 30; i++ {
		topic := topicFactory(userIDs, categoryIDs)
		if err := topic.Create(); err != nil {
			fmt.Printf("mock topic error： %v\n", err)
		}
	}
}
