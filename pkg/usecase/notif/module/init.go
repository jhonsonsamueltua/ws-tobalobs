package module

import (
	notifRepo "github.com/ws-tobalobs/pkg/repository/notif"
	tambakRepo "github.com/ws-tobalobs/pkg/repository/tambak"
	notifUsecase "github.com/ws-tobalobs/pkg/usecase/notif"
)

type notif struct {
	tambakRepo     tambakRepo.Repository
	fcmNotifRepo   notifRepo.RepositoryFCM
	redisNotifRepo notifRepo.RepositoryRedis
	mysqlNotifRepo notifRepo.RepositoryMysql
}

func InitNotifUsecase(t tambakRepo.Repository, f notifRepo.RepositoryFCM, red notifRepo.RepositoryRedis, m notifRepo.RepositoryMysql) notifUsecase.Usecase {
	return &notif{
		tambakRepo:     t,
		fcmNotifRepo:   f,
		redisNotifRepo: red,
		mysqlNotifRepo: m,
	}
}
