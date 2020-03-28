package module

import (
	"github.com/ws-tobalobs/pkg/models"
)

func (u *notif) GetAllNotif(userID int64, tambakID int64, typeNotif string) ([]models.MessagePushNotif, error) {
	allNotif, err := u.notifRepo.GetAllNotif(userID, tambakID, typeNotif)

	return allNotif, err
}

func (u *notif) GetDetailNotif(notifID int64) (models.Notifikasi, error) {
	notif, err := u.notifRepo.GetDetailNotif(notifID)

	return notif, err
}
