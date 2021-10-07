package repository


import (
	"gorm.io/gorm"
	foodDomain "miniproject/business/food"
	foodDB "miniproject/repository/database/food"
)

func NewRepositoryFood(conn *gorm.DB) foodDomain.Repository {
	return foodDB.NewRepositoryMySQL(conn)
}
