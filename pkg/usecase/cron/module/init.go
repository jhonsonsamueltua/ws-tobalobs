package module

import (
	notifRepo "github.com/ws-tobalobs/pkg/repository/notif"
	tambakRepo "github.com/ws-tobalobs/pkg/repository/tambak"
	userRepo "github.com/ws-tobalobs/pkg/repository/user"
	cronUsecase "github.com/ws-tobalobs/pkg/usecase/cron"
)

type cron struct {
	userRepo       userRepo.Repository
	tambakRepo     tambakRepo.Repository
	fcmNotifRepo   notifRepo.RepositoryFCM
	redisNotifRepo notifRepo.RepositoryRedis
	mysqlNotifRepo notifRepo.RepositoryMysql
}

func InitCronUsecase(r tambakRepo.Repository, f notifRepo.RepositoryFCM, red notifRepo.RepositoryRedis, m notifRepo.RepositoryMysql, u userRepo.Repository) cronUsecase.Usecase {

	return &cron{
		userRepo:       u,
		tambakRepo:     r,
		fcmNotifRepo:   f,
		redisNotifRepo: red,
		mysqlNotifRepo: m,
	}
}
