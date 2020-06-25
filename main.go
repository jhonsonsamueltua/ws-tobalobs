package main

import (
	"database/sql"
	"log"
	"os"

	"firebase.google.com/go/messaging"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"

	"github.com/ws-tobalobs/middleware"
	"github.com/ws-tobalobs/pkg/common/config"
	conn "github.com/ws-tobalobs/pkg/common/connection"
	f "github.com/ws-tobalobs/pkg/common/fcm"
	// cron "github.com/ws-tobalobs/pkg/common/cron"
	notifDeliver "github.com/ws-tobalobs/pkg/delivery/notif/http"
	tambakDeliver "github.com/ws-tobalobs/pkg/delivery/tambak/http"
	userDeliver "github.com/ws-tobalobs/pkg/delivery/user/http"
	"github.com/ws-tobalobs/pkg/models"
	fcmNotifRepo "github.com/ws-tobalobs/pkg/repository/notif/fcm"
	mysqlNotifRepo "github.com/ws-tobalobs/pkg/repository/notif/mysql"
	redisNotifRepo "github.com/ws-tobalobs/pkg/repository/notif/redis"
	tambakRepo "github.com/ws-tobalobs/pkg/repository/tambak/mysql"
	userRepo "github.com/ws-tobalobs/pkg/repository/user/mysql"
	userRepoRedis "github.com/ws-tobalobs/pkg/repository/user/redis"
	userRepoSms "github.com/ws-tobalobs/pkg/repository/user/sms"
	cronUseCase "github.com/ws-tobalobs/pkg/usecase/cron/module"
	jwtUseCase "github.com/ws-tobalobs/pkg/usecase/jwt/module"
	notifUseCase "github.com/ws-tobalobs/pkg/usecase/notif/module"
	tambakUseCase "github.com/ws-tobalobs/pkg/usecase/tambak/module"
	userUseCase "github.com/ws-tobalobs/pkg/usecase/user/module"
)

var Conf *models.Config

func main() {
	var connDB, connRedis string
	//config
	Conf = config.InitConfig()

	//environtment
	var env string
	var args = os.Args

	if len(args) > 1 {
		env = os.Args[1]
	}

	if env == "prod" {
		connDB = Conf.Database.Prod
		connRedis = Conf.Redis.Prod
	} else {
		connDB = Conf.Database.Devel
		connRedis = Conf.Redis.Devel
	}

	//SSH connect
	conn.ConnectSSH()

	//DB
	db := conn.InitDB(connDB)
	defer db.Close()

	//redis
	redis := conn.InitRedis(connRedis)
	defer redis.Close()

	//fcm
	fcm := f.InitFCM(Conf.Fcm.Key)

	//http
	e := echo.New()
	middL := middleware.InitMiddleware(redis)
	e.Use(middL.CORS)
	e.Use(middL.JwtAuthentication)

	//module
	tambak(e, db, fcm, redis)
	user(e, db, Conf, redis)
	notif(e, db, fcm, redis)
	cron(db, fcm, redis)

	log.Fatal(e.Start(":8000"))
}

func tambak(e *echo.Echo, db *sql.DB, fcm *messaging.Client, redis *redis.Client) {
	tambakRepo := tambakRepo.InitTambakRepo(db)
	fcmNotifRepo := fcmNotifRepo.InitFCMRepo(fcm)
	redisNotifRepo := redisNotifRepo.InitRedisRepo(redis)
	mysqlNotifRepo := mysqlNotifRepo.InitNotifRepo(db)
	tambakUsecase := tambakUseCase.InitTambakUsecase(tambakRepo, fcmNotifRepo, redisNotifRepo, mysqlNotifRepo)
	tambakDeliver.InitTambakHandler(e, tambakUsecase)
}

func notif(e *echo.Echo, db *sql.DB, fcm *messaging.Client, redis *redis.Client) {
	tambakRepo := tambakRepo.InitTambakRepo(db)
	fcmNotifRepo := fcmNotifRepo.InitFCMRepo(fcm)
	redisNotifRepo := redisNotifRepo.InitRedisRepo(redis)
	mysqlNotifRepo := mysqlNotifRepo.InitNotifRepo(db)
	notifUsecase := notifUseCase.InitNotifUsecase(tambakRepo, fcmNotifRepo, redisNotifRepo, mysqlNotifRepo)
	notifDeliver.InitNotifHandler(e, notifUsecase)
}

func user(e *echo.Echo, db *sql.DB, conf *models.Config, redis *redis.Client) {
	userRepo := userRepo.InitUserRepo(db)
	userRepoRedis := userRepoRedis.InitUserRepoRedis(redis)
	userRepoSms := userRepoSms.InitSendSMS()
	jwtUsecase := jwtUseCase.InitJWT(conf)
	userUsecase := userUseCase.InitUserUsecase(userRepo, jwtUsecase, conf, userRepoRedis, userRepoSms)
	userDeliver.InitUserHandler(e, userUsecase)
}

func cron(db *sql.DB, fcm *messaging.Client, redis *redis.Client) {
	tambakRepo := tambakRepo.InitTambakRepo(db)
	fcmNotifRepo := fcmNotifRepo.InitFCMRepo(fcm)
	redisNotifRepo := redisNotifRepo.InitRedisRepo(redis)
	mysqlNotifRepo := mysqlNotifRepo.InitNotifRepo(db)
	cron := cronUseCase.InitCronUsecase(tambakRepo, fcmNotifRepo, redisNotifRepo, mysqlNotifRepo)
	cron.InitCron()
}
