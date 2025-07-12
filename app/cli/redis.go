package cli

import (
	"sync"

	"github.com/go-redis/redis/v8"

	"codeup.aliyun.com/6829ea85516a9f85a08cb8c7/ad-services/ad-materials/app/log"
	"codeup.aliyun.com/6829ea85516a9f85a08cb8c7/ad-services/ad-materials/config"
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
