package notif

import (
	"github.com/labstack/echo"
)

type Delivery interface {
	GetAllNotif(c echo.Context) error
	GetDetailNotif(c echo.Context) error
	SaveNotif(c echo.Context) error
}
