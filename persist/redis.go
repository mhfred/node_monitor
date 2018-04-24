package persist

import (
	"fmt"
	"node_monitor/config"
	"github.com/go-redis/redis"
)


func newRedisClient() *redis.Client {
	host := config.Flags.RedisHost
	port := config.Flags.RedisPort
	addr := fmt.Sprintf("%s:%s", host, port)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}



