package fcm

import (
	"github.com/appleboy/go-fcm"

	tambakFCMRepo "github.com/ws-tobalobs/pkg/repository/tambak"
)

type tambak struct {
	fcm *fcm.Client
}

func InitTambakFCMRepo(fcm *fcm.Client) tambakFCMRepo.RepositoryFCM {
	return &tambak{
		fcm: fcm,
	}
}
