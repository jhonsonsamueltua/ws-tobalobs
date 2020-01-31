package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"

	"github.com/ws-tobalobs/middleware"
	"github.com/ws-tobalobs/pkg/common/config"
	conn "github.com/ws-tobalobs/pkg/common/connection"
	tambakDeliver "github.com/ws-tobalobs/pkg/delivery/tambak/http"
	"github.com/ws-tobalobs/pkg/models"
	tambakRepo "github.com/ws-tobalobs/pkg/repository/tambak/mysql"
	tambakUseCase "github.com/ws-tobalobs/pkg/usecase/tambak/module"
)

var Conf *models.Config

func main() {
	Conf = config.InitConfig()
	//http
	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	//DB
	db := conn.InitDB(Conf.Db.Conn)
	// db, err := sql.Open("mysql", Conf.Db.Conn)
	// if err != nil {
	// 	panic(err)
	// }
	defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Successfully connected!")

	//Tambak
	tambak(e, db)

	log.Fatal(e.Start(":8000"))
}

func tambak(e *echo.Echo, db *sql.DB) {
	tambakRepo := tambakRepo.InitTambakRepo(db)
	tambakUsecase := tambakUseCase.InitTambakUsecase(tambakRepo)
	tambakDeliver.InitTambakHandler(e, tambakUsecase)
}
