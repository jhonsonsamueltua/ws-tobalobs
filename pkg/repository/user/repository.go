package user

import (
	"time"

	"github.com/ws-tobalobs/pkg/models"
)

type Repository interface {
	Register(models.User) (int64, error)
	GetUser(username string) (models.User, error)
	GetByPhoneNumber(hp string) (models.User, error)
	GetDetailUser(userID int64) (models.User, error)
	UpdateUser(models.User) error
	UpdatePassword(newPass string, userID int64) error
	//manage data dynamic
	GetKondisiMenyimpang() ([]models.KondisiMenyimpang, error)
}

type RepositoryRedis interface {
	SaveDeviceID(key string, value string)
	RemoveDeviceID(key string, value string)
	Logout(token string, exp int) error
	SetValue(key string, value string, expiry time.Duration) error
	GetValue(key string) (string, error)
}

type RepositorySMS interface {
	Sendmessage(toNumber string, otpMessage string) error
}
