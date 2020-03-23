package module

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"

	"github.com/ws-tobalobs/pkg/models"
)

func (u *tambak) GetAllTambak(userID int64) ([]models.Tambak, error) {
	allTambak, err := u.tambakRepo.GetAllTambak(userID)

	return allTambak, err
}

func (u *tambak) GetTambakByID(tambakID int64, userID int64) (models.Tambak, error) {
	tambak, err := u.tambakRepo.GetTambakByID(tambakID, userID)

	return tambak, err
}

func (u *tambak) GetLastMonitorTambak(tambakID int64) (models.MonitorTambak, error) {
	monitor, err := u.tambakRepo.GetLastMonitorTambak(tambakID)

	return monitor, err
}

func (u *tambak) CreateTambak(t models.Tambak) (int64, error) {
	tambakID, err := u.tambakRepo.CreateTambak(t)
	if err == nil {
		execute(tambakID)
	}

	return tambakID, err
}

func (u *tambak) PostMonitorTambak(m models.MonitorTambak) (int64, error) {
	monitorTambakId, err := u.tambakRepo.PostMonitorTambak(m)
	// if err != nil {
	// 	log.Println("[Restoran][Usecase][CreateResto] Error : ", err)
	// 	return err
	// }

	return monitorTambakId, err
}

func (u *tambak) PostPenyimpanganKondisiTambak(m models.NotifikasiPenyimpanganKondisiTambak, registrationToken string) error {
	err := u.tambakRepo.PostPenyimpanganKondisiTambak(m)
	if err == nil {
		u.tambakFCMRepo.PushNotification(registrationToken)
	}

	return err
}

func execute(tambakID int64) {
	tambakIDStr := strconv.FormatInt(tambakID, 10)
	log.Println("./script/script.sh " + tambakIDStr)
	cmd := exec.Command("./script/script.sh", tambakIDStr, "&")
	err := cmd.Run()
	fmt.Println("Finished:", err)
	if err != nil {
		fmt.Printf("%s", err)
	}
}
