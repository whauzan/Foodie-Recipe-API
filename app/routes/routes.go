package routes

import (
	"miniproject/app/presenter/food"
	"github.com/labstack/echo/v4"
)

type HandlerList struct {
	FoodHandler food.FoodHandler
}

func (handlerList *HandlerList) RouteRegister(e *echo.Echo) {
	api := e.Group("foodie/v1/")
	api.GET("SearchFood", handlerList.FoodHandler.GetFoodByName)
}