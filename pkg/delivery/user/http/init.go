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

	e.POST("user/register", handler.Register)
	e.POST("user/login", handler.Login)
	e.POST("user/logout", handler.Logout)
}
