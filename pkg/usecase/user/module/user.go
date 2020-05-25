package module

import (
	"errors"
	"fmt"
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

func (u *user) Login(username string, password string, deviceID string) (string, string, error) {
	var err error
	token := ""
	users, _ := u.userRepo.GetUser(username)
	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
	if users.UserID != 0 && err == nil {
		//login success
		token, err = u.jwtUsecase.GenerateJWT(u.conf, users.UserID)
		if err != nil {
			return "", "", errors.New("Error Create Token")
		}

		//save deviceID to redis with key userID
		key := fmt.Sprint("device:", users.UserID)
		u.userRepoRedis.SaveDeviceID(key, deviceID)

		//cek table notif status pending

		//if exist, send push notification

	} else {
		return "", "", errors.New("Username or Password is wrong")
	}

	return token, users.Role, err
}

func (u *user) Logout(token, deviceID string, userID int64) error {
	jwtClaims, err := u.jwtUsecase.ExtractClaims(token)
	if err != nil {
		return err
	}
	exp := jwtClaims["exp"]

	//save token to blacklist token
	err = u.userRepoRedis.Logout(token, getTokenRemainingValidity(exp))

	//remove deviceID that use for notif
	key := fmt.Sprint("device:", userID)
	u.userRepoRedis.RemoveDeviceID(key, deviceID)

	return err
}

func (u *user) GetDetailUser(userID int64) (models.User, error) {
	user, err := u.userRepo.GetDetailUser(userID)
	return user, err
}

func (u *user) UpdateUser(m models.User) error {
	var err error
	users, _ := u.userRepo.GetUser(m.Username)

	if (models.User{}) != users && users.UserID != m.UserID {
		err = errors.New("Username already exist")
	} else {
		err = u.userRepo.UpdateUser(m)
	}

	return err
}

func (u *user) UpdatePassword(pass, newPass string, userID int64) error {
	users, _ := u.userRepo.GetDetailUser(userID)
	err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(pass))
	if err != nil {
		err = errors.New("Current password is wrong")
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
		if len(hashedPassword) != 0 || err == nil {
			err = u.userRepo.UpdatePassword(string(hashedPassword[:]), userID)
		} else {
			err = errors.New("Error hash password")
		}
	}

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
