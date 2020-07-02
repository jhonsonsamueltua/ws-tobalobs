package module

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ws-tobalobs/pkg/models"
)

func (u *notif) GetAllNotif(userID int64, tambakID int64, typeNotif string) ([]models.MessagePushNotif, error) {
	var res []models.MessagePushNotif

	allNotif, err := u.mysqlNotifRepo.GetAllNotif(userID, tambakID, typeNotif)

	format := "2006-01-02 15:04:05"
	for _, n := range allNotif {
		dt, _ := time.Parse(format, n.WaktuTanggal)
		n.WaktuTanggal = dt.Format("2 Jan 2006 - 15:04")
		res = append(res, n)
	}

	return res, err
}

func (u *notif) GetDetailNotif(notifID int64) (models.Notifikasi, error) {
	notif, err := u.mysqlNotifRepo.GetDetailNotif(notifID)

	format := "2006-01-02 15:04:05"
	dt, _ := time.Parse(format, notif.WaktuTanggal)
	notif.WaktuTanggal = dt.Format("2 Jan 2006 - 15:04")

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
	loc, _ := time.LoadLocation("Asia/Jakarta")
	dt := time.Now().In(loc)
	notif := models.MessagePushNotif{}

	tambak, _ := u.tambakRepo.GetTambakByID(tambakID, userID)

	n := models.Notifikasi{
		GuidelineID:      1,
		NamaTambak:       tambak.NamaTambak,
		TambakID:         tambakID,
		TipeNotifikasi:   "notif-guideline",
		Keterangan:       "Beri Pakan Pagi Hari",
		StatusNotifikasi: "unread",
		WaktuTanggal:     fmt.Sprintf("%s %s", dt.Format("2006-01-02"), tambak.PakanPagi),
	}

	if typeNotif == "sore" {
		n.GuidelineID = 2
		n.Keterangan = "Beri Pakan Sore Hari"
		n.WaktuTanggal = fmt.Sprintf("%s %s", dt.Format("2006-01-02"), tambak.PakanSore)
	} else if typeNotif == "ganti-air" {
		n.GuidelineID = 3
		n.Keterangan = "Ganti Air Tambak"
		n.WaktuTanggal = fmt.Sprintf("%s 07:00", dt.Format("2006-01-02"))
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
