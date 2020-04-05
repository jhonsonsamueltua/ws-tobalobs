package tambak

import "github.com/ws-tobalobs/pkg/models"

type Usecase interface {
	GetAllTambak(userID int64) ([]models.Tambak, int, error)
	GetTambakByID(tambakID int64, userID int64) (models.Tambak, error)
	GetLastMonitorTambak(tambakID int64) (models.MonitorTambak, error)
	CreateTambak(tambak models.Tambak) (int64, error)
	UpdateTambak(tambak models.Tambak) error
	PostMonitorTambak(models.MonitorTambak) (int64, error)
	PostPenyimpanganKondisiTambak(models.Notifikasi, int64) error
	GetAllInfo() ([]models.Info, error)
	GetAllPanduan() ([]models.Panduan, error)
	GetMonitorTambak(tambakID int64, tanggal string) ([]models.MonitorTambak, error)
}
