package redis

import (
	"log"
	"time"

	"github.com/ws-tobalobs/pkg/models"
)

func (r *user) Logout(token string, exp int) error {
	err := r.redis.Set(token, token, (time.Duration(exp) * time.Second)).Err()
	if err != nil {
		log.Println("Usecase Logout error : ", err)
	}
	return err
}

func (r *user) Register(m models.User) (int64, error) {
	//not implement
	return 0, nil
}

func (r *user) Login(username string, password string) bool {
	//not implement
	return true
}

func (r *user) GetUser(username string) (models.User, error) {
	//not implement
	return models.User{}, nil
}
