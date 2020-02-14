package user

import "github.com/ws-tobalobs/pkg/models"

type Usecase interface {
	Register(models.User) (string, error)
	Login(username string, password string) (string, error)
	Logout(token string) error
}
