package routes

import (
	"miniproject/app/presenter/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserHandler user.UserHandler
}

func (handlerList *HandlerList) RouteRegister(e *echo.Echo) {
	api := e.Group("foodie/v1/")
	api.POST("user/register", handlerList.UserHandler.Register)
	api.POST("user/login", handlerList.UserHandler.Login)
}