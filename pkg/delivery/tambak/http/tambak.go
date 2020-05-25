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

	usiaLobster, _ := strconv.Atoi(c.FormValue("usiaLobster"))
	jumlahLobster, _ := strconv.Atoi(c.FormValue("jumlahLobster"))
	jumlahLobsterJantan, _ := strconv.Atoi(c.FormValue("jumlahLobsterJantan"))
	jumlahLobsterBetina, _ := strconv.Atoi(c.FormValue("jumlahLobsterBetina"))
	userID := c.Request().Context().Value("user") //Grab the id of the user that send the request
	userIDInt, _ := userID.(int64)
	pakanPagi := c.FormValue("pakanPagi")
	pakanSore := c.FormValue("pakanSore")
	gantiAir := c.FormValue("gantiAir")

	t := models.Tambak{}
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
	t.PakanPagi = pakanPagi
	t.PakanSore = pakanSore
	t.GantiAir = gantiAir

	tambakId, err := d.tambakUsecase.CreateTambak(t)
	if err != nil {
		log.Println(err)
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = tambakId
	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) UpdateTambak(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	tambakID, _ := strconv.ParseInt(c.Param("tambakID"), 10, 64)
	namaTambak := c.FormValue("namaTambak")
	panjang, _ := strconv.ParseFloat(c.FormValue("panjang"), 64)
	lebar, _ := strconv.ParseFloat(c.FormValue("lebar"), 64)
	jenisBudidaya := c.FormValue("jenisBudidaya")

	usiaLobster, _ := strconv.Atoi(c.FormValue("usiaLobster"))
	jumlahLobster, _ := strconv.Atoi(c.FormValue("jumlahLobster"))
	jumlahLobsterJantan, _ := strconv.Atoi(c.FormValue("jumlahLobsterJantan"))
	jumlahLobsterBetina, _ := strconv.Atoi(c.FormValue("jumlahLobsterBetina"))

	t := models.Tambak{}
	t.TambakID = tambakID
	t.NamaTambak = namaTambak
	t.Panjang = panjang
	t.Lebar = lebar
	t.JenisBudidaya = jenisBudidaya
	// t.TanggalMulaiBudidaya = dt.Format("2006-01-02")
	t.UsiaLobster = usiaLobster
	t.JumlahLobster = jumlahLobster
	t.JumlahLobsterJantan = jumlahLobsterJantan
	t.JumlahLobsterBetina = jumlahLobsterBetina
	// t.Status = "aktif"

	err := d.tambakUsecase.UpdateTambak(t)
	if err != nil {
		log.Println(err)
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

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
	keterangan := c.FormValue("keterangan")

	m := models.MonitorTambak{}
	m.TambakId = tambakId
	m.PH = ph
	m.DO = do
	m.Suhu = suhu
	m.WaktuTanggal = dt.Format("2006-01-02 15:04:05")
	m.Keterangan = keterangan

	monitorTambakID, err := d.tambakUsecase.PostMonitorTambak(m)
	if err != nil {
		log.Println(err)
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = monitorTambakID
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

	tambakID, _ := strconv.ParseInt(c.FormValue("tambakID"), 10, 64)
	// userID, _ := strconv.ParseInt(c.FormValue("userID"), 10, 64)
	penyimpanganKondisitambakId, _ := strconv.ParseInt(c.FormValue("penyimpanganKondisiTambakID"), 10, 64)
	keterangan := c.FormValue("keterangan")

	n := models.Notifikasi{}
	n.TambakID = tambakID
	n.PenyimpanganKondisiTambakID = penyimpanganKondisitambakId
	n.TipeNotifikasi = "notif-pool-condition"
	n.Keterangan = keterangan
	n.StatusNotifikasi = "unread"
	n.WaktuTanggal = dt.Format("2006-01-02 15:04:05")

	err := d.tambakUsecase.PostPenyimpanganKondisiTambak(n)
	if err != nil {
		log.Println(err)
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = nil
	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) GetAllInfo(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	allInfo, err := d.tambakUsecase.GetAllInfo()
	if err != nil {
		log.Println(err)
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = allInfo
	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) GetAllPanduan(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	panduan, err := d.tambakUsecase.GetAllPanduan()
	if err != nil {
		log.Println(err)
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = panduan
	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) GetMonitorTambak(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	tambakID, _ := strconv.ParseInt(c.Param("tambakID"), 10, 64)
	tanggal := c.Param("tanggal")

	m, err := d.tambakUsecase.GetMonitorTambak(tambakID, tanggal)
	if err != nil {
		log.Println(err)
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = m
	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}
