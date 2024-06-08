package session

import (
	"github.com/redis/go-redis/v9"
	"os"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
	})
}
