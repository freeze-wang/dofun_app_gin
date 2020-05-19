/**
* Created by GoLand.
* User: link1st
* Date: 2019-08-03
* Time: 15:23
 */

package chat_server

import (
	"dofun/pkg/gredis"
	"fmt"
	"dofun/app/models/chat"
	"strconv"
)

const (
	serversHashKey       = "acc:hash:servers" // 全部的服务器
	serversHashCacheTime = 2 * 60 * 60        // key过期时间
	serversHashTimeout   = 3 * 60             // 超时时间
)

func getServersHashKey() (key string) {
	key = fmt.Sprintf("%s", serversHashKey)

	return
}

func GetServerAll(currentTime uint64) (servers []*chat.Server, err error) {

	servers = make([]*chat.Server, 0)
	key := getServersHashKey()

	/*redisClient := redislib.GetClient()

	val, err := redisClient.Do("hGetAll", key).Result()

	valByte, _ := json.Marshal(val)
	fmt.Println("GetServerAll", key, string(valByte))*/

	serverMap, err := gredis.HGetAll(key)
	if err != nil {
		fmt.Println("SetServerInfo", key, err)

		return
	}

	for key, value := range serverMap {
		valueUint64, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			fmt.Println("GetServerAll", key, err)

			return nil, err
		}

		// 超时
		if valueUint64+serversHashTimeout <= currentTime {
			continue
		}

		server, err := chat.StringToServer(key)
		if err != nil {
			fmt.Println("GetServerAll", key, err)

			return nil, err
		}

		servers = append(servers, server)
	}

	return
}
