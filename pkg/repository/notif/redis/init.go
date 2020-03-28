package redis

import (
	"github.com/go-redis/redis"

	redisRepo "github.com/ws-tobalobs/pkg/repository/notif"
)

type redisNotif struct {
	redis *redis.Client
}

func InitRedisRepo(r *redis.Client) redisRepo.RepositoryRedis {
	return &redisNotif{
		redis: r,
	}
}
