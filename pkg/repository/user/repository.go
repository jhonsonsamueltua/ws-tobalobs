package user

import (
	"github.com/ws-tobalobs/pkg/models"
)

type Repository interface {
	Register(models.User) (int64, error)
	GetUser(username string) (models.User, error)
}

type RepositoryRedis interface {
	SaveDeviceID(key string, value string)
	RemoveDeviceID(key string, value string)
	Logout(token string, exp int) error
}
