package food

import (
	"time"
)

type Domain struct {
	ID          int
	Name        string
	Photo       string
	Summary     string
	Step        string
	HealthScore float64
	DishTypes   string
	Diets       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Service interface {
	// GetAll() ([]Domain, error)
	// GetByID(id int) (Domain, error)
	GetFoodByName(name string) (*Domain, error)
	GetFoodAPI(name string) (*Domain, error)
	SaveFood(food *Domain) (*Domain, error)
}

type Repository interface {
	// GetAll() ([]Domain, error)
	// GetByID(id int) (Domain, error)
	GetFoodByName(name string) (*Domain, error)
	Insert(food *Domain) (*Domain, error)
}
