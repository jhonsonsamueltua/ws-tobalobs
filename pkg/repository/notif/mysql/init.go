package mysql

import (
	"database/sql"

	notifRepo "github.com/ws-tobalobs/pkg/repository/notif"
)

type notif struct {
	DB *sql.DB
}

func InitNotifRepo(db *sql.DB) notifRepo.RepositoryMysql {
	return &notif{
		DB: db,
	}
}
