package mysql

import (
	"log"

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
		err := rows.Scan(&tambak.TambakID, &tambak.NamaTambak, &tambak.Status)
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
	statement, err := r.DB.Prepare(queryInsertTambak)
	if err != nil {
		log.Println("[Repository][CreateTambak][Prepare] Error : ", err)
		return 0, err
	}
	defer statement.Close()

	res, err := statement.Exec(t.UserID, t.NamaTambak, t.Panjang, t.Lebar, t.JenisBudidaya, t.TanggalMulaiBudidaya, t.UsiaLobster, t.JumlahLobster, t.JumlahLobsterJantan, t.JumlahLobsterBetina, t.Status)
	if err != nil {
		log.Println("[Repository][CreateTambak][Execute] Error : ", err)
		return 0, err
	}
	tambakId, _ := res.LastInsertId()
	return tambakId, err
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

func (r *tambak) UpdateNotifikasiKondisiTambak(notifID int64) {
	statement, err := r.DB.Prepare(queryUpdateNotifikasiKondisiTambak)
	if err != nil {
		log.Println("[Repository][UpdateNotifikasiKondisiTambak][Prepare] Error : ", err)
	}
	defer statement.Close()

	_, err = statement.Exec(notifID)
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
