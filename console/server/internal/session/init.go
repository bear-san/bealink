package session

import (
	"github.com/redis/go-redis/v9"
	"os"
)

var client *redis.Client

const sessionKeyFormat = "bealink-session-%s"

func init() {
	client = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
	})
}
