package http

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/ws-tobalobs/pkg/models"
)

func (d *user) Register(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	username := c.FormValue("username")
	password := c.FormValue("password")
	nama := c.FormValue("nama")
	tanggalLahir := c.FormValue("tanggalLahir")
	noHp := c.FormValue("noHp")
	alamat := c.FormValue("alamat")

	user := models.User{
		Username:     username,
		Password:     password,
		Nama:         nama,
		Alamat:       alamat,
		NoHp:         noHp,
		TanggalLahir: tanggalLahir,
	}

	fmt.Printf("%+v\n", user)

	token, err := d.userUsecase.Register(user)
	if err != nil {
		log.Println(err)
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	auth := models.AuthResponse{
		Token:    token,
		Username: username,
	}

	fmt.Println("Berhasil register...")
	resp.Data = auth
	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *user) Login(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	username := c.FormValue("username")
	password := c.FormValue("password")

	token, err := d.userUsecase.Login(username, password)
	if err != nil {
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	auth := models.AuthResponse{
		Token:    token,
		Username: username,
	}
	log.Println("token : ", token)
	fmt.Println("Berhasil login...")
	resp.Data = auth
	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *user) Logout(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	token := c.Request().Header.Get("Authorization")
	log.Println("token : ", token)
	err := d.userUsecase.Logout(token)
	if err != nil {
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	fmt.Println("Berhasil logout...")
	resp.Data = nil
	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}
