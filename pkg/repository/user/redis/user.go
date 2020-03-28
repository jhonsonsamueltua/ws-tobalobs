package redis

import (
	"log"
	"time"
)

func (r *user) Logout(token string, exp int) error {
	err := r.redis.Set(token, token, (time.Duration(exp) * time.Second)).Err()
	if err != nil {
		log.Println("Repo Logout error : ", err)
	}
	return err
}

func (r *user) SaveDeviceID(key string, value string) {
	r.RemoveDeviceID(key, value)
	err := r.redis.LPush(key, value).Err()
	if err != nil {
		log.Println("Repo Save DeviceID error : ", err)
	}
}

func (r *user) RemoveDeviceID(key string, value string) {
	err := r.redis.LRem(key, -2, value).Err()
	if err != nil {
		log.Println("Repo Remove DeviceID error : ", err)
	}
}
