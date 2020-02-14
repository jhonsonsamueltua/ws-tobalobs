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
	Db    DatabaseConfig `mapstructure:"database"`
	Token TokenData
}

type DatabaseConfig struct {
	Conn string
}

type TokenData struct {
	Issuer    string
	ExpiresAt time.Time
	Key       string
}
