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
	e.POST("/tambak", handler.PostMonitorTambak)
}
