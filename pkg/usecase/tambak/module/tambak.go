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

func (u *tambak) CreateTambak(t models.Tambak) error {
	tambakID, err := u.tambakRepo.CreateTambak(t)
	if err == nil {
		log.Println(tambakID)
		// execute(tambakID)
	}

	return err
}

func (u *tambak) PostMonitorTambak(m models.MonitorTambak) error {
	err := u.tambakRepo.PostMonitorTambak(m)
	// if err != nil {
	// 	log.Println("[Restoran][Usecase][CreateResto] Error : ", err)
	// 	return err
	// }

	return err
}

func execute(tambakID int64) {
	tambakIDStr := strconv.FormatInt(tambakID, 10)
	// command := "./script/script.sh " + userIDStr
	log.Println("./script/script.sh " + tambakIDStr)
	cmd := exec.Command("./script/script.sh", tambakIDStr, "&")
	err := cmd.Run()
	fmt.Println("Finished:", err)
	if err != nil {
		fmt.Printf("%s", err)
	}
	// log.Println("Selesai")

	// // output := string(cmd.Output()[:])
	// fmt.Println(output)
}
