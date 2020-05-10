package user

import "github.com/ws-tobalobs/pkg/models"

type Usecase interface {
	Register(models.User) (string, error)
	Login(username string, password string, deviceID string) (string, error)
	Logout(token, deviceID string, userID int64) error
	GetDetailUser(userID int64) (models.User, error)
	UpdateUser(models.User) error
	UpdatePassword(pass, newPass string, userID int64) error
}
