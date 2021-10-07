package main

import (
	"log"

	_apiRepo "miniproject/repository/database/OpenAPI/Spoonacular"
	_routes "miniproject/app/routes"
	_foodAPIHandler "miniproject/app/presenter/foodAPI"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	apiRepo := _apiRepo.NewFoodAPI()
	foodAPIHandler := _foodAPIHandler.NewFoodAPIHandler(apiRepo)

	routesInit := _routes.HandlerList{
		FoodAPIHandler: *foodAPIHandler,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":8000"))
}
