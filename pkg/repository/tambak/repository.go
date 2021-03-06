package tambak

import (
	"github.com/ws-tobalobs/pkg/models"
)

type Repository interface {
	GetAllTambak(userID int64) ([]models.Tambak, error)
	GetAllTambakForAdmin() ([]models.Tambak, error)
	GetTambakByID(tambakID int64, userID int64) (models.Tambak, error)
	GetLastMonitorTambak(tambakID int64) (models.MonitorTambak, error)
	CreateTambak(tambak models.Tambak) (int64, error)
	UpdateTambak(tambak models.Tambak) error
	PostMonitorTambak(models.MonitorTambak) (int64, error)
	PostPenyimpanganKondisiTambak(models.Notifikasi) (int64, error)
	UpdateNotifikasiKondisiTambak(status string, notifID int64)
	GetMonitorTambak(tambakID int64, tanggal string) ([]models.MonitorTambak, error)
	GetAllTambakID() ([]int64, []int64, []string, error)
	GetUserIDByTambak(tambakID int64) int64
	GetAllInfo() ([]models.Info, error)
	CreateInfo(models.Info) error
	UpdateInfo(models.Info) error
	DeleteInfo(int64) error
	GetAllPanduan() ([]models.Panduan, error)
	CreatePanduan(models.Panduan) error
	UpdatePanduan(models.Panduan) error
	DeletePanduan(int64) error
	UpdateJadwal(tambakID int64, val string, _type string) error
	GetKondisiSekarang(ip string) (models.KondisiSekarang, error)

	// GUIDELINE
	GetAllGuideline() ([]models.Guideline, error)
	CreateGuideline(models.Guideline) error
	UpdateGuideline(m models.Guideline) error

	// TUNNEL
	GetTunnel(ID int64) models.Tunnel
	SaveTunnel(models.Tunnel)
}
