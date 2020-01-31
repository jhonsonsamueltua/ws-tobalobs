package connection

import (
	"database/sql"
	"fmt"

	// _ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB(conn string) *sql.DB {
	db, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}
