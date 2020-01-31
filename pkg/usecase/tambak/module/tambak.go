package module

import (
	"github.com/ws-tobalobs/pkg/models"
)

func (u *tambak) GetAllTambak() ([]models.Tambak, error) {
	allTambak, err := u.tambakRepo.GetAllTambak()
	// if err != nil {
	// 	log.Println(err)
	// }

	return allTambak, err
}
