package module

import (
	"log"
	"strconv"
	"time"

	c "github.com/robfig/cron"

	"github.com/ws-tobalobs/pkg/models"
)

func (u *cron) InitCron() {
	cr := c.New()
	cr.AddFunc("* 7 * * *", func() { u.CronNotifGuideline() })   //setiap jam 7:00
	cr.AddFunc("* 17 * * *", func() { u.CronNotifGuideline() })  //setiap jam 17:00
	cr.AddFunc("30 18 * * *", func() { u.CronNotifGuideline() }) //setiap jam 18:30

	// cr.AddFunc("0 7 * * *", func() { u.CronPakan("pagi") })
	// cr.AddFunc("0 17 * * *", func() { u.CronPakan("sore") })
	// cr.AddFunc("0 7 */3 * *", func() { log.Println("ganti air") })
	log.Println("start Cron...")
	cr.Start()
}

func (u *cron) CronPakan(waktu string) error {
	dt := time.Now()
	userID, tambakID, namaTambak, err := u.tambakRepo.GetAllTambakID()
	if err != nil {
		return err
	}

	n := models.Notifikasi{
		GuidelineID:      1,
		TipeNotifikasi:   "notif-guideline",
		Keterangan:       "Beri Pakan Pagi Hari",
		StatusNotifikasi: "unread",
		WaktuTanggal:     dt.Format("2006-01-02 15:04:05"),
	}

	if waktu == "sore" {
		n.GuidelineID = 2
		n.Keterangan = "Beri Pakan Sore Hari"
	}

	for i := 0; i < len(tambakID); i++ {
		n.TambakID = tambakID[i]
		n.NamaTambak = namaTambak[i]
		nID, err := u.mysqlNotifRepo.SaveNotifGuideline(n)
		if err != nil {
			continue
		}

		deviceIDs := u.redisNotifRepo.GetDeviceID(userID[i])
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

func (u *cron) CronNotifGuideline() error {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	notif, _ := u.mysqlNotifRepo.GetNotifWaiting(now.Format("2006-01-02 15:04:05"))
	if len(notif) > 0 {
		for i := 0; i < len(notif); i++ {
			deviceIDs := u.redisNotifRepo.GetDeviceID(notif[i].UserID)
			if len(deviceIDs) == 0 {
				//if deviceID not exist in redis, update status notification to pending
				u.tambakRepo.UpdateNotifikasiKondisiTambak("pending", notif[i].NotifikasiID)
			} else {
				notifIDStr := strconv.FormatInt(notif[i].NotifikasiID, 10)
				msg := models.MessagePushNotif{
					ID:    notifIDStr,
					Title: notif[i].NamaTambak,
					Body:  notif[i].Keterangan,
				}
				u.fcmNotifRepo.PushNotification(deviceIDs, msg)
				u.tambakRepo.UpdateNotifikasiKondisiTambak("unread", notif[i].NotifikasiID) //update status ke unread
			}
		}
	}

	return nil
}
