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
	e.GET("/api/tambak/:tambakID", handler.GetTambakByID)
	e.GET("/api/tambak/last-monitor/:tambakID", handler.GetLastMonitorTambak)
	e.POST("/api/tambak/monitor", handler.PostMonitorTambak)
	e.POST("/api/tambak/monitor-menyimpang", handler.PostPenyimpanganKondisiTambak)
	e.POST("/api/tambak", handler.CreateTambak)
}
