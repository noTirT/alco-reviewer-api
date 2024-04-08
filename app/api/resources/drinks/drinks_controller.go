package drinks

import (
	"net/http"
	"noTirT/alcotracker/app/api/shared"
	"noTirT/alcotracker/util"
)

type DrinksController interface {
	CreateDrink(w http.ResponseWriter, r *http.Request)
	GetDrinksByLocationId(w http.ResponseWriter, r *http.Request)
}

type drinksController struct {
	service DrinksService
}

func NewDrinksController(service DrinksService) DrinksController {
	return &drinksController{
		service: service,
	}
}

func (controller *drinksController) GetDrinksByLocationId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := &DrinksByLocationRequest{}
	err := util.FromJSON(request, r.Body)

	if cancel := shared.HandleParseError(err, w); cancel {
		return
	}

	drinks, err := controller.service.GetDrinksByLocationId(request)

	if cancel := shared.HandleGetError(err, w); cancel {
		return
	}

	w.WriteHeader(http.StatusOK)
	util.ToJSON(&shared.GenericResponse{
		Status:  true,
		Message: "Retrieved drinks by location",
		Data:    drinks,
	}, w)
}

func (controller *drinksController) CreateDrink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := &CreateDrinkRequest{}
	err := util.FromJSON(request, r.Body)

	if cancel := shared.HandleParseError(err, w); cancel {
		return
	}

	err = controller.service.CreateDrink(request)

	if cancel := shared.HandlePostError(err, w); cancel {
		return
	}

	w.WriteHeader(http.StatusOK)
	util.ToJSON(&shared.GenericResponse{
		Status:  true,
		Message: "Created drink successfully",
	}, w)
}
