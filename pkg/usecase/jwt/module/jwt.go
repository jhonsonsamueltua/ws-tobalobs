package module

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/ws-tobalobs/pkg/models"
)

func (u *token) GenerateJWT(conf *models.Config, userID int64) (string, error) {

	tk := &models.Token{
		StandardClaims: jwt.StandardClaims{
			Issuer:    u.conf.Token.Issuer,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		UserId: userID,
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(u.conf.Token.Key))
	if err != nil {
		log.Println("[Usecase][JWT][GenerateJWT] Error : ", err)
		return "", err
	}
	return tokenString, err
}

func (u *token) ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	hmacSecretString := "asdfghjkl"
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		log.Println("Error Ekstract Token : ", err)
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, err
	} else {
		log.Printf("Invalid JWT Token")
		return nil, err
	}
}
