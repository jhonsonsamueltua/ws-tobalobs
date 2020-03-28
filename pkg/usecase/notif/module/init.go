package module

import (
	notifRepo "github.com/ws-tobalobs/pkg/repository/notif"
	notifUsecase "github.com/ws-tobalobs/pkg/usecase/notif"
)

type notif struct {
	notifRepo notifRepo.RepositoryMysql
}

func InitNotifUsecase(r notifRepo.RepositoryMysql) notifUsecase.Usecase {
	return &notif{
		notifRepo: r,
	}
}
