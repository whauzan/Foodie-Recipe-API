package food

import (
	"miniproject/business/food"

	"gorm.io/gorm"
)

type repositoryFood struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) food.Repository {
	return &repositoryFood{
		DB: db,
	}
}

// func (rep *repositoryFood) GetAll() ([]food.Domain, error) {
// 	var foodie []Food
// 	result := rep.DB.Find(&foodie)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return toListDomain(foodie), nil
// }

// func (rep *repositoryFood) GetByID(id int) (food.Domain, error) {
// 	var foodie Food
// 	result := rep.DB.First(&foodie, "ID = ?", id)

// 	if result.Error != nil {
// 		return food.Domain{}, result.Error
// 	}

// 	return toDomain(foodie), nil
// }

func (repository repositoryFood) GetFoodByName(name string) (*food.Domain, error) {
	recordFood := Food{}
	if err := repository.DB.Where("name LIKE ?", "%"+name+"%").First(&recordFood).Error; err != nil {
		return &food.Domain{}, err
	}
	result := toDomain(recordFood)
	return &result, nil
}

func (rep *repositoryFood) Insert(foods *food.Domain) (*food.Domain, error) {
	recordFood := fromDomain(*foods)
	if err := rep.DB.Create(&recordFood).Error; err != nil {
		return &food.Domain{}, err
	}
	res := toDomain(recordFood)
	return &res, nil
}
