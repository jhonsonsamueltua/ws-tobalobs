package tambak

import (
	"github.com/labstack/echo"
)

type Delivery interface {
	GetAllTambak(c echo.Context) error
	PostMonitorTambak(c echo.Context) error
}
