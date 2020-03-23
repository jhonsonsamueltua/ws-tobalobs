package main

import (
	"database/sql"
	"log"

	"github.com/appleboy/go-fcm"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"

	"github.com/ws-tobalobs/middleware"
	"github.com/ws-tobalobs/pkg/common/config"
	conn "github.com/ws-tobalobs/pkg/common/connection"
	f "github.com/ws-tobalobs/pkg/common/fcm"
	// cron "github.com/ws-tobalobs/pkg/common/cron"
	tambakDeliver "github.com/ws-tobalobs/pkg/delivery/tambak/http"
	userDeliver "github.com/ws-tobalobs/pkg/delivery/user/http"
	"github.com/ws-tobalobs/pkg/models"
	tambakFCMRepo "github.com/ws-tobalobs/pkg/repository/tambak/fcm"
	tambakRepo "github.com/ws-tobalobs/pkg/repository/tambak/mysql"
	userRepo "github.com/ws-tobalobs/pkg/repository/user/mysql"
	userRepoRedis "github.com/ws-tobalobs/pkg/repository/user/redis"
	jwtUseCase "github.com/ws-tobalobs/pkg/usecase/jwt/module"
	tambakUseCase "github.com/ws-tobalobs/pkg/usecase/tambak/module"
	userUseCase "github.com/ws-tobalobs/pkg/usecase/user/module"
)

var Conf *models.Config

func main() {
	Conf = config.InitConfig()
	//DB
	db := conn.InitDB(Conf.Db.Conn)
	defer db.Close()
	//redis
	redis := conn.InitRedis()
	//fcm
	serverKey := "AAAA2Na2Q0M:APA91bHDEdJnN_CjOzLvpScB1DeIYv5SbGldN8-FbvsBLOkTDeyTtfwEiq8uAMvKciQPz41DfHcv2SXJLcIbs13tENnMMJmDAFyjj6vSplVdDwsbXbN4OItWAJh6B2xSD6krelEJt7tV"
	fcm := f.InitFCM(serverKey)

	//http
	e := echo.New()
	middL := middleware.InitMiddleware(redis)
	e.Use(middL.CORS)
	e.Use(middL.JwtAuthentication)

	//module
	tambak(e, db, fcm)
	user(e, db, Conf, redis)

	//Cron
	// cron.InitCron()

	log.Fatal(e.Start(":8000"))
}

func tambak(e *echo.Echo, db *sql.DB, fcm *fcm.Client) {
	tambakRepo := tambakRepo.InitTambakRepo(db)
	tambakFCMRepo := tambakFCMRepo.InitTambakFCMRepo(fcm)
	tambakUsecase := tambakUseCase.InitTambakUsecase(tambakRepo, tambakFCMRepo)
	tambakDeliver.InitTambakHandler(e, tambakUsecase)
}

func user(e *echo.Echo, db *sql.DB, conf *models.Config, redis *redis.Client) {
	userRepo := userRepo.InitUserRepo(db)
	userRepoRedis := userRepoRedis.InitUserRepoRedis(redis)
	jwtUsecase := jwtUseCase.InitJWT(conf)
	userUsecase := userUseCase.InitUserUsecase(userRepo, jwtUsecase, conf, userRepoRedis)
	userDeliver.InitUserHandler(e, userUsecase)
}
