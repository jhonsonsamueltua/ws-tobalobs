package http

import (
	"github.com/labstack/echo"

	tambakUsecase "github.com/ws-tobalobs/pkg/usecase/tambak"
)

type tambak struct {
	tambakUsecase tambakUsecase.Usecase
}

func InitTambakHandler(e *echo.Echo, p tambakUsecase.Usecase) {
	handler := &tambak{
		tambakUsecase: p,
	}

	//register handler
	e.GET("/tambak", handler.GetAllTambak)
}
