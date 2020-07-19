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
	loc, _ := time.LoadLocation("Asia/Jakarta")
	dt := time.Now().In(loc)
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
	pakanPagi := "08:00"
	pakanSore := "17:00"
	gantiAir := "3"

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
	loc, _ := time.LoadLocation("Asia/Jakarta")
	dt := time.Now().In(loc)
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
	loc, _ := time.LoadLocation("Asia/Jakarta")
	dt := time.Now().In(loc)
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

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

	err := d.tambakUsecase.PostPenyimpanganKondisiTambak(n)
	if err != nil {
		log.Println(err)
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

func (d *tambak) CreateInfo(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	judul := c.FormValue("judul")
	penjelasan := c.FormValue("penjelasan")

	i := models.Info{}
	i.Judul = judul
	i.Penjelasan = penjelasan

	err := d.tambakUsecase.CreateInfo(i)
	if err != nil {
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) UpdateInfo(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	infoID, _ := strconv.ParseInt(c.Param("infoID"), 10, 64)
	judul := c.FormValue("judul")
	penjelasan := c.FormValue("penjelasan")

	i := models.Info{}
	i.InfoID = infoID
	i.Judul = judul
	i.Penjelasan = penjelasan

	err := d.tambakUsecase.UpdateInfo(i)
	if err != nil {
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) DeleteInfo(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	infoID, _ := strconv.ParseInt(c.Param("infoID"), 10, 64)

	err := d.tambakUsecase.DeleteInfo(infoID)
	if err != nil {
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

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

func (d *tambak) CreatePanduan(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	judul := c.FormValue("judul")
	penjelasan := c.FormValue("penjelasan")

	p := models.Panduan{}
	p.Judul = judul
	p.Penjelasan = penjelasan

	err := d.tambakUsecase.CreatePanduan(p)
	if err != nil {
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) UpdatePanduan(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	panduanID, _ := strconv.ParseInt(c.Param("panduanAplikasiID"), 10, 64)
	judul := c.FormValue("judul")
	penjelasan := c.FormValue("penjelasan")

	p := models.Panduan{}
	p.PanduanAplikasiID = panduanID
	p.Judul = judul
	p.Penjelasan = penjelasan

	err := d.tambakUsecase.UpdatePanduan(p)
	if err != nil {
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) DeletePanduan(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	panduanID, _ := strconv.ParseInt(c.Param("panduanAplikasiID"), 10, 64)

	err := d.tambakUsecase.DeletePanduan(panduanID)
	if err != nil {
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

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

func (d *tambak) UpdateJadwal(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	tambakID, _ := strconv.ParseInt(c.Param("tambakID"), 10, 64)
	_type := c.FormValue("type")
	val := c.FormValue("value")

	err := d.tambakUsecase.UpdateJadwal(tambakID, val, _type)
	if err != nil {
		resp.Data = nil
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) GetAllGuideline(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	res, err := d.tambakUsecase.GetAllGuideline()
	if err != nil {
		log.Println(err)
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = res
	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) CreateGuideline(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	aksi := c.FormValue("aksiGuideline")
	notifikasi := c.FormValue("notifikasi")
	tipeBudidaya := c.FormValue("tipeBudidaya")
	tipeJadwal := c.FormValue("tipeJadwal")
	interval := c.FormValue("interval")
	waktu := c.FormValue("waktu")

	g := models.Guideline{}
	g.AksiGuideline = aksi
	g.Notifikasi = notifikasi
	g.TipeBudidaya = tipeBudidaya
	g.TipeJadwal = tipeJadwal
	g.Interval = interval
	g.Waktu = waktu

	err := d.tambakUsecase.CreateGuideline(g)
	if err != nil {
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) UpdateGuideline(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ID, _ := strconv.ParseInt(c.Param("guidelineID"), 10, 64)
	aksi := c.FormValue("aksiGuideline")
	notifikasi := c.FormValue("notifikasi")
	tipeBudidaya := c.FormValue("tipeBudidaya")
	tipeJadwal := c.FormValue("tipeJadwal")
	interval := c.FormValue("interval")
	waktu := c.FormValue("waktu")

	g := models.Guideline{}
	g.GuidelineID = ID
	g.AksiGuideline = aksi
	g.Notifikasi = notifikasi
	g.TipeBudidaya = tipeBudidaya
	g.TipeJadwal = tipeJadwal
	g.Interval = interval
	g.Waktu = waktu

	err := d.tambakUsecase.UpdateGuideline(g)
	if err != nil {
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) SaveTunnel(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	id, _ := strconv.ParseInt(c.FormValue("id"), 10, 64)
	ip := c.FormValue("ip")
	port := c.FormValue("port")

	t := models.Tunnel{}
	t.ID = id
	t.IP = ip
	t.Port = port

	d.tambakUsecase.SaveTunnel(t)

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *tambak) GetKondisiSekarang(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	log.Println("Test")
	res, err := d.tambakUsecase.GetKondisiSekarang()
	if err != nil {
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = res
	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}
