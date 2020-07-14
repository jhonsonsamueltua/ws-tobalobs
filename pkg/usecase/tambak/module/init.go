package module

import (
	notifRepo "github.com/ws-tobalobs/pkg/repository/notif"
	tambakRepo "github.com/ws-tobalobs/pkg/repository/tambak"
	userRepo "github.com/ws-tobalobs/pkg/repository/user"
	tambakUsecase "github.com/ws-tobalobs/pkg/usecase/tambak"
)

type tambak struct {
	userRepo       userRepo.Repository
	tambakRepo     tambakRepo.Repository
	fcmNotifRepo   notifRepo.RepositoryFCM
	redisNotifRepo notifRepo.RepositoryRedis
	mysqlNotifRepo notifRepo.RepositoryMysql
}

func InitTambakUsecase(r tambakRepo.Repository, f notifRepo.RepositoryFCM, red notifRepo.RepositoryRedis, m notifRepo.RepositoryMysql, u userRepo.Repository) tambakUsecase.Usecase {
	return &tambak{
		userRepo:       u,
		tambakRepo:     r,
		fcmNotifRepo:   f,
		redisNotifRepo: red,
		mysqlNotifRepo: m,
	}
}
