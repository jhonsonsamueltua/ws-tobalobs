package module

import (
	tambakRepo "github.com/ws-tobalobs/pkg/repository/tambak"
	tambakUsecase "github.com/ws-tobalobs/pkg/usecase/tambak"
)

type tambak struct {
	tambakRepo    tambakRepo.Repository
	tambakFCMRepo tambakRepo.RepositoryFCM
}

func InitTambakUsecase(r tambakRepo.Repository, f tambakRepo.RepositoryFCM) tambakUsecase.Usecase {
	return &tambak{
		tambakRepo:    r,
		tambakFCMRepo: f,
	}
}
