package notif

import "github.com/ws-tobalobs/pkg/models"

type Usecase interface {
	GetAllNotif(userID int64, tambakID int64, typeNotif string) ([]models.MessagePushNotif, error)
	GetDetailNotif(notifID int64) (models.Notifikasi, error)
	PushNotif(userID, tambakID int64, typeNotif string) error
	SaveNotif(userID, tambakID int64, typeNotif string) (models.MessagePushNotif, error)
}
