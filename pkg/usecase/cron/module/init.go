package module

import (
	notifRepo "github.com/ws-tobalobs/pkg/repository/notif"
	tambakRepo "github.com/ws-tobalobs/pkg/repository/tambak"
	cronUsecase "github.com/ws-tobalobs/pkg/usecase/cron"
)

type cron struct {
	tambakRepo     tambakRepo.Repository
	fcmNotifRepo   notifRepo.RepositoryFCM
	redisNotifRepo notifRepo.RepositoryRedis
	mysqlNotifRepo notifRepo.RepositoryMysql
}

func InitCronUsecase(r tambakRepo.Repository, f notifRepo.RepositoryFCM, red notifRepo.RepositoryRedis, m notifRepo.RepositoryMysql) cronUsecase.Usecase {

	return &cron{
		tambakRepo:     r,
		fcmNotifRepo:   f,
		redisNotifRepo: red,
		mysqlNotifRepo: m,
	}
}
