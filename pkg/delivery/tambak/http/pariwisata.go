package http

import (
	"context"
	"log"
	"net/http"

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

	allTambak, err := d.tambakUsecase.GetAllTambak()
	if err != nil {
		log.Println(err)
	}
	resp.Data = allTambak

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}
