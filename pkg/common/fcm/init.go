package fcm

import (
	"log"

	"github.com/appleboy/go-fcm"
)

func InitFCM(serverKey string) *fcm.Client {
	client, err := fcm.NewClient(serverKey)
	if err != nil {
		log.Println(err)
	}

	return client
}
