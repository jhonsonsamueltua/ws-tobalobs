package module

import (
	tambakRepo "github.com/ws-tobalobs/pkg/repository/tambak"
	tambakUsecase "github.com/ws-tobalobs/pkg/usecase/tambak"
)

type tambak struct {
	tambakRepo tambakRepo.Repository
}

func InitTambakUsecase(p tambakRepo.Repository) tambakUsecase.Usecase {
	return &tambak{
		tambakRepo: p,
	}
}
