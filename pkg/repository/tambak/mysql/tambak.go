package mysql

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/sfreiberg/simplessh"

	"github.com/ws-tobalobs/pkg/models"
)

func (r *tambak) GetAllTambak(userID int64) ([]models.Tambak, error) {
	allTambak := []models.Tambak{}
	statement, err := r.DB.Prepare(queryGetAllTambak)
	if err != nil {
		log.Println("[Repository][GetAllTambak][Prepare] Error : ", err)
		return allTambak, err
	}
	rows, err := statement.Query(userID)
	if err != nil {
		log.Println("Repository error : ", err)
		return allTambak, err
	}

	for rows.Next() {
		tambak := models.Tambak{}
		err := rows.Scan(&tambak.TambakID, &tambak.NamaTambak, &tambak.Status, &tambak.PakanPagi, &tambak.PakanSore, &tambak.GantiAir)
		if err != nil {
			log.Println(err)
		}
		allTambak = append(allTambak, tambak)
	}

	return allTambak, nil
}

func (r *tambak) GetTambakByID(tambakID int64, userID int64) (models.Tambak, error) {
	tambak := models.Tambak{}
	statement, err := r.DB.Prepare(queryGetTambakByID)
	if err != nil {
		log.Println("[Repository][GetTambakByID][Prepare] Error : ", err)
		return tambak, err
	}
	rows, err := statement.Query(userID, tambakID)
	if err != nil {
		log.Println("Repository error : ", err)
		return tambak, err
	}

	for rows.Next() {
		err := rows.Scan(&tambak.TambakID, &tambak.NamaTambak, &tambak.Panjang, &tambak.Lebar, &tambak.JenisBudidaya, &tambak.TanggalMulaiBudidaya, &tambak.UsiaLobster, &tambak.JumlahLobster, &tambak.JumlahLobsterJantan, &tambak.JumlahLobsterBetina, &tambak.Status)
		if err != nil {
			log.Println("[Repository][GetTambakByID][Scan] Error : ", err)
			return tambak, err
		}
	}

	return tambak, nil
}

func (r *tambak) GetUserIDByTambak(tambakID int64) int64 {
	var userID int64
	statement, err := r.DB.Prepare(queryGetUserIDByTambak)
	if err != nil {
		log.Println("[Repository][GetUserIDByTambak][Prepare] Error : ", err)
		return userID
	}
	rows, err := statement.Query(tambakID)
	if err != nil {
		log.Println("Repository error : ", err)
		return userID
	}

	for rows.Next() {
		err := rows.Scan(&userID)
		if err != nil {
			log.Println("[Repository][GetUserIDByTambak][Scan] Error : ", err)
			return userID
		}
	}

	return userID
}

func (r *tambak) GetLastMonitorTambak(tambakID int64) (models.MonitorTambak, error) {
	monitor := models.MonitorTambak{}
	statement, err := r.DB.Prepare(queryGetLastMonitorTambak)
	if err != nil {
		log.Println("[Repository][GetLastMonitorTambak][Prepare] Error : ", err)
		return monitor, err
	}
	rows, err := statement.Query(tambakID)
	if err != nil {
		log.Println("Repository error : ", err)
		return monitor, err
	}

	for rows.Next() {
		err := rows.Scan(&monitor.TambakId, &monitor.NamaTambak, &monitor.PH, &monitor.DO, &monitor.Suhu, &monitor.WaktuTanggal, &monitor.Keterangan)
		if err != nil {
			log.Println("[Repository][GetLastMonitorTambak][Scan] Error : ", err)
			return monitor, err
		}
	}

	return monitor, nil
}

func (r *tambak) CreateTambak(t models.Tambak) (int64, error) {
	tambakId := int64(0)
	tx, err := r.DB.Begin()
	if err != nil {
		return tambakId, err
	}

	stmt, err := tx.Prepare(queryInsertTambak)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(t.UserID, t.NamaTambak, t.Panjang, t.Lebar, t.JenisBudidaya, t.TanggalMulaiBudidaya, t.UsiaLobster, t.JumlahLobster, t.JumlahLobsterJantan, t.JumlahLobsterBetina, t.Status, t.PakanPagi, t.PakanSore, t.GantiAir)
	if err != nil {
		log.Println("[Repository][CreateTambak][Execute] Error : ", err)
		return 0, err
	}

	tambakId, _ = res.LastInsertId()
	err = r.execute(tambakId)
	if err != nil {
		log.Println("Error remote raspberry")
		tx.Rollback()
		str := fmt.Sprintf("Gagal menghubungkan ke perangkat IOT dengan error : %s", err.Error())
		er := errors.New(str)
		return 0, er
	}
	tx.Commit()

	return tambakId, err
}

func (r *tambak) execute(tambakID int64) error {
	tambakIDStr := strconv.FormatInt(tambakID, 10)
	var client *simplessh.Client
	var err error

	// get tunnel
	tunnel := r.GetTunnel(int64(1))

	host := tunnel.IP
	port := tunnel.Port
	hostClient := fmt.Sprintf("%s:%s", host, port)
	if client, err = simplessh.ConnectWithPassword(hostClient, "pi", "tobalobs2020"); err != nil {
		// log.Println(err)
		return err
	}

	// Now run the commands on the remote machine:
	comm := fmt.Sprintf("sed -i '$ a python sketchbook/tobalobs/monitor-tambak.py %s &' sketchbook/tobalobs/script.sh", tambakIDStr)
	if _, err := client.Exec(comm); err != nil {
		log.Println(err)
	}
	defer client.Close()

	cmd := exec.Command("./var/www/go/src/github.com/ws-tobalobs/script/script.sh", host, port, tambakIDStr, "&")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("%s", err)
	}

	return err
}

func (r *tambak) UpdateTambak(m models.Tambak) error {
	statement, err := r.DB.Prepare(QueryUpdateTambak)
	if err != nil {
		log.Println("[Repository][UpdateTambak][Prepare] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(m.NamaTambak, m.Panjang, m.Lebar, m.JenisBudidaya, m.UsiaLobster, m.JumlahLobster, m.JumlahLobsterJantan, m.JumlahLobsterBetina, m.TambakID)
	if err != nil {
		log.Println("[Repository][UpdateTambak][Execute] Error : ", err)
		return err
	}
	return err
}

func (r *tambak) PostMonitorTambak(m models.MonitorTambak) (int64, error) {
	statement, err := r.DB.Prepare(queryInsertMonitoringTambak)
	if err != nil {
		log.Println("[Repository][PostMonitorTambak][Prepare] Error : ", err)
		return 0, err
	}
	defer statement.Close()

	res, err := statement.Exec(m.TambakId, m.PH, m.DO, m.Suhu, m.WaktuTanggal, m.Keterangan)
	if err != nil {
		log.Println("[Repository][PostMonitorTambak][Execute] Error : ", err)
		return 0, err
	}
	monitorTambakId, err := res.LastInsertId()
	return monitorTambakId, err
}

func (r *tambak) PostPenyimpanganKondisiTambak(n models.Notifikasi) (int64, error) {
	statement, err := r.DB.Prepare(queryInsertNotifikasiKondisiTambak)
	if err != nil {
		log.Println("[Repository][PostPenyimpanganKondisiTambak][Prepare] Error : ", err)
		return 0, err
	}
	defer statement.Close()

	res, err := statement.Exec(n.TambakID, n.PenyimpanganKondisiTambakID, n.TipeNotifikasi, n.Keterangan, n.StatusNotifikasi, n.WaktuTanggal)
	if err != nil {
		log.Println("[Repository][PostPenyimpanganKondisiTambak][Execute] Error : ", err)
		return 0, err
	}
	notifID, err := res.LastInsertId()
	return notifID, err
}

func (r *tambak) UpdateNotifikasiKondisiTambak(status string, notifID int64) {
	statement, err := r.DB.Prepare(queryUpdateNotifikasiKondisiTambak)
	if err != nil {
		log.Println("[Repository][UpdateNotifikasiKondisiTambak][Prepare] Error : ", err)
	}
	defer statement.Close()

	_, err = statement.Exec(status, notifID)
	if err != nil {
		log.Println("[Repository][UpdateNotifikasiKondisiTambak][Execute] Error : ", err)
	}
}

func (r *tambak) GetAllInfo() ([]models.Info, error) {
	allInfo := []models.Info{}
	statement, err := r.DB.Prepare(queryGetAllInfo)
	if err != nil {
		log.Println("[Repository][GetAllInfo][Prepare] Error : ", err)
		return allInfo, err
	}
	rows, err := statement.Query()
	if err != nil {
		log.Println("Repository error : ", err)
		return allInfo, err
	}

	for rows.Next() {
		info := models.Info{}
		err := rows.Scan(&info.InfoID, &info.Judul, &info.Penjelasan)
		if err != nil {
			log.Println(err)
		}
		allInfo = append(allInfo, info)
	}

	return allInfo, nil
}

func (r *tambak) CreateInfo(i models.Info) error {
	statement, err := r.DB.Prepare(queryInsertInfo)
	if err != nil {
		log.Println("[Repository][CreateInfo][Prepare] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(i.Judul, i.Penjelasan)
	if err != nil {
		log.Println("[Repository][CreateInfo][Execute] Error : ", err)
	}
	return err
}

func (r *tambak) UpdateInfo(m models.Info) error {
	statement, err := r.DB.Prepare(QueryUpdateInfo)
	if err != nil {
		log.Println("[Repository][UpdateInfo][Prepare] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(m.Judul, m.Penjelasan, m.InfoID)
	if err != nil {
		log.Println("[Repository][UpdateInfo][Execute] Error : ", err)
		return err
	}
	return err
}

func (r *tambak) DeleteInfo(id int64) error {
	statement, err := r.DB.Prepare(queryDeleteInfo)
	if err != nil {
		log.Println("[Repository][DeleteInfo][Prepare] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		log.Println("[Repository][DeleteInfo][Execute] Error : ", err)
		return err
	}
	return err
}

func (r *tambak) GetAllPanduan() ([]models.Panduan, error) {
	panduan := []models.Panduan{}
	statement, err := r.DB.Prepare(queryGetAllPanduan)
	if err != nil {
		log.Println("[Repository][GetAllPanduan][Prepare] Error : ", err)
		return panduan, err
	}
	rows, err := statement.Query()
	if err != nil {
		log.Println("Repository error : ", err)
		return panduan, err
	}

	for rows.Next() {
		p := models.Panduan{}
		err := rows.Scan(&p.PanduanAplikasiID, &p.Judul, &p.Penjelasan)
		if err != nil {
			log.Println(err)
		}
		panduan = append(panduan, p)
	}

	return panduan, nil
}

func (r *tambak) CreatePanduan(p models.Panduan) error {
	statement, err := r.DB.Prepare(queryInsertPanduan)
	if err != nil {
		log.Println("[Repository][CreatePanduan][Prepare] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(p.Judul, p.Penjelasan)
	if err != nil {
		log.Println("[Repository][CreatePanduan][Execute] Error : ", err)
	}
	return err
}

func (r *tambak) UpdatePanduan(m models.Panduan) error {
	statement, err := r.DB.Prepare(QueryUpdatePanduan)
	if err != nil {
		log.Println("[Repository][UpdatePanduan][Prepare] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(m.Judul, m.Penjelasan, m.PanduanAplikasiID)
	if err != nil {
		log.Println("[Repository][UpdatePanduan][Execute] Error : ", err)
		return err
	}
	return err
}

func (r *tambak) DeletePanduan(id int64) error {
	statement, err := r.DB.Prepare(queryDeletePanduan)
	if err != nil {
		log.Println("[Repository][DeletePanduan][Prepare] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		log.Println("[Repository][UpdatePanduan][Execute] Error : ", err)
		return err
	}
	return err
}

func (r *tambak) GetMonitorTambak(tambakID int64, tanggal string) ([]models.MonitorTambak, error) {
	monitor := []models.MonitorTambak{}
	statement, err := r.DB.Prepare(queryGetMonitorTambak)
	if err != nil {
		log.Println("[Repository][GetMonitorTambak][Prepare] Error : ", err)
		return monitor, err
	}
	rows, err := statement.Query(tambakID, tanggal)
	if err != nil {
		log.Println("Repository error : ", err)
		return monitor, err
	}

	for rows.Next() {
		m := models.MonitorTambak{}
		err := rows.Scan(&m.MonitorTambakId, &m.PH, &m.DO, &m.Suhu, &m.WaktuTanggal, &m.Keterangan)
		if err != nil {
			log.Println(err)
		}
		monitor = append(monitor, m)
	}

	return monitor, nil
}

func (r *tambak) GetAllTambakID() ([]int64, []int64, []string, error) {
	userID := []int64{}
	tambakID := []int64{}
	namaTambak := []string{}

	statement, err := r.DB.Prepare(queryGetAllTambakID)
	if err != nil {
		log.Println("[Repository][GetAllTambakID][Prepare] Error : ", err)
		return userID, tambakID, namaTambak, err
	}
	rows, err := statement.Query()
	if err != nil {
		log.Println("Repository error : ", err)
		return userID, tambakID, namaTambak, err
	}

	for rows.Next() {
		var uID int64
		var tID int64
		var nama string
		err := rows.Scan(&tID, &uID, &nama)
		if err != nil {
			log.Println(err)
		}
		userID = append(userID, uID)
		tambakID = append(tambakID, tID)
		namaTambak = append(namaTambak, nama)
	}

	return userID, tambakID, namaTambak, err
}

func (r *tambak) UpdateJadwal(tambakID int64, val string, _type string) error {
	query := queryUpdateJadwalPakanPagi
	if _type == "pakan_sore" {
		query = queryUpdateJadwalPakanSore
	} else if _type == "ganti_air" {
		query = queryUpdateJadwalGantiAir
	}

	statement, err := r.DB.Prepare(query)
	if err != nil {
		log.Println("[Repository][UpdateJadwal][Prepare] Error : ", err)
	}
	defer statement.Close()

	_, err = statement.Exec(val, tambakID)
	if err != nil {
		log.Println("[Repository][UpdateJadwal][Execute] Error : ", err)
	}

	return err
}

func (r *tambak) GetAllGuideline() ([]models.Guideline, error) {
	res := []models.Guideline{}
	statement, err := r.DB.Prepare(queryGetAllGuideline)
	if err != nil {
		log.Println("[Repository][GetAllGuideline][Prepare] Error : ", err)
		return res, err
	}
	rows, err := statement.Query()
	if err != nil {
		log.Println("Repository error : ", err)
		return res, err
	}

	for rows.Next() {
		r := models.Guideline{}
		err := rows.Scan(&r.GuidelineID, &r.AksiGuideline, &r.Notifikasi, &r.TipeBudidaya, &r.TipeJadwal, &r.Interval, &r.Waktu)
		if err != nil {
			log.Println(err)
		}
		res = append(res, r)
	}

	return res, nil
}

func (r *tambak) CreateGuideline(m models.Guideline) error {
	statement, err := r.DB.Prepare(queryAddGuideline)
	if err != nil {
		log.Println("[Repository][CreateGuideline][Prepare] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(m.AksiGuideline, m.Notifikasi, m.TipeBudidaya, m.TipeJadwal, m.Interval, m.Waktu)
	if err != nil {
		log.Println("[Repository][CreateGuideline][Execute] Error : ", err)
	}
	return err
}

func (r *tambak) UpdateGuideline(m models.Guideline) error {
	statement, err := r.DB.Prepare(queryUpdateGuideline)
	if err != nil {
		log.Println("[Repository][UpdateGuideline][Prepare] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(m.AksiGuideline, m.Notifikasi, m.TipeBudidaya, m.TipeJadwal, m.Interval, m.Waktu, m.GuidelineID)
	if err != nil {
		log.Println("[Repository][UpdateGuideline][Execute] Error : ", err)
		return err
	}
	return err
}

func (r *tambak) GetTunnel(ID int64) models.Tunnel {
	res := models.Tunnel{}
	statement, err := r.DB.Prepare(queryGetTunnel)
	if err != nil {
		log.Println("[Repository][GetTunnel][Prepare] Error : ", err)
		return res
	}
	rows, err := statement.Query(ID)
	if err != nil {
		log.Println("Repository error : ", err)
		return res
	}

	for rows.Next() {
		err := rows.Scan(&res.ID, &res.IP, &res.Port)
		if err != nil {
			log.Println("[Repository][GetTunnel][Scan] Error : ", err)
			return res
		}
	}

	return res
}

func (r *tambak) SaveTunnel(m models.Tunnel) {
	statement, err := r.DB.Prepare(queryUpdateTunnel)
	if err != nil {
		log.Println("[Repository][SaveTunnel][Prepare] Error : ", err)
		return
	}
	defer statement.Close()

	_, err = statement.Exec(m.IP, m.Port, m.ID)
	if err != nil {
		log.Println("[Repository][SaveTunnel][Execute] Error : ", err)
	}
	return
}

func (r *tambak) GetKondisiSekarang(ip string) (models.KondisiSekarang, error) {
	res := models.KondisiSekarang{}
	endpoint := fmt.Sprintf("%s/get-monitor", ip)
	resp, err := http.Get(endpoint)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	jsonErr := json.Unmarshal(body, &res)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Printf("Results: %v\n", res)
	return res, err
}
