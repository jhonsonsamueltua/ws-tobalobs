package models

import "time"

const (
	StatusSucces   = "OK"
	StatusFailed   = "failed"
	MessageSucces  = "Berhasil"
	MeassageFailed = "Tidak Berhasil"
)

type Responses struct {
	Message string      `json:"message,omitempty"`
	Status  string      `json:"status,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Config struct {
	Database DatabaseConfig
	Token    TokenData
	Redis    Redis
	Fcm      Fcm
}

type DatabaseConfig struct {
	Prod  string
	Devel string
}

type TokenData struct {
	Issuer    string
	ExpiresAt time.Time
	Key       string
	Test      string
}

type Redis struct {
	Prod  string
	Devel string
}

type Fcm struct {
	Key string
}
