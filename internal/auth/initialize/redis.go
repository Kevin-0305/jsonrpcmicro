package initialize

import (
	"fmt"
	"jsonrpcmicro/internal/auth/config"

	"github.com/go-redis/redis"
)

func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Conf.Cache.Hosts[0],
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("redis connect ping failed, err: " + err.Error())
		return nil
	} else {
		fmt.Println("redis connect ping response:")
		return client
	}
}
