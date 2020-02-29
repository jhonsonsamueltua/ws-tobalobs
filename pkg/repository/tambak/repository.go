package tambak

import (
	"github.com/ws-tobalobs/pkg/models"
)

type Repository interface {
	GetAllTambak(userID int64) ([]models.Tambak, error)
	GetTambakByID(tambakID int64, userID int64) (models.Tambak, error)
	GetLastMonitorTambak(tambakID int64) (models.MonitorTambak, error)
	CreateTambak(tambak models.Tambak) (int64, error)
	PostMonitorTambak(models.MonitorTambak) error
}
