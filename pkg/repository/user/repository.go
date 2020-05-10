package user

import (
	"github.com/ws-tobalobs/pkg/models"
)

type Repository interface {
	Register(models.User) (int64, error)
	GetUser(username string) (models.User, error)
	GetDetailUser(userID int64) (models.User, error)
	UpdateUser(models.User) error
	UpdatePassword(newPass string, userID int64) error
}

type RepositoryRedis interface {
	SaveDeviceID(key string, value string)
	RemoveDeviceID(key string, value string)
	Logout(token string, exp int) error
}
