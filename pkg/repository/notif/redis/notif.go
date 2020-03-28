package redis

import (
	"fmt"
	"log"
)

func (r *redisNotif) GetDeviceID(userID int64) []string {
	key := fmt.Sprint("device:", userID)
	deviceID, err := r.redis.LRange(key, 0, -1).Result()
	if err != nil {
		log.Println("Get DeviceID err : ", err)
	}

	return deviceID
}
