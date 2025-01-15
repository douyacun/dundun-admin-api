package cli

import (
	"sync"

	"github.com/go-redis/redis/v8"

	"github.com/douyacun/go-websocket-protobuf-ts/app/log"
	"github.com/douyacun/go-websocket-protobuf-ts/config"
)

var (
	RDB       *redis.Client
	redisOnce sync.Once
)

func InitRDB() {
	if config.Client.RedisAddr() != "" {
		log.Errorf("InitRDB redis_addr empty !!!")
		return
	}
	redisOnce.Do(func() {
		RDB = redis.NewClient(&redis.Options{
			Addr:     config.Client.RedisAddr(),
			Password: config.Client.RedisPassword(), // no password set
			DB:       config.Client.RedisDB(),       // use default DB
		})
	})
}
