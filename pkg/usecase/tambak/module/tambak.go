package module

import (
	"fmt"
	"log"
	// "math"
	"strconv"
	"time"

	"github.com/sfreiberg/simplessh"

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

func (u *tambak) GetAllTambakForAdmin() ([]models.Tambak, error) {
	allTambak, err := u.tambakRepo.GetAllTambakForAdmin()

	return allTambak, err
}

func (u *tambak) GetTambakByID(tambakID int64, userID int64) (models.Tambak, error) {
	tambak, err := u.tambakRepo.GetTambakByID(tambakID, userID)

	format := "2006-01-02"
	dt, _ := time.Parse(format, tambak.TanggalMulaiBudidaya)
	tambak.TanggalMulaiBudidaya = dt.Format("2 Jan 2006")

	return tambak, err
}

func (u *tambak) GetLastMonitorTambak(tambakID int64) (models.MonitorTambak, error) {
	monitor, err := u.tambakRepo.GetLastMonitorTambak(tambakID)

	// monitor.PH = math.Floor(monitor.PH*100) / 100
	// monitor.Suhu = math.Floor(monitor.Suhu*100) / 100
	// monitor.DO = math.Floor(monitor.DO*100) / 100

	format := "2006-01-02 15:04:05"
	dt, _ := time.Parse(format, monitor.WaktuTanggal)
	monitor.WaktuTanggal = dt.Format("2 Jan 2006 - 15:04")

	return monitor, err
}

func (u *tambak) CreateTambak(t models.Tambak) (int64, error) {
	tambakID, err := u.tambakRepo.CreateTambak(t)
	if err != nil {
		return tambakID, err
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	n := models.Notifikasi{
		TambakID:         tambakID,
		TipeNotifikasi:   "notif-guideline",
		StatusNotifikasi: "waiting",
	}

	guide, _ := u.tambakRepo.GetAllGuideline()
	if t.JenisBudidaya == "pembenihan" {
		for _, g := range guide {
			var dt time.Time
			interval, _ := strconv.Atoi(g.Interval)
			if g.TipeBudidaya == "pembenihan" || g.TipeBudidaya == "semua" {
				if g.TipeJadwal == "sekali" {
					dt = now.AddDate(0, 0, interval)
					n.GuidelineID = g.GuidelineID
					n.Keterangan = g.Notifikasi
					n.WaktuTanggal = fmt.Sprintf("%s %s:00", dt.Format("2006-01-02"), g.Waktu)
					_, err := u.mysqlNotifRepo.SaveNotifGuideline(n)
					if err != nil {
						log.Println(err)
					}
				} else if g.TipeJadwal == "berulang" && interval > 5 {
					lamaBudidaya := 8 - t.UsiaLobster
					for i := 1; i <= (lamaBudidaya * 30 / interval); i++ {
						if i == 1 {
							dt = now.AddDate(0, 0, interval)
						} else {
							dt = dt.AddDate(0, 0, interval)
						}
						n.GuidelineID = g.GuidelineID
						n.Keterangan = g.Notifikasi
						n.WaktuTanggal = fmt.Sprintf("%s %s:00", dt.Format("2006-01-02"), g.Waktu)
						_, err := u.mysqlNotifRepo.SaveNotifGuideline(n)
						if err != nil {
							log.Println(err)
						}
					}
				}
			}
		}
	} else if t.JenisBudidaya == "pembesaran" {
		for _, g := range guide {
			var dt time.Time
			interval, _ := strconv.Atoi(g.Interval)
			if g.TipeBudidaya == "pembesaran" || g.TipeBudidaya == "semua" {
				if g.TipeJadwal == "sekali" {
					if interval == 0 { // interval 0 artinya tergantung umur lobster
						lamaBudidaya := 8 - t.UsiaLobster
						dt = now.AddDate(0, 0, lamaBudidaya*30)
						n.GuidelineID = g.GuidelineID
						n.Keterangan = g.Notifikasi
						n.WaktuTanggal = fmt.Sprintf("%s %s:00", dt.Format("2006-01-02"), g.Waktu)
						_, err := u.mysqlNotifRepo.SaveNotifGuideline(n)
						if err != nil {
							log.Println(err)
						}
					} else {
						dt = now.AddDate(0, 0, interval)
						n.GuidelineID = g.GuidelineID
						n.Keterangan = g.Notifikasi
						n.WaktuTanggal = fmt.Sprintf("%s %s:00", dt.Format("2006-01-02"), g.Waktu)
						_, err := u.mysqlNotifRepo.SaveNotifGuideline(n)
						if err != nil {
							log.Println(err)
						}
					}
				} else if g.TipeJadwal == "berulang" {
					lamaBudidaya := 8 - t.UsiaLobster
					if interval > 5 {
						for i := 1; i <= (lamaBudidaya * 30 / interval); i++ {
							if i == 1 {
								dt = now.AddDate(0, 0, interval)
							} else {
								dt = dt.AddDate(0, 0, interval)
							}
							n.GuidelineID = g.GuidelineID
							n.Keterangan = g.Notifikasi
							n.WaktuTanggal = fmt.Sprintf("%s %s:00", dt.Format("2006-01-02"), g.Waktu)
							_, err := u.mysqlNotifRepo.SaveNotifGuideline(n)
							if err != nil {
								log.Println(err)
							}
						}
					}
				}
			}
		}
	}

	return tambakID, err
}

// func (u *tambak) CreateTambak(t models.Tambak) (int64, error) {
// 	tambakID, err := u.tambakRepo.CreateTambak(t)
// 	if err == nil {
// 		// remote raspberry
// 		execute(tambakID)

// 		// store notif guideline to db
// 		loc, _ := time.LoadLocation("Asia/Jakarta")
// 		now := time.Now().In(loc)

// 		var pemindahanInduk, pemisahanInduk, pemberianPakanSayur, pemberianPakanKeong, panenBenih, panenKonsumsi time.Time
// 		n := models.Notifikasi{
// 			TambakID:         tambakID,
// 			TipeNotifikasi:   "notif-guideline",
// 			StatusNotifikasi: "waiting",
// 		}

// 		if t.JenisBudidaya == "pembenihan" {
// 			//pemindahan induk
// 			pemindahanInduk = now.AddDate(0, 0, 14)
// 			n.GuidelineID = 4
// 			n.Keterangan = "Pemindahan induk"
// 			n.WaktuTanggal = pemindahanInduk.Format("2006-01-02 07:00:00")
// 			_, err := u.mysqlNotifRepo.SaveNotifGuideline(n)
// 			if err != nil {
// 				log.Println(err)
// 			}
// 			//pemisahan induk
// 			pemisahanInduk = pemindahanInduk.AddDate(0, 0, 45)
// 			n.GuidelineID = 5
// 			n.Keterangan = "Pemisahan induk dari anakan (telur menetas dan turun anak)"
// 			n.WaktuTanggal = pemisahanInduk.Format("2006-01-02 07:00:00")
// 			_, err = u.mysqlNotifRepo.SaveNotifGuideline(n)
// 			if err != nil {
// 				log.Println(err)
// 			}
// 			//panen benih
// 			panenBenih = pemindahanInduk.AddDate(0, 2, 0)
// 			n.GuidelineID = 8
// 			n.Keterangan = "Panen benih lobster untuk benih"
// 			n.WaktuTanggal = fmt.Sprintf("%s 18:30:00", panenBenih.Format("2006-01-02"))
// 			_, err = u.mysqlNotifRepo.SaveNotifGuideline(n)
// 			if err != nil {
// 				log.Println(err)
// 			}

// 		} else {
// 			//pemberian pakan sayuran
// 			lamaBudidaya := 8 - t.UsiaLobster
// 			for i := 1; i <= (lamaBudidaya * 2); i++ {
// 				if i == 1 {
// 					pemberianPakanSayur = now.AddDate(0, 0, 14)
// 				} else {
// 					pemberianPakanSayur = pemberianPakanSayur.AddDate(0, 0, 14)
// 				}
// 				n.GuidelineID = 6
// 				n.Keterangan = "Pemberian pakan sayuran"
// 				n.WaktuTanggal = pemberianPakanSayur.Format("2006-01-02 07:00:00")
// 				_, err := u.mysqlNotifRepo.SaveNotifGuideline(n)
// 				if err != nil {
// 					log.Println(err)
// 				}
// 			}

// 			//pemberian pakan keong
// 			for i := 1; i <= lamaBudidaya; i++ {
// 				if i == 1 {
// 					pemberianPakanKeong = now.AddDate(0, 1, 0)
// 				} else {
// 					pemberianPakanKeong = pemberianPakanKeong.AddDate(0, 1, 0)
// 				}
// 				n.GuidelineID = 7
// 				n.Keterangan = "Pemberian pakan keong"
// 				n.WaktuTanggal = pemberianPakanKeong.Format("2006-01-02 07:00:00")
// 				_, err := u.mysqlNotifRepo.SaveNotifGuideline(n)
// 				if err != nil {
// 					log.Println(err)
// 				}
// 			}

// 			//panen konsumsi
// 			panenKonsumsi = now.AddDate(0, lamaBudidaya, 0)
// 			n.GuidelineID = 8
// 			n.Keterangan = "Panen lobster pembesaran"
// 			n.WaktuTanggal = fmt.Sprintf("%s 18:30:00", panenKonsumsi.Format("2006-01-02"))
// 			_, err = u.mysqlNotifRepo.SaveNotifGuideline(n)
// 			if err != nil {
// 				log.Println(err)
// 			}

// 		}
// 	}

// 	return tambakID, err
// }

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
		// deviceIDs := u.redisNotifRepo.GetDeviceID(userID)
		deviceIDs := u.userRepo.GetDeviceID(userID)
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
	monitor := []models.MonitorTambak{}
	m, err := u.tambakRepo.GetMonitorTambak(tambakID, tanggal)

	for _, d := range m {
		data := models.MonitorTambak{}
		data = d
		if d.Keterangan == "Kondisi tambak normal" {
			data.Keterangan = "Normal"
		} else {
			data.Keterangan = "Bermasalah"
		}
		monitor = append(monitor, data)
	}

	return monitor, err
}

func (u *tambak) UpdateJadwal(tambakID int64, val string, _type string) error {
	err := u.tambakRepo.UpdateJadwal(tambakID, val, _type)
	return err
}

func execute(tambakID int64) error {
	tambakIDStr := strconv.FormatInt(tambakID, 10)
	// log.Println("./script/script.sh " + tambakIDStr)
	// cmd := exec.Command("./script/script.sh", tambakIDStr, "&")
	// ssh -l pi proxy73.rt3.io -p 35745 ./sketchbook/tobalobs/script-remote.sh $1 &
	// cmd := exec.LookPath
	// cmd := exec.Command("ssh", "pi@0.tcp.ngrok.io -p 18899", "./sketchbook/tobalobs/script-remote.sh", tambakIDStr, "&")
	// err := cmd.Run()
	// fmt.Println("Finished:", err)
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// }
	var client *simplessh.Client
	var err error

	if client, err = simplessh.ConnectWithPassword("2.tcp.ngrok.io:18899", "pi", "raspberry"); err != nil {
		log.Println(err)
	}

	// Now run the commands on the remote machine:
	cmd := fmt.Sprintf("./sketchbook/tobalobs/script-remote.sh %s &", tambakIDStr)
	if _, err := client.Exec(cmd); err != nil {
		log.Println(err)
	}
	defer client.Close()

	return err
}

func (u *tambak) GetAllGuideline() ([]models.Guideline, error) {
	res, err := u.tambakRepo.GetAllGuideline()

	return res, err
}

func (u *tambak) CreateGuideline(m models.Guideline) error {
	err := u.tambakRepo.CreateGuideline(m)

	return err
}

func (u *tambak) UpdateGuideline(m models.Guideline) error {
	err := u.tambakRepo.UpdateGuideline(m)

	return err
}

func (u *tambak) GetTunnel(ID int64) models.Tunnel {
	res := u.tambakRepo.GetTunnel(ID)

	return res
}

func (u *tambak) SaveTunnel(m models.Tunnel) {
	u.tambakRepo.SaveTunnel(m)

	return
}

func (u *tambak) GetKondisiSekarang() (models.KondisiSekarang, error) {
	// get tunnel
	tunnel := u.tambakRepo.GetTunnel(int64(2))
	log.Println(tunnel.IP)
	res, err := u.tambakRepo.GetKondisiSekarang(tunnel.IP)
	return res, err
}
