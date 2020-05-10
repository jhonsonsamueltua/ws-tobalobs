package http

import (
	"context"
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
	deviceID := c.Request().Header.Get("deviceID")
	username := c.FormValue("username")
	password := c.FormValue("password")

	token, err := d.userUsecase.Login(username, password, deviceID)
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
	deviceID := c.Request().Header.Get("deviceID")
	userID := c.Request().Context().Value("user")
	userIDInt, _ := userID.(int64)

	err := d.userUsecase.Logout(token, deviceID, userIDInt)
	if err != nil {
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = nil
	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *user) GetDetailUser(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	userID := c.Request().Context().Value("user") //Grab the id of the user that send the request
	userIDInt, _ := userID.(int64)

	user, err := d.userUsecase.GetDetailUser(userIDInt)
	if err != nil {
		log.Println(err)
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}
	user.Password = ""
	resp.Data = user

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *user) UpdateUser(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	userID := c.Request().Context().Value("user") //Grab the id of the user that send the request
	userIDInt, _ := userID.(int64)
	username := c.FormValue("username")
	nama := c.FormValue("nama")
	tanggalLahir := c.FormValue("tanggalLahir")
	noHp := c.FormValue("noHp")
	alamat := c.FormValue("alamat")

	user := models.User{
		UserID:       userIDInt,
		Username:     username,
		Nama:         nama,
		Alamat:       alamat,
		NoHp:         noHp,
		TanggalLahir: tanggalLahir,
	}

	err := d.userUsecase.UpdateUser(user)
	if err != nil {
		log.Println(err)
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = nil
	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *user) UpdatePassword(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	userID := c.Request().Context().Value("user") //Grab the id of the user that send the request
	userIDInt, _ := userID.(int64)
	pass := c.FormValue("password")
	newPass := c.FormValue("newPassword")

	err := d.userUsecase.UpdatePassword(pass, newPass, userIDInt)
	if err != nil {
		log.Println(err)
		resp.Data = nil
		resp.Status = models.StatusFailed
		resp.Message = err.Error()
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = nil
	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}
