package module

import (
	"strconv"
	"time"

	"github.com/ws-tobalobs/pkg/models"
)

func (u *notif) GetAllNotif(userID int64, tambakID int64, typeNotif string) ([]models.MessagePushNotif, error) {
	allNotif, err := u.mysqlNotifRepo.GetAllNotif(userID, tambakID, typeNotif)

	return allNotif, err
}

func (u *notif) GetDetailNotif(notifID int64) (models.Notifikasi, error) {
	notif, err := u.mysqlNotifRepo.GetDetailNotif(notifID)

	return notif, err
}

func (u *notif) PushNotif(userID, tambakID int64, typeNotif string) error {
	dt := time.Now()

	tambak, _ := u.tambakRepo.GetTambakByID(tambakID, userID)

	n := models.Notifikasi{
		GuidelineID:      1,
		NamaTambak:       tambak.NamaTambak,
		TambakID:         tambakID,
		TipeNotifikasi:   "notif-guideline",
		Keterangan:       "Beri Pakan Pagi Hari",
		StatusNotifikasi: "unread",
		WaktuTanggal:     dt.Format("2006-01-02 15:04:05"),
	}

	if typeNotif == "sore" {
		n.GuidelineID = 2
		n.Keterangan = "Beri Pakan Sore Hari"
	} else if typeNotif == "ganti-air" {
		n.GuidelineID = 3
		n.Keterangan = "Ganti Air Tambak"
	}

	nID, err := u.mysqlNotifRepo.SaveNotifGuideline(n)
	if err == nil {
		deviceIDs := u.redisNotifRepo.GetDeviceID(userID)
		if len(deviceIDs) == 0 {
			//if deviceID not exist in redis, update status notification to pending
			u.tambakRepo.UpdateNotifikasiKondisiTambak("pending", nID)
		} else {
			notifIDStr := strconv.FormatInt(nID, 10)
			msg := models.MessagePushNotif{
				ID:    notifIDStr,
				Title: n.NamaTambak,
				Body:  n.Keterangan,
			}
			u.fcmNotifRepo.PushNotification(deviceIDs, msg)
		}
	}

	return err
}

func (u *notif) SaveNotif(userID, tambakID int64, typeNotif string) (models.MessagePushNotif, error) {
	dt := time.Now()
	notif := models.MessagePushNotif{}

	tambak, _ := u.tambakRepo.GetTambakByID(tambakID, userID)

	n := models.Notifikasi{
		GuidelineID:      1,
		NamaTambak:       tambak.NamaTambak,
		TambakID:         tambakID,
		TipeNotifikasi:   "notif-guideline",
		Keterangan:       "Beri Pakan Pagi Hari",
		StatusNotifikasi: "unread",
		WaktuTanggal:     dt.Format("2006-01-02 15:04:05"),
	}

	if typeNotif == "sore" {
		n.GuidelineID = 2
		n.Keterangan = "Beri Pakan Sore Hari"
	} else if typeNotif == "ganti-air" {
		n.GuidelineID = 3
		n.Keterangan = "Ganti Air Tambak"
	}

	nID, err := u.mysqlNotifRepo.SaveNotifGuideline(n)
	if err != nil {
		return notif, err
	}

	notif.ID = strconv.FormatInt(nID, 10)
	notif.Title = n.NamaTambak
	notif.Body = n.Keterangan

	return notif, err
}
