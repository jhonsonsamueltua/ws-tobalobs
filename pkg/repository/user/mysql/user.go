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
