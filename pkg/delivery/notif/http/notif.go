package http

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/ws-tobalobs/pkg/models"
)

func (d *notif) GetAllNotif(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	userID := c.Request().Context().Value("user")
	userIDInt, _ := userID.(int64)
	typeNotif := c.Param("type")
	tambakID := c.Param("tambakID")
	tambakIDInt, _ := strconv.ParseInt(tambakID, 10, 16)

	allNotif, err := d.notifUsecase.GetAllNotif(userIDInt, tambakIDInt, typeNotif)
	if err != nil {
		log.Println(err)
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp.Data = allNotif
	
	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *notif) GetDetailNotif(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	notifID := c.Param("notifID")
	notifIDInt, _ := strconv.ParseInt(notifID, 10, 16)

	notif, err := d.notifUsecase.GetDetailNotif(notifIDInt)
	if err != nil {
		log.Println(err)
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp.Data = notif

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}
