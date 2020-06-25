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
	userRepoSms   userRepo.RepositorySMS
	jwtUsecase    jwtUsecase.Usecase
	conf          *models.Config
}

func InitUserUsecase(r userRepo.Repository, jwt jwtUsecase.Usecase, c *models.Config, rRedis userRepo.RepositoryRedis, rSms userRepo.RepositorySMS) userUsecase.Usecase {
	return &user{
		userRepo:      r,
		userRepoRedis: rRedis,
		userRepoSms:   rSms,
		jwtUsecase:    jwt,
		conf:          c,
	}
}
