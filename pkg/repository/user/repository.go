package user

import (
	"time"

	"github.com/ws-tobalobs/pkg/models"
)

type Repository interface {
	Register(models.User) (int64, error)
	Login(username string, password string) bool
	GetUser(username string) (models.User, error)
	Logout(token string, exp time.Duration) error
}

type RepositoryRedis interface {
	Register(models.User) (int64, error)
	Login(username string, password string) bool
	GetUser(username string) (models.User, error)
	Logout(token string, exp int) error
}
