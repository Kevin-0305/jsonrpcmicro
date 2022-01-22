package initialize

import (
	"fmt"

	"github.com/go-redis/redis"
)

func Redis(address string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
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
