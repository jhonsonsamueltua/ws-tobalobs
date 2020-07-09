package http

import (
	"github.com/labstack/echo"

	userUsecase "github.com/ws-tobalobs/pkg/usecase/user"
)

type user struct {
	userUsecase userUsecase.Usecase
}

func InitUserHandler(e *echo.Echo, u userUsecase.Usecase) {
	handler := &user{
		userUsecase: u,
	}

	e.POST("/api/user/register", handler.Register)
	e.POST("/api/user/forgot", handler.ForgotPassword)
	e.POST("/api/user/verify", handler.Verify)
	e.POST("/api/user/login", handler.Login)
	e.POST("/api/user/logout", handler.Logout)
	e.GET("/api/user", handler.GetDetailUser)
	e.PUT("/api/user", handler.UpdateUser)
	e.PUT("/api/user/password", handler.UpdatePassword)

	//manage dynamic content
	e.GET("/api/penyimpangan-kondisi-tambak", handler.GetKondisiPenyimpangan)
	e.PUT("/api/penyimpangan-kondisi-tambak/:id", handler.UpdateKondisiMenyimpang)
}
