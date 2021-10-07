package routes

import (
	"miniproject/app/presenter/foodAPI"
	"github.com/labstack/echo/v4"
)

type HandlerList struct {
	FoodAPIHandler foodAPI.FoodAPIHandler
}

func (handlerList *HandlerList) RouteRegister(e *echo.Echo) {
	api := e.Group("foodie/v1/")
	api.GET("foodSearch", handlerList.FoodAPIHandler.GetRecipeAPI)
}