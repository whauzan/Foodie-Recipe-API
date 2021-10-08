package spoonacular

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"miniproject/business/foodAPI"
	"net/http"
	"strings"
)

type spoonacularAPI struct {
	client http.Client
}

func NewFoodAPI() foodAPI.Repository {
	return spoonacularAPI{
		client: http.Client{},
	}
}

func (s spoonacularAPI) GetFoodAPI(name string) ([]*foodAPI.Domain, error) {
	apiKey := "dd6eecd6e5694ba2af2e94916aeed314"
	splitQuery := strings.Split(name, " ")
	joinQuery := strings.Join(splitQuery, "%20")
	endpoint := fmt.Sprintf("https://api.spoonacular.com/recipes/complexSearch?apiKey=%s&query=%s&addRecipeInformation=True", apiKey, joinQuery)
	log.Println(endpoint)

	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	responseData, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	food := RecipeSource{}
	err = json.Unmarshal(responseData, &food)
	if err != nil {
		panic(err)
	}

	return toListFoodDomain(&food), nil
}

func (s spoonacularAPI) GetRecipeAPI(name string) ([]foodAPI.Domain, error) {
	apiKey := "dd6eecd6e5694ba2af2e94916aeed314"
	splitQuery := strings.Split(name, " ")
	joinQuery := strings.Join(splitQuery, "%20")
	endpoint := fmt.Sprintf("https://api.spoonacular.com/recipes/complexSearch?apiKey=%s&query=%s&addRecipeInformation=True", apiKey, joinQuery)
	log.Println(endpoint)

	resp, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	responseData, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	food := RecipeSource{}
	err = json.Unmarshal(responseData, &food)
	if err != nil {
		panic(err)
	}

	return toListDomain(food), nil
}