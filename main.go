package main

import (
	"log"

	_apiRepo "miniproject/repository/database/OpenAPI/Spoonacular"
	_routes "miniproject/app/routes"
	_dbDriver "miniproject/repository/SQL"
	_foodService "miniproject/business/food"
	_foodHandler "miniproject/app/presenter/food"
	_foodRepo "miniproject/repository/database/food"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config/config.json`)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_foodRepo.Food{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB {
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host: viper.GetString(`database.host`),
		DB_Port: viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	db := configDB.InitDB()
	dbMigrate(db)

	e := echo.New()

	apiRepo := _apiRepo.NewFoodAPI()
	foodRepo := _foodRepo.NewRepositoryMySQL(db)
	foodService := _foodService.NewService(foodRepo, apiRepo)
	foodHandler := _foodHandler.NewFoodHandler(foodService)

	routesInit := _routes.HandlerList{
		FoodHandler:    *foodHandler,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":8000"))
}
