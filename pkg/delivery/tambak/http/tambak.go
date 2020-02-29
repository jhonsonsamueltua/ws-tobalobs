package http

import (
	"context"
	"log"
	"net/http"
	"strconv"

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

	allTambak, err := d.tambakUsecase.GetAllTambak(userIDInt)
	if err != nil {
		log.Println(err)
	}
	resp.Data = allTambak

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
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	// tambakId, _ := strconv.ParseInt(c.FormValue("tambakID"), 10, 64)
	namaTambak := c.FormValue("namaTambak")
	panjang, _ := strconv.ParseFloat(c.FormValue("panjang"), 64)
	lebar, _ := strconv.ParseFloat(c.FormValue("lebar"), 64)
	jenisBudidaya := c.FormValue("jenisBudidaya")
	tanggalMulaiBudidaya := c.FormValue("tanggalMulaiBudidaya")
	usiaLobster, _ := strconv.Atoi(c.FormValue("usiaLobster"))
	jumlahLobster, _ := strconv.Atoi(c.FormValue("jumlahLobster"))
	jumlahLobsterJantan, _ := strconv.Atoi(c.FormValue("jumlahLobsterJantan"))
	jumlahLobsterBetina, _ := strconv.Atoi(c.FormValue("jumlahLobsterBetina"))
	status := c.FormValue("status")
	userID := c.Request().Context().Value("user") //Grab the id of the user that send the request
	userIDInt, _ := userID.(int64)

	t := models.Tambak{}
	// t.TambakID = tambakId
	t.UserID = userIDInt
	t.NamaTambak = namaTambak
	t.Panjang = panjang
	t.Lebar = lebar
	t.JenisBudidaya = jenisBudidaya
	t.TanggalMulaiBudidaya = tanggalMulaiBudidaya
	t.UsiaLobster = usiaLobster
	t.JumlahLobster = jumlahLobster
	t.JumlahLobsterJantan = jumlahLobsterJantan
	t.JumlahLobsterBetina = jumlahLobsterBetina
	t.Status = status

	log.Println(t)

	err := d.tambakUsecase.CreateTambak(t)
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

func (d *tambak) PostMonitorTambak(c echo.Context) error {
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
	waktuTanggal := c.FormValue("waktuTanggal")
	keterangan := c.FormValue("keterangan")

	m := models.MonitorTambak{}
	m.TambakId = tambakId
	m.PH = ph
	m.DO = do
	m.Suhu = suhu
	m.WaktuTanggal = waktuTanggal
	m.Keterangan = keterangan

	log.Println(m)
	return nil

	err := d.tambakUsecase.PostMonitorTambak(m)
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
