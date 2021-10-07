package foodAPI

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
	// GetHealthyRecipe(healthy string) (*Domain, error)
	// GetRecipeByCuisine(cuisine string) (*Domain, error)
	// GetRecipeByDiet(diet string) (*Domain, error)
}