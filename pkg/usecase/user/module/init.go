package module

import (
	"github.com/ws-tobalobs/pkg/models"
	userRepo "github.com/ws-tobalobs/pkg/repository/user"
	jwtUsecase "github.com/ws-tobalobs/pkg/usecase/jwt"
	userUsecase "github.com/ws-tobalobs/pkg/usecase/user"
)

type user struct {
	userRepo      userRepo.Repository
	userRepoRedis userRepo.RepositoryRedis
	jwtUsecase    jwtUsecase.Usecase
	conf          *models.Config
}

func InitUserUsecase(r userRepo.Repository, jwt jwtUsecase.Usecase, c *models.Config, rRedis userRepo.RepositoryRedis) userUsecase.Usecase {
	return &user{
		userRepo:      r,
		userRepoRedis: rRedis,
		jwtUsecase:    jwt,
		conf:          c,
	}
}
