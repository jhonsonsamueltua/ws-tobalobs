package jwt

import (
	"github.com/dgrijalva/jwt-go"

	"github.com/ws-tobalobs/pkg/models"
)

type Usecase interface {
	GenerateJWT(conf *models.Config, userID int64) (string, error)
	ExtractClaims(tokenStr string) (jwt.MapClaims, error)
}
