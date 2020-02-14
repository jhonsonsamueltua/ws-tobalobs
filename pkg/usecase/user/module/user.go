package module

import (
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/ws-tobalobs/pkg/models"
)

func (u *user) Register(m models.User) (string, error) {
	token := ""
	users, _ := u.userRepo.GetUser(m.Username)
	if (models.User{}) == users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
		if len(hashedPassword) != 0 || err == nil {
			m.Password = string(hashedPassword[:])
			userID, err := u.userRepo.Register(m)
			if err != nil {
				return "", err
			}

			token, err = u.jwtUsecase.GenerateJWT(u.conf, userID)
			if err != nil {
				return "", err
			}
			return token, err
		} else {
			log.Println("[Usecase][User][Register][HashPassword] Error : ", err)
			return "", errors.New("Error Hash Password")
		}
	} else {
		return "", errors.New("Username already exist")
	}
}

func (u *user) Login(username string, password string) (string, error) {
	var err error
	token := ""
	users, _ := u.userRepo.GetUser(username)
	passwordTes := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
	if users.UserID != 0 && passwordTes == nil {
		//login success
		token, err = u.jwtUsecase.GenerateJWT(u.conf, users.UserID)
		if err != nil {
			return "", errors.New("Error Create Token")
		}
	} else {
		return "", errors.New("Username or Password is wrong")
	}

	return token, err
}

func (u *user) Logout(token string) error {
	jwtClaims, err := u.jwtUsecase.ExtractClaims(token)
	if err != nil {
		return err
	}
	exp := jwtClaims["exp"]

	err = u.userRepoRedis.Logout(token, getTokenRemainingValidity(exp))

	return err
}

func getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + 60)
		}
	}
	return 60
}
