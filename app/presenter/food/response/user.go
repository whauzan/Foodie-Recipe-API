package response

import (
	"miniproject/business/food"
	"time"
)

type FoodInsertResponse struct {
	Message     string    `json:"message"`
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Photo       string    `json:"photo"`
	Summary     string    `json:"summary"`
	HealthScore float64   `json:"healthscore"`
	DishTypes   string    `json:"dishtypes"`
	Diets       string    `json:"diets"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomainInsert(domain food.Domain) FoodInsertResponse {
	return FoodInsertResponse{
		Message:     "Insert Media Succes",
		ID:          domain.ID,
		Name:        domain.Name,
		Photo:       domain.Photo,
		Summary:     domain.Summary,
		HealthScore: domain.HealthScore,
		DishTypes:   domain.DishTypes,
		Diets:       domain.Diets,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}