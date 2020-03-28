package http

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"

	"github.com/ws-tobalobs/pkg/models"
)

func (d *tambak) GetAllTambak(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	userID := c.Request().Context().Value("user") //Grab the id of the user that send the request
	userIDInt, _ := userID.(int64)

	allTambak, totalNotif, err := d.tambakUsecase.GetAllTambak(userIDInt)
	if err != nil {
		log.Println(err)
	}
	type ListData struct {
		TotalNotif int             `json:"totalNotif"`
		Data       []models.Tambak `json:"data"`
	}

	resp.Data = ListData{
		Data:       allTambak,
		TotalNotif: totalNotif,
	}
	log.Println("Total : ", totalNotif)

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) GetTambakByID(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	userID := c.Request().Context().Value("user") //Grab the id of the user that send the request
	userIDInt, _ := userID.(int64)
	tambakID := c.Param("tambakID")
	log.Println("Get By Tambak ID : ", tambakID)
	tambakIDInt, _ := strconv.ParseInt(tambakID, 10, 16)

	tambak, err := d.tambakUsecase.GetTambakByID(tambakIDInt, userIDInt)
	if err != nil {
		log.Println(err)
	}
	resp.Data = tambak

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) GetLastMonitorTambak(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	tambakID := c.Param("tambakID")
	tambakIDInt, _ := strconv.ParseInt(tambakID, 10, 16)
	log.Println("Last Monitor Tambak ID : ", tambakID)
	monitor, err := d.tambakUsecase.GetLastMonitorTambak(tambakIDInt)
	if err != nil {
		log.Println(err)
	}
	resp.Data = monitor

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) CreateTambak(c echo.Context) error {
	dt := time.Now()
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	namaTambak := c.FormValue("namaTambak")
	panjang, _ := strconv.ParseFloat(c.FormValue("panjang"), 64)
	lebar, _ := strconv.ParseFloat(c.FormValue("lebar"), 64)
	jenisBudidaya := c.FormValue("jenisBudidaya")
	// tanggalMulaiBudidaya := c.FormValue("tanggalMulaiBudidaya")

	usiaLobster, _ := strconv.Atoi(c.FormValue("usiaLobster"))
	jumlahLobster, _ := strconv.Atoi(c.FormValue("jumlahLobster"))
	jumlahLobsterJantan, _ := strconv.Atoi(c.FormValue("jumlahLobsterJantan"))
	jumlahLobsterBetina, _ := strconv.Atoi(c.FormValue("jumlahLobsterBetina"))
	// status := c.FormValue("status")
	userID := c.Request().Context().Value("user") //Grab the id of the user that send the request
	userIDInt, _ := userID.(int64)

	t := models.Tambak{}
	// t.TambakID = tambakId
	t.UserID = userIDInt
	t.NamaTambak = namaTambak
	t.Panjang = panjang
	t.Lebar = lebar
	t.JenisBudidaya = jenisBudidaya
	t.TanggalMulaiBudidaya = dt.Format("2006-01-02")
	t.UsiaLobster = usiaLobster
	t.JumlahLobster = jumlahLobster
	t.JumlahLobsterJantan = jumlahLobsterJantan
	t.JumlahLobsterBetina = jumlahLobsterBetina
	t.Status = "aktif"

	log.Println(t)

	tambakId, err := d.tambakUsecase.CreateTambak(t)
	if err != nil {
		log.Println(err)
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = "Error Save Data"
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = tambakId
	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) PostMonitorTambak(c echo.Context) error {
	dt := time.Now()
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	tambakId, _ := strconv.ParseInt(c.FormValue("tambakID"), 10, 64)
	ph, _ := strconv.ParseFloat(c.FormValue("ph"), 64)
	suhu, _ := strconv.ParseFloat(c.FormValue("suhu"), 64)
	do, _ := strconv.ParseFloat(c.FormValue("do"), 64)
	// waktuTanggal := c.FormValue("waktuTanggal")
	keterangan := c.FormValue("keterangan")

	m := models.MonitorTambak{}
	m.TambakId = tambakId
	m.PH = ph
	m.DO = do
	m.Suhu = suhu
	m.WaktuTanggal = dt.Format("2006-01-02 15:04:05")
	m.Keterangan = keterangan

	log.Println(m)
	// return nil

	monitorTambakID, err := d.tambakUsecase.PostMonitorTambak(m)
	if err != nil {
		log.Println(err)
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = "Error Save Data"
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = monitorTambakID
	log.Println(resp.Data)
	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) PostPenyimpanganKondisiTambak(c echo.Context) error {
	dt := time.Now()
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	userID := c.Request().Context().Value("user")
	userIDInt, _ := userID.(int64)
	tambakID, _ := strconv.ParseInt(c.FormValue("tambakID"), 10, 64)
	penyimpanganKondisitambakId, _ := strconv.ParseInt(c.FormValue("penyimpanganKondisiTambakID"), 10, 64)
	keterangan := c.FormValue("keterangan")

	n := models.Notifikasi{}
	n.TambakID = tambakID
	n.PenyimpanganKondisiTambakID = penyimpanganKondisitambakId
	n.TipeNotifikasi = "notif-pool-condition"
	n.Keterangan = keterangan
	n.StatusNotifikasi = "unread"
	n.WaktuTanggal = dt.Format("2006-01-02 15:04:05")

	//set sementara user id, karena dari raspberry tidak ada jwt token
	userIDInt = int64(114)
	err := d.tambakUsecase.PostPenyimpanganKondisiTambak(n, userIDInt)
	if err != nil {
		log.Println(err)
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = "Error Save Data"
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = nil
	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}
