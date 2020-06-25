package sms

import (
	// "github.com/souvikhaldar/gobudgetsms"

	userRepo "github.com/ws-tobalobs/pkg/repository/user"
)

type sms struct {
}

func InitSendSMS() userRepo.RepositorySMS {
	return &sms{
		// redis: redis,
	}
}
