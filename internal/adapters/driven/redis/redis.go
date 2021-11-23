package redis

import (
	"github.com/go-redis/redis/v8"
	"log"
)

func InitRedis(addr, user, pass string, db int) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: user,
		Password: pass,
		DB:       db,
	})
	if cli == nil {
		log.Fatal("can't initialize redis , check address and db name or user pass")
	}
	return cli
}
