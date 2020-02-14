package tambak

import (
	"github.com/ws-tobalobs/pkg/models"
)

type Repository interface {
	GetAllTambak() ([]models.Tambak, error)
	PostMonitorTambak(models.MonitorTambak) error
}
