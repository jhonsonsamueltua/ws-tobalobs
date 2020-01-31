package mysql

import (
	"database/sql"

	tambakRepo "github.com/ws-tobalobs/pkg/repository/tambak"
)

type tambak struct {
	DB *sql.DB
}

func InitTambakRepo(db *sql.DB) tambakRepo.Repository {
	return &tambak{
		DB: db,
	}
}
