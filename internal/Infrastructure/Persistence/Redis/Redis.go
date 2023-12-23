package Redis

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/dig"
)

// type RedisProvider struct{}

func UseRedisProvider() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "neptune",
		DB:       0,
	})
}

func UseRedisRepositoryProviders(c *dig.Container) {

}
