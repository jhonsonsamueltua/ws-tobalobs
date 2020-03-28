package fcm

import (
	"firebase.google.com/go/messaging"

	fcmRepo "github.com/ws-tobalobs/pkg/repository/notif"
)

type FCM struct {
	fcm *messaging.Client
}

func InitFCMRepo(fcm *messaging.Client) fcmRepo.RepositoryFCM {
	return &FCM{
		fcm: fcm,
	}
}
