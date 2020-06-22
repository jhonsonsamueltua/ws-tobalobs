package module

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"

	"github.com/ws-tobalobs/pkg/models"
)

func (u *tambak) GetAllTambak(userID int64) ([]models.Tambak, int, error) {
	totalNotif := int(0)
	allTambak, err := u.tambakRepo.GetAllTambak(userID)
	if err == nil {
		//get total notif unread
		totalNotif = u.mysqlNotifRepo.GetTotalNofikasiUnread(userID)
	}
	return allTambak, totalNotif, err
}

func (u *tambak) GetTambakByID(tambakID int64, userID int64) (models.Tambak, error) {
	tambak, err := u.tambakRepo.GetTambakByID(tambakID, userID)

	return tambak, err
}

func (u *tambak) GetLastMonitorTambak(tambakID int64) (models.MonitorTambak, error) {
	monitor, err := u.tambakRepo.GetLastMonitorTambak(tambakID)

	return monitor, err
}

func (u *tambak) CreateTambak(t models.Tambak) (int64, error) {
	tambakID, err := u.tambakRepo.CreateTambak(t)
	if err == nil {
		// remote raspberry
		// execute(tambakID)

		//store notif guideline to db
		loc, _ := time.LoadLocation("Asia/Jakarta")
		now := time.Now().In(loc)

		var pemindahanInduk, pemisahanInduk, pemberianPakanSayur, pemberianPakanKeong, panenBenih, panenKonsumsi time.Time
		n := models.Notifikasi{
			TambakID:         tambakID,
			TipeNotifikasi:   "notif-guideline",
			StatusNotifikasi: "waiting",
		}

		if t.JenisBudidaya == "pembenihan" {
			//pemindahan induk
			pemindahanInduk = now.AddDate(0, 0, 14)
			n.GuidelineID = 4
			n.Keterangan = "Pemindahan induk"
			n.WaktuTanggal = pemindahanInduk.Format("2006-01-02 07:00:00")
			_, err := u.mysqlNotifRepo.SaveNotifGuideline(n)
			if err != nil {
				log.Println(err)
			}
			//pemisahan induk
			pemisahanInduk = pemindahanInduk.AddDate(0, 0, 45)
			n.GuidelineID = 5
			n.Keterangan = "Pemisahan induk dari anakan (telur menetas dan turun anak)"
			n.WaktuTanggal = pemisahanInduk.Format("2006-01-02 07:00:00")
			_, err = u.mysqlNotifRepo.SaveNotifGuideline(n)
			if err != nil {
				log.Println(err)
			}
			//panen benih
			panenBenih = pemindahanInduk.AddDate(0, 2, 0)
			n.GuidelineID = 8
			n.Keterangan = "Panen benih lobster untuk benih"
			n.WaktuTanggal = fmt.Sprintf("%s 18:30:00", panenBenih.Format("2006-01-02"))
			_, err = u.mysqlNotifRepo.SaveNotifGuideline(n)
			if err != nil {
				log.Println(err)
			}

		} else {
			//pemberian pakan sayuran
			lamaBudidaya := 8 - t.UsiaLobster
			for i := 1; i <= (lamaBudidaya * 2); i++ {
				if i == 1 {
					pemberianPakanSayur = now.AddDate(0, 0, 14)
				} else {
					pemberianPakanSayur = pemberianPakanSayur.AddDate(0, 0, 14)
				}
				n.GuidelineID = 6
				n.Keterangan = "Pemberian pakan sayuran"
				n.WaktuTanggal = pemberianPakanSayur.Format("2006-01-02 07:00:00")
				_, err := u.mysqlNotifRepo.SaveNotifGuideline(n)
				if err != nil {
					log.Println(err)
				}
			}

			//pemberian pakan keong
			for i := 1; i <= lamaBudidaya; i++ {
				if i == 1 {
					pemberianPakanKeong = now.AddDate(0, 1, 0)
				} else {
					pemberianPakanKeong = pemberianPakanKeong.AddDate(0, 1, 0)
				}
				n.GuidelineID = 7
				n.Keterangan = "Pemberian pakan keong"
				n.WaktuTanggal = pemberianPakanKeong.Format("2006-01-02 07:00:00")
				_, err := u.mysqlNotifRepo.SaveNotifGuideline(n)
				if err != nil {
					log.Println(err)
				}
			}

			//panen konsumsi
			panenKonsumsi = now.AddDate(0, lamaBudidaya, 0)
			n.GuidelineID = 8
			n.Keterangan = "Panen lobster pembesaran"
			n.WaktuTanggal = fmt.Sprintf("%s 18:30:00", panenKonsumsi.Format("2006-01-02"))
			_, err = u.mysqlNotifRepo.SaveNotifGuideline(n)
			if err != nil {
				log.Println(err)
			}

		}
	}

	return tambakID, err
}

func (u *tambak) UpdateTambak(t models.Tambak) error {
	err := u.tambakRepo.UpdateTambak(t)

	return err
}

func (u *tambak) PostMonitorTambak(m models.MonitorTambak) (int64, error) {
	monitorTambakId, err := u.tambakRepo.PostMonitorTambak(m)

	return monitorTambakId, err
}

func (u *tambak) PostPenyimpanganKondisiTambak(n models.Notifikasi) error {
	notifID, err := u.tambakRepo.PostPenyimpanganKondisiTambak(n)
	userID := u.tambakRepo.GetUserIDByTambak(n.TambakID)
	if err == nil {
		deviceIDs := u.redisNotifRepo.GetDeviceID(userID)
		if len(deviceIDs) == 0 {
			//if deviceID not exist in redis, update status notification to pending
			u.tambakRepo.UpdateNotifikasiKondisiTambak("pending", notifID)
		} else {
			notifIDStr := strconv.FormatInt(notifID, 10)
			msg := models.MessagePushNotif{
				ID:    notifIDStr,
				Title: "Notifikasi Kondisi Tambak",
				Body:  n.Keterangan,
			}
			u.fcmNotifRepo.PushNotification(deviceIDs, msg)
		}
	}

	return err
}

func (u *tambak) GetAllInfo() ([]models.Info, error) {
	allInfo, err := u.tambakRepo.GetAllInfo()

	return allInfo, err
}

func (u *tambak) CreateInfo(i models.Info) error {
	err := u.tambakRepo.CreateInfo(i)

	return err
}

func (u *tambak) UpdateInfo(i models.Info) error {
	err := u.tambakRepo.UpdateInfo(i)

	return err
}

func (u *tambak) DeleteInfo(id int64) error {
	err := u.tambakRepo.DeleteInfo(id)

	return err
}

func (u *tambak) GetAllPanduan() ([]models.Panduan, error) {
	panduan, err := u.tambakRepo.GetAllPanduan()

	return panduan, err
}

func (u *tambak) CreatePanduan(p models.Panduan) error {
	err := u.tambakRepo.CreatePanduan(p)

	return err
}

func (u *tambak) UpdatePanduan(p models.Panduan) error {
	err := u.tambakRepo.UpdatePanduan(p)

	return err
}

func (u *tambak) DeletePanduan(id int64) error {
	err := u.tambakRepo.DeletePanduan(id)

	return err
}

func (u *tambak) GetMonitorTambak(tambakID int64, tanggal string) ([]models.MonitorTambak, error) {
	m, err := u.tambakRepo.GetMonitorTambak(tambakID, tanggal)

	return m, err
}

func (u *tambak) UpdateJadwal(tambakID int64, val string, _type string) error {
	err := u.tambakRepo.UpdateJadwal(tambakID, val, _type)
	return err
}

func execute(tambakID int64) {
	tambakIDStr := strconv.FormatInt(tambakID, 10)
	log.Println("./script/script.sh " + tambakIDStr)
	cmd := exec.Command("./script/script.sh", tambakIDStr, "&")
	err := cmd.Run()
	fmt.Println("Finished:", err)
	if err != nil {
		fmt.Printf("%s", err)
	}
}
