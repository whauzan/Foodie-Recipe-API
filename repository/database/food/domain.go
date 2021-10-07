package food

import (
	"miniproject/business/food"
	"time"

	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Name        string
	Photo       string
	Summary     string
	Number      int
	Step        string
	HealthScore float64
	DishTypes   string
	Diets       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func toDomain(record Food) food.Domain {
	return food.Domain{
		ID:          record.ID,
		Name:        record.Name,
		Photo:       record.Photo,
		Summary:     record.Summary,
		Step:        record.Step,
		HealthScore: record.HealthScore,
		DishTypes:   record.DishTypes,
		Diets:       record.Diets,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
	}
}

func toListDomain(record []Food) (result []food.Domain) {
	result = []food.Domain{}
	for _, record := range record {
		result = append(result, toDomain(record))
	}
	return result
}

func fromDomain(record food.Domain) Food {
	return Food{
		ID:          record.ID,
		Name:        record.Name,
		Photo:       record.Photo,
		Summary:     record.Summary,
		Step:        record.Step,
		HealthScore: record.HealthScore,
		DishTypes:   record.DishTypes,
		Diets:       record.Diets,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
	}
}
