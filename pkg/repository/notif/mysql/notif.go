package mysql

import (
	"log"

	"github.com/ws-tobalobs/pkg/models"
)

func (r *notif) GetAllNotif(userID int64, tambakID int64, typeNotif string) ([]models.MessagePushNotif, error) {
	allNotif := []models.MessagePushNotif{}
	var query string

	if typeNotif == "unread-per-tambak" {
		query = queryGetAllNotifUnreadPerTambak
	} else if typeNotif == "all-per-tambak" {
		query = queryGetAllNotifPerTambak
	} else if typeNotif == "all-tambak" {
		query = queryGetAllNotif
	}

	statement, err := r.DB.Prepare(query)
	if err != nil {
		log.Println("[Repository][GetAllNotif][Prepare] Error : ", err)
		return allNotif, err
	}

	ID := tambakID
	if typeNotif == "all-tambak" {
		ID = userID
	}
	rows, err := statement.Query(ID)
	if err != nil {
		log.Println("Repository error : ", err)
		return allNotif, err
	}

	for rows.Next() {
		notif := models.MessagePushNotif{}
		err := rows.Scan(&notif.ID, &notif.Title, &notif.Body, &notif.StatusNotifikasi, &notif.TipeNotifikasi)
		if err != nil {
			log.Println(err)
		}
		allNotif = append(allNotif, notif)
	}

	return allNotif, nil
}

func (r *notif) GetDetailNotif(notifID int64) (models.Notifikasi, error) {
	notif := models.Notifikasi{}
	statement, err := r.DB.Prepare(queryGetDetailNotif)
	if err != nil {
		log.Println("[Repository][GetDetailNotif][Prepare] Error : ", err)
		return notif, err
	}
	rows, err := statement.Query(notifID)
	if err != nil {
		log.Println("Repository error : ", err)
		return notif, err
	}

	for rows.Next() {
		err := rows.Scan(&notif.NotifikasiID, &notif.NamaTambak, &notif.TipeNotifikasi, &notif.Keterangan, &notif.WaktuTanggal, &notif.StatusNotifikasi, &notif.AksiPenyimpangan, &notif.KondisiPenyimpangan, &notif.AksiGuideline, &notif.KondisiGuideline)
		if err != nil {
			log.Println("[Repository][GetDetailNotif][Scan] Error : ", err)
			return notif, err
		}
	}

	//update status notifikasi if status is unread
	if notif.StatusNotifikasi == "unread" {
		r.UpdateStatusNotifikasi(notifID)
	}

	return notif, nil
}

func (r *notif) UpdateStatusNotifikasi(notifID int64) {
	statement, err := r.DB.Prepare(queryUpdateStatusNotifikasi)
	if err != nil {
		log.Println("[Repository][UpdateStatusNotifikasi][Prepare] Error : ", err)
	}
	defer statement.Close()

	_, err = statement.Exec(notifID)
	if err != nil {
		log.Println("[Repository][UpdateStatusNotifikasi][Execute] Error : ", err)
	}
}

func (r *notif) GetTotalNofikasiUnread(userID int64) int {
	statement, err := r.DB.Prepare(queryGetTotalNotifikasiUnread)
	if err != nil {
		log.Println("[Repository][GetTotalNofikasiUnread][Prepare] Error : ", err)
	}
	defer statement.Close()

	rows, err := statement.Query(userID)
	if err != nil {
		log.Println("Repository error : ", err)
		return 0
	}

	var totalNotif int

	for rows.Next() {
		err := rows.Scan(&totalNotif)
		if err != nil {
			log.Println("[Repository][GetTotalNofikasiUnread][Scan] Error : ", err)
			return 0
		}
	}
	return totalNotif
}
