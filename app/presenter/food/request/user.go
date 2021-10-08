package request

import "miniproject/business/food"

// type FoodCreate struct {
// 	Food_ID string `json:"food_id"`
// }

type FoodSave struct {
	Food_Name string `json:"food_name"`
}

type Food struct {
	ID int `json:"id"`
	Name        string  `json:"name"`
	Photo       string  `json:"photo"`
	Summary     string  `json:"summary"`
	Step        string  `json:"step"`
	HealthScore float64 `json:"healthsore"`
	DishTypes   string  `json:"dishtypes"`
	Diets       string  `json:"diets"`
}

func (request *Food) ToDomain() *food.Domain {
	return &food.Domain{
		Name:        request.Name,
		Photo:       request.Photo,
		Summary:     request.Summary,
		Step:        request.Step,
		HealthScore: request.HealthScore,
		DishTypes:   request.DishTypes,
		Diets:       request.Diets,
	}
}