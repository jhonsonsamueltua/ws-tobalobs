package mysql

import (
	"log"

	"github.com/ws-tobalobs/pkg/models"
)

func (r *tambak) GetAllTambak() ([]models.Tambak, error) {
	allTambak := []models.Tambak{}

	rows, err := r.DB.Query(queryGetAllTambak)
	if err != nil {
		log.Println("Repository error : ", err)
		return allTambak, err
	}

	for rows.Next() {
		tambak := models.Tambak{}
		err := rows.Scan(&tambak.ID, &tambak.Name, &tambak.Location, &tambak.Description)
		if err != nil {
			log.Println(err)
		}
		allTambak = append(allTambak, tambak)
	}

	return allTambak, nil
}

func (r *tambak) PostMonitorTambak(m models.MonitorTambak) error {
	statement, err := r.DB.Prepare(queryInsertMonitoringTambak)
	if err != nil {
		log.Println("[Repository][PostMonitorTambak][Prepare] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(m.TambakId, m.PH, m.DO, m.Suhu, m.WaktuTanggal, m.Keterangan)
	if err != nil {
		log.Println("[Repository][PostMonitorTambak][Execute] Error : ", err)
		return err
	}
	return err
}
