package persist

import (
	"node_monitor/model"
	"time"
)

const ChainInfoKey = "chainInfo"


func SaveChainInfo(chainInfo *model.ChainInfo) {
	redisClient := newRedisClient()
	defer redisClient.Close()

	redisClient.Set(ChainInfoKey,chainInfo.String(),time.Duration(1*time.Second)).Result()
}


func GetChainInfo() string {
	redisClient := newRedisClient()
	defer redisClient.Close()

	result, err := redisClient.Get(ChainInfoKey).Result()
	if err != nil {
		return "{}"
	}

	return result
}
