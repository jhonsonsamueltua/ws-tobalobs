package module

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"

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
	// if err == nil {
	// remote raspberry
	// execute(tambakID)

	//store notif guideline to db
	// loc, _ := time.LoadLocation("Asia/Jakarta")
	// now := time.Now().In(loc)

	// var pemindahanInduk, pemisahanInduk, pemberianPakanSayur, pemberianPakanKeong, panenBenih, panenKonsumsi time.Time

	// if t.JenisBudidaya == "pembenihan" {
	// 	pemindahanInduk = now.AddDate(0, 0, 14)

	// } else {

	// }
	// log.Println("pemindahanInduk : ", pemindahanInduk.Format("2006-01-02 07:00:00"))
	// }

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
			u.tambakRepo.UpdateNotifikasiKondisiTambak(notifID)
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

func (u *tambak) GetAllPanduan() ([]models.Panduan, error) {
	panduan, err := u.tambakRepo.GetAllPanduan()

	return panduan, err
}

func (u *tambak) GetMonitorTambak(tambakID int64, tanggal string) ([]models.MonitorTambak, error) {
	m, err := u.tambakRepo.GetMonitorTambak(tambakID, tanggal)

	return m, err
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
