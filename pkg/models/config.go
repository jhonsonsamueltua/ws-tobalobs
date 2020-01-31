package models

const (
	StatusSucces   = "OK"
	StatusFailed   = "failed"
	MessageSucces  = "Berhasil"
	MeassageFailed = "Tidak Berhasil"
)

type DatabaseConfig struct {
	Conn string
}

type Config struct {
	Db DatabaseConfig `mapstructure:"database"`
}

type Responses struct {
	Message string      `json:"message,omitempty"`
	Status  string      `json:"status,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
