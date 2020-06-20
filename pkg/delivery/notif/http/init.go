package http

import (
	"github.com/labstack/echo"

	notifUsecase "github.com/ws-tobalobs/pkg/usecase/notif"
)

type notif struct {
	notifUsecase notifUsecase.Usecase
}

func InitNotifHandler(e *echo.Echo, u notifUsecase.Usecase) {
	handler := &notif{
		notifUsecase: u,
	}

	//register handler
	e.GET("/api/notif/:tambakID/:type", handler.GetAllNotif)
	e.GET("/api/notif/detail/:notifID", handler.GetDetailNotif)
	e.POST("/api/save-notif", handler.SaveNotif)
}
