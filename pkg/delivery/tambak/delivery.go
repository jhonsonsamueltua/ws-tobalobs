package tambak

import (
	"github.com/labstack/echo"
)

type Delivery interface {
	GetAllTambak(c echo.Context) error
	GetTambakByID(c echo.Context) error
	GetLastMonitorTambak(c echo.Context) error
	PostMonitorTambak(c echo.Context) error
	PostPenyimpanganKondisiTambak(c echo.Context) error
	GetAllInfo(c echo.Context) error
}
