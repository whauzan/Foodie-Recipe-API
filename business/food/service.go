package food

import (
	"miniproject/business/foodAPI"
)

type serviceFood struct {
	repository    Repository
	recipeAPIRepo foodAPI.Repository
}

func NewService(repositoryFood Repository, recipeAPIRepo foodAPI.Repository) Service {
	return &serviceFood{
		repository:    repositoryFood,
		recipeAPIRepo: recipeAPIRepo,
	}
}


// func (service *serviceFood) GetAll() ([]Domain, error) {

// 	media, _ := service.repository.GetAll()
// 	if media == nil {
// 		return nil, business.ErrNotFound
// 	}

// 	return media, nil
// }

// func (service *serviceFood) GetByID(id int) (Domain, error) {

// 	media, err := service.repository.GetByID(id)
// 	if err != nil {
// 		return Domain{}, business.ErrNotFound
// 	}

// 	return media, nil
// }

func (service *serviceFood) GetFoodByName(name string) (*Domain, error) {
	result, err := service.repository.GetFoodByName(name)
	if err != nil {
		apiFood, err := service.GetFoodAPI(name)
		if err != nil {
			return &Domain{}, err
		}
		insert, err := service.SaveFood(apiFood)
		if err != nil {
			return &Domain{}, err
		}
		return insert, nil
	}
	return result, nil
}

func (service *serviceFood) GetFoodAPI(name string) (*Domain, error) {
	result, err := service.recipeAPIRepo.GetFoodAPI(name)
	if err != nil {
		return &Domain{}, err
	}
	newRes := Domain{
		ID:          result[0].ID,
		Name:        result[0].Name,
		Photo:       result[0].Photo,
		Summary:     result[0].Summary,
		Step:        result[0].Step,
		HealthScore: result[0].HealthScore,
		DishTypes:   result[0].DishTypes,
		Diets:       result[0].Diets,
	}
	return &newRes, nil
}

func (service *serviceFood) SaveFood(food *Domain) (*Domain, error) {
	result, err := service.repository.Insert(food)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
