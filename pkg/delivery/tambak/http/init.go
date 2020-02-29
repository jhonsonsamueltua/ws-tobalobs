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
	e.GET("/tambak", handler.GetAllTambak)
	e.GET("/tambak/:tambakID", handler.GetTambakByID)
	e.GET("/tambak/last-monitor/:tambakID", handler.GetLastMonitorTambak)
	e.POST("/tambak/monitor", handler.PostMonitorTambak)
	e.POST("/tambak", handler.CreateTambak)
}
