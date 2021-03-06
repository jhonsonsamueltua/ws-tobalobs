package mysql

import (
	"log"

	"github.com/ws-tobalobs/pkg/models"
)

func (r *user) Register(m models.User) (int64, error) {
	statement, err := r.DB.Prepare(queryInsertUser)
	if err != nil {
		log.Println("[Repository][Register][Prepare] Error : ", err)
		return 0, err
	}

	defer statement.Close()

	res, err := statement.Exec(m.Username, m.Password, m.Nama, m.Alamat, m.NoHp, m.TanggalLahir, m.Role)
	if err != nil {
		log.Println("[Repository][Register][Execute] Error : ", err)
		return 0, err
	}
	userID, _ := res.LastInsertId()
	return userID, err
}

func (r *user) GetUser(username string) (models.User, error) {
	var users = models.User{}
	statement, err := r.DB.Prepare(querySelectUser)
	if err != nil {
		log.Println("[Repository][GetUser][Prepare] Error : ", err)
		return users, err
	}

	defer statement.Close()

	err = statement.QueryRow(username).Scan(&users.UserID, &users.Username, &users.Password, &users.Nama, &users.Alamat, &users.NoHp, &users.TanggalLahir, &users.Role)

	return users, err
}

func (r *user) GetByPhoneNumber(hp string) (models.User, error) {
	var users = models.User{}
	statement, err := r.DB.Prepare(querySelectUserByHP)
	if err != nil {
		log.Println("[Repository][GetByPhoneNumber][Prepare] Error : ", err)
		return users, err
	}

	defer statement.Close()

	err = statement.QueryRow(hp).Scan(&users.UserID, &users.Username, &users.Password, &users.Nama, &users.Alamat, &users.NoHp, &users.TanggalLahir, &users.Role)

	return users, err
}

func (r *user) GetDetailUser(userID int64) (models.User, error) {
	var users = models.User{}
	statement, err := r.DB.Prepare(querySelectDetailUser)
	if err != nil {
		log.Println("[Repository][GetDetailUser][Prepare] Error : ", err)
		return users, err
	}

	defer statement.Close()

	err = statement.QueryRow(userID).Scan(&users.UserID, &users.Username, &users.Password, &users.Nama, &users.Alamat, &users.NoHp, &users.TanggalLahir)

	return users, err
}

func (r *user) UpdateUser(m models.User) error {
	statement, err := r.DB.Prepare(QueryUpdateUser)
	if err != nil {
		log.Println("[Repository][UpdateUser][Prepare] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(m.Username, m.Nama, m.Alamat, m.NoHp, m.TanggalLahir, m.UserID)
	if err != nil {
		log.Println("[Repository][UpdateUser][Execute] Error : ", err)
		return err
	}
	return err
}

func (r *user) UpdatePassword(newPass string, userID int64) error {
	statement, err := r.DB.Prepare(QueryUpdatePassword)
	if err != nil {
		log.Println("[Repository][UpdatePassword][Prepare] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(newPass, userID)
	if err != nil {
		log.Println("[Repository][UpdatePassword][Execute] Error : ", err)
		return err
	}
	return err
}

func (r *user) GetKondisiMenyimpang() ([]models.KondisiMenyimpang, error) {
	res := []models.KondisiMenyimpang{}
	statement, err := r.DB.Prepare(queryGetAllKondisi)
	if err != nil {
		log.Println("[Repository][GetKondisiMenyimpang][Prepare] Error : ", err)
		return res, err
	}
	rows, err := statement.Query()
	if err != nil {
		log.Println("Repository error : ", err)
		return res, err
	}

	for rows.Next() {
		cond := models.KondisiMenyimpang{}
		err := rows.Scan(&cond.ID, &cond.AksiPenyimpangan, &cond.Kondisi, &cond.Tipe, &cond.Nilai)
		if err != nil {
			log.Println(err)
		}
		res = append(res, cond)
	}

	return res, nil
}

func (r *user) UpdateKondisiMenyimpang(m models.KondisiMenyimpang) error {
	statement, err := r.DB.Prepare(QueryUpdateKondisi)
	if err != nil {
		log.Println("[Repository][UpdateKondisiMenyimpang][Prepare] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(m.AksiPenyimpangan, m.Kondisi, m.Tipe, m.Nilai, m.ID)
	if err != nil {
		log.Println("[Repository][UpdateKondisiMenyimpang][Execute] Error : ", err)
		return err
	}
	return err
}

func (r *user) GetDeviceID(userID int64) []string {
	res := []string{}
	statement, err := r.DB.Prepare(queryGetDeviceID)
	if err != nil {
		log.Println("[Repository][GetDeviceID][Prepare] Error : ", err)
		return res
	}
	rows, err := statement.Query(userID)
	if err != nil {
		log.Println("Repository error : ", err)
		return res
	}

	for rows.Next() {
		var deviceId string
		err := rows.Scan(&deviceId)
		if err != nil {
			log.Println(err)
		}
		res = append(res, deviceId)
	}

	return res
}

func (r *user) SaveDeviceID(userID int64, deviceID string) {
	statement, err := r.DB.Prepare(querySaveDeviceID)
	if err != nil {
		log.Println("[Repository][SaveDeviceID][Prepare] Error : ", err)
	}

	defer statement.Close()

	_, err = statement.Exec(userID, deviceID)
	if err != nil {
		log.Println("[Repository][SaveDeviceID][Execute] Error : ", err)
	}
}

func (r *user) DeleteDeviceID(userID int64, deviceID string) {
	statement, err := r.DB.Prepare(queryDeleteDeviceID)
	if err != nil {
		log.Println("[Repository][DeleteDeviceID][Prepare] Error : ", err)
	}

	defer statement.Close()

	_, err = statement.Exec(userID, deviceID)
	if err != nil {
		log.Println("[Repository][DeleteDeviceID][Execute] Error : ", err)
	}
}
