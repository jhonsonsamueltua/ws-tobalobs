package http

import (
	"github.com/labstack/echo"

	tambakUsecase "github.com/ws-tobalobs/pkg/usecase/tambak"
)

type tambak struct {
	tambakUsecase tambakUsecase.Usecase
}

func InitTambakHandler(e *echo.Echo, u tambakUsecase.Usecase) {
	handler := &tambak{
		tambakUsecase: u,
	}

	//register handler
	e.GET("/api/tambak", handler.GetAllTambak)
	e.GET("/api/tambak/admin", handler.GetAllTambakForAdmin)
	e.GET("/api/tambak/:tambakID", handler.GetTambakByID)
	e.GET("/api/tambak/last-monitor/:tambakID", handler.GetLastMonitorTambak)
	e.GET("/api/tambak/monitor/:tambakID/:tanggal", handler.GetMonitorTambak)
	e.POST("/api/tambak/monitor", handler.PostMonitorTambak)
	e.POST("/api/tambak/monitor-menyimpang", handler.PostPenyimpanganKondisiTambak)
	e.POST("/api/tambak", handler.CreateTambak)
	e.PUT("/api/tambak/:tambakID", handler.UpdateTambak)
	e.GET("/api/info", handler.GetAllInfo)
	e.POST("/api/info", handler.CreateInfo)
	e.PUT("/api/info/:infoID", handler.UpdateInfo)
	e.DELETE("/api/info/:infoID", handler.DeleteInfo)
	e.GET("/api/panduan", handler.GetAllPanduan)
	e.POST("/api/panduan", handler.CreatePanduan)
	e.PUT("/api/panduan/:panduanAplikasiID", handler.UpdatePanduan)
	e.DELETE("/api/panduan/:panduanAplikasiID", handler.DeletePanduan)
	e.PUT("/api/jadwal/:tambakID", handler.UpdateJadwal)
	e.GET("/api/get-kondisi-sekarang", handler.GetKondisiSekarang)

	//GUIDELINE
	e.GET("/api/guideline", handler.GetAllGuideline)
	e.POST("/api/guideline", handler.CreateGuideline)
	e.PUT("/api/guideline/:guidelineID", handler.UpdateGuideline)

	//TUNNEL
	e.POST("/api/save-tunnel", handler.SaveTunnel)
}
