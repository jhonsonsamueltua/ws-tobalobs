package notif

import (
	"github.com/ws-tobalobs/pkg/models"
)

type RepositoryFCM interface {
	PushNotification(deviceID []string, msg models.MessagePushNotif)
}

type RepositoryRedis interface {
	GetDeviceID(userID int64) []string
}

type RepositoryMysql interface {
	GetAllNotif(userID int64, tambakID int64, typeNotif string) ([]models.MessagePushNotif, error)
	GetDetailNotif(notifID int64) (models.Notifikasi, error)
	UpdateStatusNotifikasi(notifID int64)
	GetTotalNofikasiUnread(userID int64) int
	SaveNotifGuideline(n models.Notifikasi) (int64, error)
	GetNotifWaiting(waktu string) ([]models.Notifikasi, error)
}
