package module

import (
	"log"

	"github.com/ws-tobalobs/pkg/models"
)

func (u *tambak) GetAllTambak() ([]models.Tambak, error) {
	allTambak, err := u.tambakRepo.GetAllTambak()
	// if err != nil {
	// 	log.Println(err)
	// }

	return allTambak, err
}

func (u *tambak) PostMonitorTambak(m models.MonitorTambak) error {
	err := u.tambakRepo.PostMonitorTambak(m)
	if err != nil {
		log.Println("[Restoran][Usecase][CreateResto] Error : ", err)
		return err
	}

	return err
}
