package redis

import (
	"github.com/souvikhaldar/gobudgetsms"

	userRepo "github.com/ws-tobalobs/pkg/repository/user"
)

type user struct {
	redis *redis.Client
}

func InitUserRepoRedis(redis *redis.Client) userRepo.RepositoryRedis {
	return &user{
		redis: redis,
	}
}
