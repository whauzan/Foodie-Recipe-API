package FoodAPI

type Domain struct {
	ID int
	Name string
	Photo string
	Summary string
	Step string
	HealthScore float64
	DishTypes string
	Diets string
}

type Repository interface {
	GetRecipeAPI(name string) ([]Domain, error)
}