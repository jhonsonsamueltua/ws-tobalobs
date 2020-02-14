package module

import (
	"github.com/ws-tobalobs/pkg/models"
	// tambakRepo "github.com/ws-tobalobs/pkg/repository/tambak"
	jwtUsecase "github.com/ws-tobalobs/pkg/usecase/jwt"
)

type token struct {
	conf *models.Config
}

func InitJWT(c *models.Config) jwtUsecase.Usecase {
	return &token{
		conf: c,
	}
}
