package module

import (
	c "github.com/robfig/cron"

	notifRepo "github.com/ws-tobalobs/pkg/repository/notif"
	tambakRepo "github.com/ws-tobalobs/pkg/repository/tambak"
	cronUsecase "github.com/ws-tobalobs/pkg/usecase/cron"
)

type cron struct {
	cr             *c.Cron
	tambakRepo     tambakRepo.Repository
	fcmNotifRepo   notifRepo.RepositoryFCM
	redisNotifRepo notifRepo.RepositoryRedis
	mysqlNotifRepo notifRepo.RepositoryMysql
}

func InitCronUsecase(r tambakRepo.Repository, f notifRepo.RepositoryFCM, red notifRepo.RepositoryRedis, m notifRepo.RepositoryMysql, cr *c.Cron) cronUsecase.Usecase {

	return &cron{
		cr:             cr,
		tambakRepo:     r,
		fcmNotifRepo:   f,
		redisNotifRepo: red,
		mysqlNotifRepo: m,
	}
}
