package module

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"github.com/ws-tobalobs/pkg/models"
)

func (u *user) Register(m models.User, smsNonse string, otp string) (string, error) {
	token := ""

	// get phone number
	OriginalHp, err := u.userRepoRedis.GetValue(smsNonse)
	if err != nil {
		return "", errors.New("Kode OTP kaladuarsa")
	}

	//get OTP
	originalOtp, err := u.userRepoRedis.GetValue(OriginalHp)
	if err != nil {
		return "", errors.New("Kode OTP kaladuarsa")
	}

	// cek & compare OTP redis
	if originalOtp != otp {
		return "", errors.New("Kode OTP salah")
	}

	// users, _ := u.userRepo.GetUser(m.Username)
	// if (models.User{}) == users {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if len(hashedPassword) != 0 || err == nil {
		m.Password = string(hashedPassword[:])
		userID, err := u.userRepo.Register(m)
		if err != nil {
			return "", err
		}

		token, _ = u.jwtUsecase.GenerateJWT(u.conf, userID)
		return token, err
	} else {
		log.Println(err)
		return "", errors.New("Error Hash Password")
	}
	// } else {
	// 	return "", errors.New("Username already exist")
	// }
}

func (u *user) ForgotPassword(smsNonse string, otp string, deviceID string) (string, string, error) {
	// get phone number
	OriginalHp, err := u.userRepoRedis.GetValue(smsNonse)
	if err != nil {
		return "", "", errors.New("Kode OTP kaladuarsa")
	}

	//get OTP
	originalOtp, err := u.userRepoRedis.GetValue(OriginalHp)
	if err != nil {
		return "", "", errors.New("Kode OTP kaladuarsa")
	}

	// cek & compare OTP redis
	if originalOtp != otp {
		return "", "", errors.New("Kode OTP salah")
	}

	users, err := u.userRepo.GetByPhoneNumber(OriginalHp)
	if err != nil {
		return "", "", errors.New("Error get user")
	}

	// Generate JWT token
	token, err := u.jwtUsecase.GenerateJWT(u.conf, users.UserID)
	if err != nil {
		return "", "", errors.New("Error Create Token")
	}

	//save deviceID to redis with key userID
	key := fmt.Sprint("device:", users.UserID)
	u.userRepoRedis.SaveDeviceID(key, deviceID)

	return token, users.Role, err
}

func (u *user) Verify(username, hp string, _type string) (string, error) {
	var err error
	token, _ := randToken()
	otp, _ := getRandNum()

	if _type == "register" {
		user, _ := u.userRepo.GetUser(username)
		if (models.User{}) != user {
			return "", errors.New("Username already exist")
		}

		user, _ = u.userRepo.GetByPhoneNumber(hp)
		if (models.User{}) != user {
			return "", errors.New("Phone number already exist")
		}
	} else {
		user, _ := u.userRepo.GetByPhoneNumber(hp)
		if (models.User{}) == user {
			return "", errors.New("Phone number is not registered")
		}
	}

	//save to redis : key = token, val = hp
	err = u.userRepoRedis.SetValue(token, hp, 5*time.Minute)
	if err != nil {
		return token, err
	}
	//save to redis : key = hp, val = otp
	err = u.userRepoRedis.SetValue(hp, otp, 5*time.Minute)
	if err != nil {
		return token, err
	}

	err = u.userRepoSms.Sendmessage(hp, otp)

	return token, err
}

// getRandNum returns a random number of size four for OTP code
func getRandNum() (string, error) {
	nBig, e := rand.Int(rand.Reader, big.NewInt(8999))
	if e != nil {
		return "", e
	}
	return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}

//token smsNonse or token validate
func randToken() (string, error) {
	// Using UUID V5 for generating the Token
	u4, err := uuid.NewV4()
	UUIDtoken := u4.String()
	if err != nil {
		logrus.Errorln("error:", err)
		return "", err
	}
	return UUIDtoken, nil
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

func (u *user) UpdatePassword(newPass string, userID int64) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	if len(hashedPassword) != 0 || err == nil {
		err = u.userRepo.UpdatePassword(string(hashedPassword[:]), userID)
	} else {
		err = errors.New("Error hash password")
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

func (u *user) GetKondisiMenyimpang() ([]models.KondisiMenyimpang, error) {
	res, err := u.userRepo.GetKondisiMenyimpang()

	return res, err
}
