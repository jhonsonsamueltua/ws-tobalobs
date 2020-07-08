package user

import "github.com/ws-tobalobs/pkg/models"

type Usecase interface {
	Register(m models.User, smsNonse string, otp string) (string, error)
	Verify(username, hp string, _type string) (string, error)
	ForgotPassword(smsNonse string, otp string, deviceID string) (string, string, error)
	Login(username string, password string, deviceID string) (string, string, error)
	Logout(token, deviceID string, userID int64) error
	GetDetailUser(userID int64) (models.User, error)
	UpdateUser(models.User) error
	UpdatePassword(newPass string, userID int64) error

	// manage data dynamic
	GetKondisiMenyimpang() ([]models.KondisiMenyimpang, error)
}
